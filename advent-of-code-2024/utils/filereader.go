package utils

import (
	"bufio"
	"os"
)

// returns a scanner for specified file
func getFileScanner(fileLoc string) (error, *bufio.Scanner) {
	file, err := os.Open(fileLoc)
	if err != nil {
		return err, nil
	}

	return nil, bufio.NewScanner(file)
}

// returns each line in slice
func LineSlice(fileLoc string) (error, []string) {
	err, scanner := getFileScanner(fileLoc)
	if err != nil {
		return err, nil
	}

	// use stringreader/writer
	var output []string
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}
	return nil, output
}
