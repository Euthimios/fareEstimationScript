package csv

import (
	"reflect"
	"testing"
)

type read struct {
	name     string
	expected [][]string
	filePath string
}

var testRead = []read{
	{
		name:     "wrong file path",
		expected: [][]string{},
		filePath: "wrong file path",
	},
	{
		name:     "wrong file path",
		expected: [][]string{},
		filePath: "test/testdata/minorset.csv",
	},
	{
		name:     "malformed data csv",
		expected: [][]string{},
		filePath: "../test/testdata/malformed.csv",
	},
	{
		name: "Correct file path",
		expected: [][]string{
			0: {0: "123", 1: "38.019564", 2: "23.735323", 3: "1405592132"},
			1: {0: "123", 1: "38.019562", 2: "23.735345", 3: "1405592142"},
			2: {0: "123", 1: "38.019562", 2: "23.735345", 3: "1405592152"},
		},
		filePath: "../test/testdata/minorset.csv",
	},
}

func TestReadFromFile(t *testing.T) {

	for _, row := range testRead {
		result, err := ReadFromFile(row.filePath)

		if result != nil && !reflect.DeepEqual(result, row.expected) {
			t.Errorf("[TestReadFromFile] Failed: expected  result  data : %v and test data  %v to be equal ", result, row.expected)
		}

		if result != nil && err != nil && !reflect.DeepEqual(result, row.expected) {
			t.Errorf("[TestReadFromFile] Failed: expected error %v and nil result and get result %v ", err, result)
		}
	}

}
