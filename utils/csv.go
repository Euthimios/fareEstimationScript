package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ReadFromFile gets a file path as parameter, opens a csv file, reads it
func ReadFromFile(path string) ([][]string, error) {

	fullPath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("Invalid file path: %v; err: %v ", path, err)
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file: %v; err: %v ", fullPath, err)
	}

	reader := csv.NewReader(bufio.NewReader(file))
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data, nil
}

// WriteToFile gets a file name  and writes them in a file
func WriteToFile(path string, input [][]string) error {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return fmt.Errorf("Invalid file path: %v; err: %v ", path, err)
	}

	dirPath := filepath.Dir(fullPath)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Could not create path: %v; err: %v ", path, err)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("cannot create file; err: %v", err)
	}

	writer := csv.NewWriter(file)

	for row := range input {
		err := writer.Write(input[row])
		if err != nil {
			return fmt.Errorf("cannot write row in file; err: %v", err)
		}
	}

	writer.Flush()
	file.Close()
	return nil
}
