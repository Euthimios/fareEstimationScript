package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// ReadFromFileChannel gets a file path as parameter, opens a csv file, reads it
func ReadFromFileChannel(path string) ([][]string, error) {
	fullpath, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("Invalid file path: %v; err: %v ", path, err)
	}

	file, err := os.Open(fullpath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file: %v; err: %v ", fullpath, err)
	}

	reader := csv.NewReader(bufio.NewReader(file))
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data, nil
}
