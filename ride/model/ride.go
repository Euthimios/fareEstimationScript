package model

// Ride is a structure with the signals reported by the driver with specific ride id
type Ride struct {
	ID              string
	LocationSignals []Signal
}
