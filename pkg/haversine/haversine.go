package haversine

import (
	"math"
)

var earthRadiusMetres float64 = 6371000

// Point provides information for a location
type Point struct {
	Lat float64
	Lon float64
}

// Delta struct represents the delta between points
type Delta struct {
	Lat float64
	Lon float64
}

// Delta method calculates the delta between starting and end point
func (p Point) Delta(point Point) Delta {
	return Delta{
		Lat: p.Lat - point.Lat,
		Lon: p.Lon - point.Lon,
	}
}

// toRadians method transforms a point coordinate system (lat/long) to radian
func (p Point) toRadians() Point {
	return Point{
		Lat: degreesToRadians(p.Lat),
		Lon: degreesToRadians(p.Lon),
	}
}

// degreesToRadians transforms lat/log to radian
func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// Distance will calculate the spherical distance as the
// crow flies between lat and lon for two given points by the Haversine formula in khm
func Distance(origin, position Point) float64 {
	origin = origin.toRadians()
	position = position.toRadians()

	change := origin.Delta(position)

	a := math.Pow(math.Sin(change.Lat/2), 2) + math.Cos(origin.Lat)*math.Cos(position.Lat)*math.Pow(math.Sin(change.Lon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return (earthRadiusMetres * c) / 1000
}
