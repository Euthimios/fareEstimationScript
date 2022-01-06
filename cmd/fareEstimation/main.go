package main

import (
	"fareEstimationScript/internal/app/fareEstimation/ride/farecalculation"
	"fareEstimationScript/internal/app/fareEstimation/ride/parser"
	"fareEstimationScript/pkg/csv"
	"flag"
	"fmt"
)

const (
	inputFile  = "assets/input.csv"
	outputFile = "assets/output.csv"
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
	var fareEstimation [][]string
	// read from file
	read, err := csv.ReadFromFile(input)
	if err != nil {
		return fmt.Errorf("failed to open and/or read the file : %v", err)
	}
	// parse the data  from the file into a Ride structure
	rides := parser.ParseData(read)
	// for each Ride proceed with fare calculation
	for _, ride := range rides {
		rideEstimation := farecalculation.CalculateFare(ride)
		stringEstimation := []string{rideEstimation.IDRide, fmt.Sprintf("%.2f", rideEstimation.Total)}
		fareEstimation = append(fareEstimation, stringEstimation)
	}
	// write the data at desired file
	err = csv.WriteToFile(output, fareEstimation)
	if err != nil {
		return fmt.Errorf("error writing to file: %s", err)
	}

	return nil
}

func prepare() (string, string) {
	// prints a usage message documenting all defined command-line flags
	flag.Usage = func() {
		fmt.Printf(" Fare Ride Calculation Script\n")
		fmt.Printf(" Please read bellow how to use the script , and how to use the arguments\n\n")
		flag.PrintDefaults()
	}

	input := flag.String("in", inputFile, "please enter the path for the file that has the Ride data")
	output := flag.String("out", outputFile, "please enter the path for the file that will have the calculated data for each Ride")
	flag.Parse()
	return *input, *output
}
