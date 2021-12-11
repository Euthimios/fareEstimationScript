package haversine

import (
	"testing"
)

var testsCases = []struct {
	from     Point
	to       Point
	expected float64
}{
	{
		Point{Lat: 22.55, Lon: 43.12},
		Point{Lat: 13.45, Lon: 100.28},
		6094544.408786774,
	},
	{
		Point{Lat: 51.510357, Lon: -0.116773},
		Point{Lat: 38.889931, Lon: -77.009003},
		5897658.288856054,
	},
}

func TestDistance(t *testing.T) {
	for _, input := range testsCases {
		distance := Distance(input.from, input.to)

		if input.expected != distance {
			t.Errorf("FAIL: Want distance from %v to %v to be: %v but we got %v",
				input.from,
				input.to,
				input.expected,
				distance,
			)
		}
	}
}

func BenchmarkDistance(b *testing.B) {

	from := Point{Lat: 22.55, Lon: 43.12}
	to := Point{Lat: 13.45, Lon: 100.28}
	for i := 0; i < b.N; i++ {
		_ = Distance(from, to)
	}
}
