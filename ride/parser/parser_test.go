package parser

import (
	"testing"
	"thaBeat/ride/model"
)

type parsePositionResult struct {
	row            []string
	expectedID     string
	expectedSignal *model.Signal
	expectedError  error
}

var positions = []parsePositionResult{
	{[]string{"123", "38.018001", "23.730222", "1405591942"}, "123", &model.Signal{Latitude: 38.018001, Longitude: 23.730222, Timestamp: 1405591942}, nil},
	{[]string{"1", "38.021894", "23.735748", "1405592246"}, "1", &model.Signal{Latitude: 38.021894, Longitude: 23.735748, Timestamp: 1405592246}, nil},
}

func TestParsePosition(t *testing.T) {

	for _, test := range positions {
		id, result, err := parsePosition(test.row)

		if id != test.expectedID {
			t.Errorf(" Parsing Failed: expect  : %v, and  get : %v", test.expectedID, id)
		}

		if result != nil && test.expectedSignal != nil && *result != *test.expectedSignal {
			t.Errorf(" Parsing Failed: expect : %v, and  get : %v", test.expectedID, id)
		}

		if (err == nil && test.expectedError != nil) || (err != nil && test.expectedError == nil) {
			t.Errorf(" Parsing Failed: expect : %v, and  get : %v", test.expectedID, id)
		}
	}
}
