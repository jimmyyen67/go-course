package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("failed to read the file")
	}
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		_ = file.Close()
		return nil, errors.New("failed to read line in file")
	}

	_ = file.Close()
	return lines, nil
}
