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
func WriteToFile(path string, input <-chan []string) (<-chan int, <-chan error, error) {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return nil, nil, fmt.Errorf("Invalid file path: %v; err: %v ", path, err)
	}

	dirPath := filepath.Dir(fullPath)
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return nil, nil, fmt.Errorf("Could not create path: %v; err: %v ", path, err)
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return nil, nil, fmt.Errorf("error: %v", err)
	}

	writer := csv.NewWriter(file)
	done := make(chan int)
	chanError := make(chan error)

	go func() {
		for row := range input {
			err := writer.Write(row)
			if err != nil {
				chanError <- fmt.Errorf("error: %v", err)
			}
		}
		writer.Flush()
		file.Close()
		done <- 0
	}()
	return done, chanError, nil
}