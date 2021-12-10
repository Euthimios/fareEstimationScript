package model

// Signal represents the position of drivers , provides information for a single reported location (lat/long) with unix time
type Signal struct {
	Latitude  float64
	Longitude float64
	Timestamp int32
}
