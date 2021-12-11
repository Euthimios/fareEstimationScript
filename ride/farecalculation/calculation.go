package fareCalculation

import (
	"fmt"
	"math"
	"thaBeat/ride/haversine"
	"thaBeat/ride/model"
	"time"
)

const (
	minSpeed             = float64(10)    // defines the idle
	maxSpeed             = float64(100)   // max speed in order to remove from the set
	flagRate             = float64(1.3)   //default charge
	idleHourRate         = float64(11.90) // idle charge
	movingRateDayShift   = float64(0.74)  // day  sift charge
	movingRateNightShift = float64(1.30)  // night sift charge
	minTotal             = float64(3.47)  // minimum ride fare
)

func GetCalculatedFare(inputCh <-chan model.Ride) <-chan []string {

	data := make(chan []string)

	go func() {
		for ride := range inputCh {
			estimation := CalculateFare(ride)
			data <- []string{estimation.IDRide, fmt.Sprintf("%.2f", estimation.Total)}
		}
		close(data)
	}()
	return data
}

func CalculateFare(ride model.Ride) *FareRide {

	total := flagRate
	for i := 0; i < len(ride.LocationSignals)-1; i++ {
		for j := i + 1; j < len(ride.LocationSignals); j++ {

			startPoint := ride.LocationSignals[i]
			endPoint := ride.LocationSignals[j]

			origin := haversine.Point{Lat: startPoint.Latitude, Lon: startPoint.Longitude}
			position := haversine.Point{Lat: endPoint.Latitude, Lon: endPoint.Longitude}

			//the elapsed time Δt as the absolute difference of the segment endpoint timestamps
			deltaTimeSeconds := float64(endPoint.Timestamp - startPoint.Timestamp)
			//the distance covered Δs as the Haversine distance of the segment endpoint coordinates.
			deltaDistanceKm := haversine.Distance(origin, position)

			// calculate the segment’s speed in khm
			speed := (deltaDistanceKm / deltaTimeSeconds) * 3600

			//if speed is > 100km/h remove the second element from the set
			if speed > maxSpeed {
				i++
				// skip the  point
				continue
			}
			// calculate idle rate
			if speed <= minSpeed {
				total += (deltaTimeSeconds / 3600) * idleHourRate
				break
			}
			// calculate distance rate by hour
			if isDayRide(startPoint.Timestamp) {
				total += deltaDistanceKm * movingRateDayShift
			} else {
				total += deltaDistanceKm * movingRateNightShift
			}
			break
		}
	}
	// select  the greatest
	total = math.Max(total, minTotal)
	return &FareRide{
		IDRide: ride.ID,
		Total:  total,
	}
}

//calculate if the given timestamp is at day/night sift
// TODO improve accuracy on minutes in order to catch a scenario where the ride starts few minutes befare shift changes
func isDayRide(timestamp int32) bool {
	t := time.Unix(int64(timestamp), 0).UTC()
	hour := t.Hour()
	if hour >= 5 && hour < 24 {
		return true
	}
	return false
}
