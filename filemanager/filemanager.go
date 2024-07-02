package filemanager

import (
	"bufio"
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
