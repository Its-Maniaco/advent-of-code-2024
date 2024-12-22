package utils

import (
	"bufio"
	"log"
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
		numInt, err := strconv.Atoi(num)
		if err != nil {
			//log.Println("error converting string to number: ", err)
			continue
		}
		output = append(output, numInt)
	}
	return output
}

// get whole input as a 2D slice of string
func File2DSlice(fileLoc string) (error, [][]string) {
	var output [][]string

	err, lines := LineSlice(fileLoc)
	if err != nil {
		return err, nil
	}

	for _, line := range lines {
		var tmp []string
		for _, ch := range line {
			tmp = append(tmp, string(ch))
		}
		output = append(output, tmp)
	}

	return nil, output
}

func File2DInt(fileLoc string) ([][]int, error) {
	err, s := LineSlice(fileLoc)
	if err != nil {
		return nil, err
	}

	out := [][]int{}
	for _, v := range s {

		tmp := StringToInts(v, "")
		out = append(out, tmp)

	}
	return out, nil
}

// return each line parsed as int
func LineInt(fileLoc string) (error, []int) {
	err, fs := LineSlice(fileLoc)
	if err != nil {
		return err, nil
	}
	nums := []int{}
	for _, s := range fs {
		num, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal("Could not parse number: ", num)
		}
		nums = append(nums, num)
	}

	return nil, nums
}
