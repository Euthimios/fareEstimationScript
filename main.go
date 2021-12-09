package main

import (
	"fmt"
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
	fmt.Println(read)
	return nil
}
