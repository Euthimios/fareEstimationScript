package csv

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

type StreamData interface {
	ReadData() ([][]string, error)
	WriteData() error
}

type Stream struct {
	InputPath  string
	OutputPath string
	Input      [][]string
}

// ReadData ReadFromFile gets a file path as parameter, opens a csv file, reads it
func (inStream *Stream) ReadData() ([][]string, error) {
	// absolute representation of the specified path
	fullPath, err := filepath.Abs(inStream.InputPath)
	if err != nil {
		return nil, fmt.Errorf("Invalid file path: %v; err: %v ", inStream.InputPath, err)
	}
	//open the file
	file, err := os.Open(fullPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to open file: %v; err: %v ", fullPath, err)
	}
	// read the file
	reader := csv.NewReader(bufio.NewReader(file))
	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Malformed data to  file: %v; err: %v ", fullPath, err)
	}
	return data, nil
}

// WriteData WriteToFile gets a file name  and writes them in a file
func (inStream *Stream) WriteData() error {
	// absolute representation of the specified path
	fullPath, err := filepath.Abs(inStream.OutputPath)
	if err != nil {
		return fmt.Errorf("Invalid file path: %v; err: %v ", inStream.OutputPath, err)
	}

	dirPath := filepath.Dir(fullPath)
	// create path
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Could not create path: %v; err: %v ", inStream.OutputPath, err)
	}
	// create file
	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("cannot create file; err: %v", err)
	}
	// create a new writer
	writer := csv.NewWriter(file)

	for row := range inStream.Input {
		// write to csv
		err := writer.Write(inStream.Input[row])
		if err != nil {
			return fmt.Errorf("cannot write row in file; err: %v", err)
		}
	}

	writer.Flush()
	file.Close()
	return nil
}
