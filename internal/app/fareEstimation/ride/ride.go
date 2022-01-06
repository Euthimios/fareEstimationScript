package ride

// Ride is a structure with the signals reported by the driver with specific ride id
type Ride struct {
	ID     string
	Points []Point
}

// Point provides information for a single reported location (lat/long) with unix time
type Point struct {
	Latitude  float64
	Longitude float64
	Timestamp int32
}
