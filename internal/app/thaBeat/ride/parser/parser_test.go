package parser

import (
	"errors"
	"testing"
	"thaBeat/internal/app/thaBeat/ride"
)

type parsePositionRow struct {
	row           []string
	expectedID    string
	expectedPoint *ride.Point
	expectedError error
}

var positions = []parsePositionRow{
	{[]string{"123", "38.018001", "23.730222", "1405591942"}, "123", &ride.Point{Latitude: 38.018001, Longitude: 23.730222, Timestamp: 1405591942}, nil},
	{[]string{"1234", "38.021894", "23.735748", "1405592246"}, "1234", &ride.Point{Latitude: 38.021894, Longitude: 23.735748, Timestamp: 1405592246}, nil},
	{[]string{"14", "37.964168", "23.726123"}, "", nil, errors.New(" failed to parse row")},
	{[]string{"112344", "37.964168", "23.726123"}, "", nil, errors.New("expectd 4 elements but row hasn't")},
}

func TestParsePosition(t *testing.T) {

	for _, test := range positions {
		id, result, err := parseRow(test.row)

		if id != test.expectedID {
			t.Errorf(" Parsing Failed: expect  : %v, and  get : %v", test.expectedID, id)
		}

		if result != nil && test.expectedPoint != nil && *result != *test.expectedPoint {
			t.Errorf(" Parsing Failed: expect : %v, and  get : %v", test.expectedID, id)
		}

		if err == nil && test.expectedPoint == nil {
			t.Errorf(" Parsing Failed: expect : %v, and  get : %v", test.expectedID, id)
		}
	}
}
