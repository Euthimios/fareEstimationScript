package csv

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
)

var testRead = []struct {
	name     string
	expected [][]string
	filePath string
}{
	{name: "wrong file path", expected: [][]string{}, filePath: "wrong file path"},
	{name: "wrong file path", expected: [][]string{}, filePath: "test/testdata/minorset.csv"},
	{name: "malformed data csv", expected: [][]string{}, filePath: "../test/testdata/malformed.csv"},
	{name: "Correct file path", expected: [][]string{0: {0: "123", 1: "38.019564", 2: "23.735323", 3: "1405592132"},
		1: {0: "123", 1: "38.019562", 2: "23.735345", 3: "1405592142"},
		2: {0: "123", 1: "38.019562", 2: "23.735345", 3: "1405592152"}},
		filePath: "../test/testdata/minorset.csv",
	},
}

func TestReadFromFile(t *testing.T) {

	for _, row := range testRead {
		result, err := ReadFromFile(row.filePath)

		if result != nil && !reflect.DeepEqual(result, row.expected) {
			t.Errorf("[TestReadFromFile] Failed %v: expected  result  data : %v and test data  %v to be equal ", row.name, result, row.expected)
		}
		if result != nil && err != nil && !reflect.DeepEqual(result, row.expected) {
			t.Errorf("[TestReadFromFile] Failed %v: expected error %v and nil result and get result %v ", row.name, err, result)
		}
	}
}

var testWrite = []struct {
	name      string
	inputData [][]string
	filePath  string
}{
	{name: "Could not create path", inputData: [][]string{}, filePath: "/terst/±±±!!@@!!/!!!!@@@/##@@"},
	{name: "Cannot create file", inputData: [][]string{}, filePath: ""},
	{name: "Correct Data,file path",
		inputData: [][]string{
			0: {0: "123", 1: "3.47"},
			1: {0: "1234", 1: "8.41"},
			2: {0: "12345", 1: "56.60"},
		},
		filePath: "../test/testdata/output.csv",
	},
}

func TestWriteToFile(t *testing.T) {
	for _, row := range testWrite {
		err := WriteToFile(row.filePath, row.inputData)
		if err == nil && !fileComparison(row.filePath, "../test/testdata/expectedresult.csv") {
			t.Errorf("[TestWriteToFile] Failed %v: expected file %v and %v to be equal", row.name, err, "../test/testdata/expectedresult.csv")
		}
	}
}

// fileComparison compares two files row by row without avoiding the load of the entire files in memory
func fileComparison(file1, file2 string) bool {

	sf, err := os.Open(file1)
	if err != nil {
		fmt.Printf("Failed to open file 1: %v; err: %v ", file1, err)
	}

	df, err := os.Open(file2)
	if err != nil {
		fmt.Printf("Failed to open file 2: %v; err: %v ", file2, err)
	}

	sscan := bufio.NewScanner(sf)
	dscan := bufio.NewScanner(df)

	for sscan.Scan() {
		dscan.Scan()
		if !bytes.Equal(sscan.Bytes(), dscan.Bytes()) {
			return false
		}
	}

	return true
}
