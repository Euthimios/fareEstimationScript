package main

import (
	"os"
	"testing"
)

var prepareTestsCases = []struct {
	arguments []string
	expected  []string
}{
	{[]string{"-in", "Hello", "-out", "World"}, []string{0: inputFile, 1: outputFile}},
}

func TestPrepare(t *testing.T) {
	for _, row := range prepareTestsCases {
		os.Args = row.arguments
		input, output := prepare()
		if input != row.expected[0] && output != row.expected[1] {
			t.Errorf("[TestPrepare] Failed: expected input |output :%v | %v  to be equal with %v | %v, ", input, output, row.expected[0], row.expected[1])
		}
	}
}

var testMain = []struct {
	name    string
	inPath  string
	outPath string
}{
	{name: "wrong file path", inPath: "test/testdata/minorset.csv", outPath: ""},
	{name: "malformed data csv", inPath: "../../test/testdata/malformed.csv", outPath: "/terst/±±±!!@@!!/!!!!@@@/##@@"},
	{name: "Correct file path", inPath: "../../test/testdata/minorset.csv", outPath: "../../test/testdata/output.csv"},
}

func TestEstimator(t *testing.T) {
	for _, row := range testMain {
		err := Estimator(row.inPath, row.outPath)
		if err == nil && sizeFile(row.outPath) == 0 {
			t.Errorf("[TestEstimator] Failed: expected content in file :%v not to be empty", row.outPath)
		}
	}
}

// sizeFile return the size of a file ,for test purpose only
func sizeFile(file string) int64 {
	fi, _ := os.Stat(file)
	return fi.Size()
}
