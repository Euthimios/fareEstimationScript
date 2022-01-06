package parser

import (
	"errors"
	"fareEstimationScript/internal/app/fareEstimation/ride"
	"reflect"
	"testing"
)

var positions = []struct {
	testData     [][]string
	expectedData []ride.Ride
}{
	{testData: [][]string{0: {0: "123", 1: "38.018001", 2: "23.730222", 3: "1405591942"}, 1: {0: "123", 1: "38.018001", 2: "23.730222", 3: "1405591952"}},
		expectedData: []ride.Ride{{ID: "123", Points: []ride.Point{{Latitude: 38.018001, Longitude: 23.730222, Timestamp: 1405591942}, {Latitude: 38.018001, Longitude: 23.730222, Timestamp: 1405591952}}}},
	},
	{testData: [][]string{0: {0: "123", 1: "38.018001", 2: "23.730222", 3: "1405591942"}, 1: {0: "123", 1: "38.018001", 2: "23.730222"}},
		expectedData: []ride.Ride{{ID: "123", Points: []ride.Point{{Latitude: 38.018001, Longitude: 23.730222, Timestamp: 1405591942}}}},
	},
	{testData: [][]string{0: {0: "123", 1: "38.018001", 2: "23.730222", 3: "1405591942"}, 1: {0: "123", 1: "38.018001", 2: "23.730222", 3: "1405591952"},
		2: {0: "1234", 1: "38.018011", 2: "23.730212", 3: "1638690233"}, 3: {0: "1234", 1: "38.018011", 2: "23.730212", 3: "1638690243"},
		4: {0: "12345", 1: "38.019932", 2: "23.732528", 3: "1638704645"}, 5: {0: "12345", 1: "38.019597", 2: "23.73581", 3: "1638704745"}},
		expectedData: []ride.Ride{
			{ID: "123", Points: []ride.Point{{Latitude: 38.018001, Longitude: 23.730222, Timestamp: 1405591942}, {Latitude: 38.018001, Longitude: 23.730222, Timestamp: 1405591952}}},
			{ID: "1234", Points: []ride.Point{{Latitude: 38.018011, Longitude: 23.730212, Timestamp: 1638690233}, {Latitude: 38.018011, Longitude: 23.730212, Timestamp: 1638690243}}},
			{ID: "12345", Points: []ride.Point{{Latitude: 38.019932, Longitude: 23.732528, Timestamp: 1638704645}, {Latitude: 38.019597, Longitude: 23.73581, Timestamp: 1638704745}}},
		},
	},
}

func TestParseData(t *testing.T) {
	for _, pos := range positions {
		result := ParseData(pos.testData)
		if !reflect.DeepEqual(result, pos.expectedData) {
			t.Errorf("[TestParseData] Failed: unexpected position data, expected :%v  , got %v, ", pos.expectedData, result)
		}
	}
}

var rows = []struct {
	row           []string
	expectedID    string
	expectedPoint *ride.Point
	expectedError error
}{
	{[]string{"123", "38.018001", "23.730222", "1405591942"}, "123", &ride.Point{Latitude: 38.018001, Longitude: 23.730222, Timestamp: 1405591942}, nil},
	{[]string{"1234", "38.021894", "23.735748", "1405592246"}, "1234", &ride.Point{Latitude: 38.021894, Longitude: 23.735748, Timestamp: 1405592246}, nil},
	{[]string{"14", "37.964168", "23.726123", "nil"}, "", nil, errors.New(" failed to parse row")},
	{[]string{"112344", "37.964168", "23.726123"}, "", nil, errors.New("expectd 4 elements but row hasn't")},
}

func TestParseRow(t *testing.T) {

	for _, test := range rows {
		id, result, err := parseRow(test.row)

		if id != test.expectedID {
			t.Errorf(" [TestParseRow] Failed: Parsing Failed, expect  : %v, and  get : %v", test.expectedID, id)
		}

		if result != nil && test.expectedPoint != nil && *result != *test.expectedPoint {
			t.Errorf(" [TestParseRow] Failed : Parsing Failed, expect : %v, and  get : %v", test.expectedID, id)
		}

		if err == nil && test.expectedPoint == nil {
			t.Errorf(" [TestParseRow] Failed : Parsing Failed, expect : %v, and  get : %v", test.expectedID, id)
		}
	}
}
