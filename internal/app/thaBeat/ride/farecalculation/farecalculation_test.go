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
	expected FareRide
}{{ride.Ride{ID: "1", Points: []ride.Point{}}, FareRide{IDRide: "1", Total: 1.3}},
	{ride.Ride{ID: "2", Points: []ride.Point{{38.018001, 23.730222, 1405591942}, {38.018001, 23.730222, 1405591952}}}, FareRide{IDRide: "2", Total: 3.47}},
	{ride.Ride{ID: "123", Points: []ride.Point{{38.018011, 23.730212, 1638690233}, {38.018011, 23.730212, 1638690243}}}, FareRide{IDRide: "123", Total: 3.47}},
	{ride.Ride{ID: "1243", Points: []ride.Point{{38.019576, 23.735345, 1639442756}, {38.019562, 23.736212, 1639443476}}}, FareRide{IDRide: "1243", Total: 3.6800000000000006}},
	{ride.Ride{ID: "1235", Points: []ride.Point{{38.019576, 23.735355, 1638704950}, {38.019562, 23.735345, 1638705853}}}, FareRide{IDRide: "1235", Total: 4.284916666666667}},
	{ride.Ride{ID: "12356", Points: []ride.Point{{38.019576, 23.735355, 1639442753}, {38.019562, 23.735345, 1639443593}}}, FareRide{IDRide: "12356", Total: 4.076666666666667}},
	{ride.Ride{ID: "1234", Points: []ride.Point{{38.019576, 23.735345, 1639357217}, {38.019562, 23.735345, 1639357700}, {38.019562, 23.735345, 1638692835}, {38.019567, 23.735437, 1638692900}}}, FareRide{IDRide: "1234", Total: 3.47}},
	{ride.Ride{ID: "6", Points: []ride.Point{{38.019982, 23.735563, 1638692912}, {38.021867, 23.735672, 1638694000}, {38.0205329, 23.729497, 1638698799}, {38.019564, 23.735323, 1638704845}, {38.019562, 23.735366, 163870485}}}, FareRide{IDRide: "6", Total: 3.47}}, // corrupted point + min fee
}

func TestCalculateFare(t *testing.T) {
	for _, test := range estimateResults {
		result := CalculateFare(test.ride)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("[TestCalculateTotalFareEstimate] Failed: expected: %v, actual: %v", test.expected, result)
		}
	}
}
