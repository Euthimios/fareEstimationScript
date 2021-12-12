package utils

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestReadFromFile(t *testing.T) {

	testContent := [][]string{
		0: {
			0: "123",
			1: "38.018001",
			2: "23.730222",
			3: "1405591942",
		},
		1: {
			0: "123",
			1: "38.018001",
			2: "23.730222",
			3: "1405591952",
		},
		2: {
			0: "123",
			1: "38.018106",
			2: "23.729468",
			3: "1405591960",
		},
	}

	content, err := ReadFromFile("../files/test_files/test_input.csv")
	assert.Nil(t, err)
	assert.EqualValues(t, content, testContent)

	contentData, errData := ReadFromFile("../files/test_files/wrong_data_test_input.csv")
	assert.Nil(t, contentData)
	assert.NotNil(t, errData)

	content1, err1 := ReadFromFile(" ")
	if err1 == nil && content1 == nil {
		t.Errorf(" expect  content : %v,or  err1 : %v not  be null", content1, err1)
	}

}

func TestWriteToFile(t *testing.T) {

	tests := []struct {
		name         string
		inputValues  [][]string
		filePath     string
		wantOKCount  int
		wantErrCount int
	}{
		{
			name: "Succes",
			inputValues: [][]string{
				0: {0: "123", 1: "3.47"},
				1: {0: "1234", 1: "8.41"},
				2: {0: "12345", 1: "56.60"},
			},
			filePath:     "../files/test_files/outpout.csv",
			wantOKCount:  1,
			wantErrCount: 0,
		},
		{
			name: "fail",
			inputValues: [][]string{
				0: {0: "123", 1: "3.47"},
				1: {0: "1234", 1: "8.41"},
				2: {0: "12345", 1: "56.60"},
			},
			filePath:     "../files/test_files/outpout.csv",
			wantOKCount:  1,
			wantErrCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDataChannel := make(chan []string)
			var gotOK, gotErr int
			for row := range tt.inputValues {
				testDataChannel <- tt.inputValues[row]
				done, writeToFileError, err := WriteToFile("../files/test_files/outpout.csv", testDataChannel)
				if err != nil {
					t.Errorf(" unexpected error got: %v", err)
				}
				select {
				case <-done:
					gotOK++
				case <-writeToFileError:
					gotErr++
				}
			}

		})
	}
}

//compares two files , only for test purpose

func fileComparison(path1, path2 string) (bool, error) {
	file1, err1 := ioutil.ReadFile(path1)

	if err1 != nil {
		return false, fmt.Errorf("Failed to open file 1: %v; err: %v ", path1, err1)
	}

	file2, err2 := ioutil.ReadFile(path2)

	if err2 != nil {
		return false, fmt.Errorf("Failed to open file 2: %v; err: %v ", path2, err2)
	}

	equal := bytes.Equal(file1, file2)

	return equal, nil

}
