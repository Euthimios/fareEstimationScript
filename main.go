package main

import (
	"fmt"
	"thaBeat/ride/parser"
	"thaBeat/utils"
)

func main() {
	input := "files/input.csv"
	output := "files/output.csv"
	/*	if len(os.Args) < 3 {
			fmt.Println("A filepath argument is required")
			os.Exit(1)
		}
		input := os.Args[1]
		output := os.Args[2]*/

	err := estimator(input, output)
	if err != nil {
		panic(err)
	}
	fmt.Println("complete")
}

func estimator(inpath string, outPath string) error {
	read, err := utils.ReadFromFile(inpath)
	if err != nil {
		return fmt.Errorf("failed to open and/or read the file : %v", err)
	}
	parsedData, parsedDataError := parser.ParseData(read)

	fmt.Println(parsedDataError)
	fmt.Println(parsedData)

	/*	var fareEstimation [][]string
		for row := range parsedData {
			estimation := farecalculation.CalculateFare(parsedData[row])
			mySlice := []string{estimation.IDRide, fmt.Sprintf("%.2f", estimation.Total)}
			fareEstimation = append(fareEstimation, mySlice)
		}
		utils.WriteToFile(outPath, fareEstimation)*/

	return nil
}
