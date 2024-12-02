package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

// converts a string of numbers into []int based on seperator
func StringToInts(s string, sep string) []int {
	nums := strings.Split(s, sep)
	var output []int
	for _, num := range nums {
		numInt, _ := strconv.Atoi(num)
		output = append(output, numInt)
	}
	return output
}
