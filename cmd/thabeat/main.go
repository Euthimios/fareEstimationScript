package main

import (
	"flag"
	"fmt"
	"thaBeat/csv"
	"thaBeat/internal/app/thaBeat/ride/farecalculation"
	"thaBeat/internal/app/thaBeat/ride/parser"
)

func main() {

	input, output := prepare()
	err := Estimator(input, output)
	if err != nil {
		panic(err)
	}
	fmt.Println("complete")
}

func Estimator(input string, output string) error {
	read, err := csv.ReadFromFile(input)
	if err != nil {
		return fmt.Errorf("failed to open and/or read the file : %v", err)
	}
	rides := parser.ParseData(read)
	var fareEstimation [][]string
	for _, ride := range rides {
		rideEstimation := farecalculation.CalculateFare(ride)
		stringEstimation := []string{rideEstimation.IDRide, fmt.Sprintf("%.2f", rideEstimation.Total)}
		fareEstimation = append(fareEstimation, stringEstimation)
	}
	err = csv.WriteToFile(output, fareEstimation)
	if err != nil {
		return fmt.Errorf("Error writing to file: %s", err)
	}

	return nil
}

func prepare() (string, string) {
	flag.Usage = func() {
		fmt.Printf(" Fare Ride Calculation Script\n")
		fmt.Printf(" Please read bellow how to use the script , and how to use the arguments\n\n")
		flag.PrintDefaults()
	}

	input := flag.String("in", "resources/input.csv", "pleas enter the path for the file that has the Ride data")
	output := flag.String("out", "resources/output.csv", "please enter the path for the file that will have the calculated data for each Ride")
	flag.Parse()
	return *input, *output
}
