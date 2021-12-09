package main

import (
	"fmt"
	"thaBeat/ride/farecalculation"
	"thaBeat/ride/parser"
	"thaBeat/utils"
)

func main() {
	input := "files/input.csv"
	output := "files/output.csv"

	err := Estimator(input, output)
	if err != nil {
		panic(err)
	}
	fmt.Println("complete")
}

func Estimator(inpath string, outPath string) error {
	read, err := utils.ReadFromFile(inpath)
	if err != nil {
		return fmt.Errorf("failed to open and/or read the file : %v", err)
	}
	parsedData, err := parser.ParseData(read)
	if err != nil {
		return fmt.Errorf("failed to parse the data  : %v", err)
	}
	var theArray [][]string
	for row := range parsedData {
		estimation := farecalculation.CalculateFare(parsedData[row])
		mySlice := []string{estimation.IDRide, fmt.Sprintf("%.2f", estimation.Total)}
		theArray = append(theArray, mySlice)
	}
	utils.WriteToFile(outPath, theArray)

	return nil
}
