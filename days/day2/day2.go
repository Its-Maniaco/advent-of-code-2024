package day2

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) int {
	err, fileSlice := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	var countSafe int = 0

	for _, lvlString := range fileSlice {
		// convert level's string into int
		nums := strings.Split(lvlString, " ")
		var lvlInt []int //= make([]int, len(nums)) //to store level's number in int format
		for _, num := range nums {
			numInt, _ := strconv.Atoi(num)
			lvlInt = append(lvlInt, numInt)
		}

		// checking if safe - if sorted, then check if difference acceptable
		if slices.IsSortedFunc(lvlInt, func(a, b int) int {
			return a - b
		}) {
			var flag bool = true // true for acceptable difference
			for i := 0; i < len(lvlInt)-1; i++ {
				// not safe
				if lvlInt[i+1]-lvlInt[i] == 0 || lvlInt[i+1]-lvlInt[i] > 3 {
					flag = false
					break
				}
			}
			if flag == true {
				countSafe++
			}
		} else if slices.IsSortedFunc(lvlInt, func(a, b int) int {
			return b - a
		}) {
			var flag bool = true // true for acceptable difference
			for i := 0; i < len(lvlInt)-1; i++ {
				// not safe
				if lvlInt[i+1]-lvlInt[i] == 0 || lvlInt[i+1]-lvlInt[i] < -3 {
					flag = false
					break
				}
			}
			if flag == true {
				countSafe++
			}
		}
	}

	fmt.Println("Safe levels: ", countSafe)
	return countSafe
}
