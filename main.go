package main

import (
	"flag"
	"fmt"
	"thaBeat/ride/fareCalculation"
	"thaBeat/ride/parser"
	"thaBeat/utils"
)

func main() {

	flag.Usage = func() {
		fmt.Printf(" Fare Ride Calculation Script\n")
		fmt.Printf(" Please read bellow how to use the script , and how to use the arguments\n\n")
		flag.PrintDefaults()
	}

	inPath := flag.String("inPath", "files/input.csv", "pleas enter the path for the file that has the Ride data")
	outPath := flag.String("outPath", "files/output.csv", "please enter the path for the file that will have the calculated data for each Ride")
	flag.Parse()

	err := estimator(*inPath, *outPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("complete")
}

func estimator(inPath string, outPath string) error {
	read, err := utils.ReadFromFile(inPath)
	if err != nil {
		return fmt.Errorf("failed to open and/or read the file : %v", err)
	}

	parsedData, parsedDataError := parser.ParseData(read)
	fareByRide := fareCalculation.CalculateFare(parsedData)
	done, writeToFileError, err := utils.WriteToFile(outPath, fareByRide)

	if err != nil {
		return fmt.Errorf("error during file write , err: %v", err)
	}

	select {
	case err := <-parsedDataError:
		if err != nil {
			return fmt.Errorf("error : %v", err)
		}
	case err := <-writeToFileError:
		if err != nil {
			return fmt.Errorf("error : %v", err)
		}
	case <-done:
	}
	return nil
}
