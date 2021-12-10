package fareCalculation

import (
	"fmt"
	"math"
	"thaBeat/ride/model"
	"time"
)

const (
	earthRadius          = float64(6371000) // const for harvesine
	minSpeed             = float64(10)      // defines the idle
	maxSpeed             = float64(100)     // max speed in order to remove from the set
	flagRate             = float64(1.3)     //default charge
	idleHourRate         = float64(11.90)   // idle charge
	movingRateDayShift   = float64(0.74)    // day  sift charge
	movingRateNightShift = float64(1.30)    // night sift charge
	minTotal             = float64(3.47)    // minimum ride fare
)

func CalculateFare(inputCh <-chan model.Ride) <-chan []string {

	data := make(chan []string)

	go func() {
		for ride := range inputCh {
			total := flagRate
			length := len(ride.LocationSignals)
			for i := 0; i < length-1; i++ {
				for j := i + 1; j < length; j++ {

					startSignal := ride.LocationSignals[i]
					endSignal := ride.LocationSignals[j]

					//the elapsed time Δt as the absolute difference of the segment endpoint timestamps
					deltaTimeSeconds := float64(endSignal.Timestamp - startSignal.Timestamp)
					//the distance covered Δs as the Haversine distance of the segment endpoint coordinates.
					deltaDistanceKm := calculateHaversine(startSignal.Longitude, startSignal.Latitude, endSignal.Longitude, endSignal.Latitude)
					// calculate the segment’s speed in khm
					speed := (deltaDistanceKm / deltaTimeSeconds) * 3600

					//if speed is > 100km/h remove the second element from the set
					if speed > maxSpeed {
						i++
						// skip the corrupted point
						continue
					}

					// calculate idle rate
					if speed <= minSpeed {
						total += (deltaTimeSeconds / 3600) * idleHourRate
						break
					}

					// calculate distance rate by hour
					if isDayRide(startSignal.Timestamp) {
						total += deltaDistanceKm * movingRateDayShift
					} else {
						total += deltaDistanceKm * movingRateNightShift
					}
					break
				}
			}
			// select  the greatest
			total = math.Max(total, minTotal)
			data <- []string{ride.ID, fmt.Sprintf("%.2f", total)}
			/*
				return &FareRide{
					IDRide: ride.ID,
					Total:  total,
				}*/
		}
	}()
	return data
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

// calculateHaversine will calculate the spherical distance as the
// crow flies between lat and lon for two given points in km by the Haversine formula
func calculateHaversine(lonFrom float64, latFrom float64, lonTo float64, latTo float64) (distance float64) {
	var deltaLat = (latTo - latFrom) * (math.Pi / 180)
	var deltaLon = (lonTo - lonFrom) * (math.Pi / 180)

	var a = math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(latFrom*(math.Pi/180))*math.Cos(latTo*(math.Pi/180))*
			math.Sin(deltaLon/2)*math.Sin(deltaLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance = earthRadius * c

	return
}
