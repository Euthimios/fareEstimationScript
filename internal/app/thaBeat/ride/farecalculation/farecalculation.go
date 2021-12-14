package farecalculation

import (
	"math"
	"thaBeat/internal/app/thaBeat/haversine"
	"thaBeat/internal/app/thaBeat/ride"
	"time"
)

// FareRide represents a ride with id and calculated total fee
type FareRide struct {
	IDRide string
	Total  float64
}

const (
	minSpeed             = float64(10)    // defines the idle
	maxSpeed             = float64(100)   // max speed in order to remove from the set
	flagRate             = float64(1.3)   // default charge
	idleHourRate         = float64(11.90) // idle charge
	movingRateDayShift   = float64(0.74)  // day  sift charge
	movingRateNightShift = float64(1.30)  // night sift charge
	minTotal             = float64(3.47)  // minimum ride fare
)

// CalculateFare gets as parameter Ride objects and for each of them a fare is calculated
func CalculateFare(r ride.Ride) FareRide {

	fare := FareRide{
		IDRide: r.ID,
		Total:  flagRate,
	}
	if len(r.Points) == 0 {
		return fare
	}
	// First point is the start point
	startPoint := r.Points[0]
	// We start iterating from second point
	for i := 1; i < len(r.Points); i++ {

		endPoint := r.Points[i]
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
			// skip the corrupted point
			continue
		}

		// calculate idle rate
		if speed <= minSpeed {
			fare.Total += (deltaTimeSeconds / 3600) * idleHourRate
			startPoint = endPoint
			continue
		}

		// calculate distance rate by hour
		if isDayRide(startPoint.Timestamp) {
			fare.Total += deltaDistanceKm * movingRateDayShift
		} else {
			fare.Total += deltaDistanceKm * movingRateNightShift
		}
		startPoint = endPoint
	}
	// select  the greatest
	fare.Total = math.Max(fare.Total, minTotal)

	return fare
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
