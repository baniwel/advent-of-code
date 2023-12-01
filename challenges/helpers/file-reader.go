package helpers

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, error := os.Open(path)
	if error != nil {
		return nil, error
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
