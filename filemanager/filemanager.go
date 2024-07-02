package filemanager

import (
	"bufio"
	"encoding/json"
	"os"
)

func ReadLines(filename string) ([]string, error) {
	file, err := os.Open("prices.txt")

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

func WriteJSON(path string, data any) error {
	file, err := os.Create(path)
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
