package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func New(inputFilePath string, outputFilePath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputFilePath,
		OutputFilePath: outputFilePath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		defer file.Close()
		return nil, err
	}

	lines := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		defer file.Close()
		return nil, err
	}

	defer file.Close()

	return lines, nil
}

func (fm FileManager) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		defer file.Close()
		return err
	}

	defer file.Close()
	return nil
}
