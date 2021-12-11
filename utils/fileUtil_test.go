package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestReadFromFile(t *testing.T) {
	content, err := ReadFromFile("../files/input.csv")
	if err != nil {
		t.Error("Failed to read csv data")
	}
	fmt.Print(content)
}

func TestWriteToFile(t *testing.T) {

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
