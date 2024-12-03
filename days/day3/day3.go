package day3

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	err, lines := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	exp := `mul\(\d+,\d+\)`
	matches := getMatches(lines, exp)

	result := 0

	// parse the combinations
	for _, match := range matches {
		s := strings.Split(match, ",")           // split around ','
		a, _ := strconv.Atoi(s[0][4:])           // mul(....
		b, _ := strconv.Atoi(s[1][:len(s[1])-1]) //.....)
		result += a * b
	}

	fmt.Println(result)
}

func Part2(fileLoc string) {
	err, lines := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	exp := `mul\(\d+,\d+\)|don't\(\)|do\(\)`
	matches := getMatches(lines, exp)

	result := 0

	var flag bool = true // to allow multiplying
	// parse the combinations
	for _, match := range matches {
		if match == "do()" {
			flag = true
			continue
		} else if match == "don't()" {
			flag = false
			continue
		}

		if flag == true {
			s := strings.Split(match, ",")           // split around ','
			a, _ := strconv.Atoi(s[0][4:])           // mul(....
			b, _ := strconv.Atoi(s[1][:len(s[1])-1]) //.....)
			result += a * b
		}

	}

	fmt.Println(result)
}

func getMatches(lines []string, exp string) []string {
	re, err := regexp.Compile(exp)
	if err != nil {
		log.Println("Cannot compile regex: ", err)
	}

	// get all valid combinations
	var matches []string
	for _, line := range lines {
		match := re.FindAllString(line, -1)
		fmt.Println(match)
		matches = append(matches, match...)
	}

	return matches
}
