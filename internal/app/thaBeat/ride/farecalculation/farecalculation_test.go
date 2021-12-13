package farecalculation

import (
	"reflect"
	"testing"
	"thaBeat/internal/app/thaBeat/ride"
)

var isDayRideResult = []struct {
	timestamp int32
	expected  bool
}{
	{1639316817, true},
	{1639342017, true},
	{1639356417, false},
	{1639357917, false},
}

func TestIsDayRide(t *testing.T) {
	for _, test := range isDayRideResult {
		result := isDayRide(test.timestamp)

		if result != test.expected {
			t.Errorf("[TestIsDayRide] Failed: expected: %t, actual: %t, timestamp: %d", test.expected, result, test.timestamp)
		}
	}
}

var estimateResults = []struct {
	ride     ride.Ride
	expected *FareRide
}{
	{ride.Ride{ID: "a", Points: []ride.Point{{38.020539, 23.729497, 1405592080}, {38.019564, 25.735323, 1405594132}}}, &FareRide{IDRide: "a", Total: 3.47}},
	{ride.Ride{ID: "123", Points: []ride.Point{{38.019326, 23.72896, 1638690560}, {38.019397, 23.728953, 1638690960}}}, &FareRide{IDRide: "123", Total: 7.159777115033196}},
	{ride.Ride{ID: "1235", Points: []ride.Point{{38.019576, 23.735345, 1638665085}, {38.019562, 23.735345, 1638665506}}}, &FareRide{IDRide: "1235", Total: 3.47}},
	{ride.Ride{ID: "1234", Points: []ride.Point{{38.019576, 23.735345, 1638692330}, {38.019562, 23.735345, 1638692730}, {38.019562, 23.735345, 1638692835}, {38.019567, 23.735437, 1638692900}}}, &FareRide{IDRide: "1234", Total: 3.47}},
	{ride.Ride{ID: "6", Points: []ride.Point{{38.019982, 23.735563, 1638692912}, {38.021867, 23.735672, 1638694000}, {38.0205329, 23.729497, 1638698799}, {38.019564, 23.735323, 1638704845}, {38.019562, 23.735366, 163870485}}}, &FareRide{IDRide: "6", Total: 3.47}}, // corrupted point + min fee
}

func TestCalculateFare(t *testing.T) {
	for _, test := range estimateResults {
		result := CalculateFare(test.ride)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("[TestCalculateTotalFareEstimate] Failed: expected: %v, actual: %v", *test.expected, *result)
		}
	}
}
