package day2

import (
	"fmt"
	"log"
	"slices"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) int {
	err, fileSlice := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	var countSafe int = 0

	for _, lvlString := range fileSlice {
		lvlInt := utils.StringToInts(lvlString, " ")
		if checkSafe(lvlInt) {
			countSafe++
		}
	}

	return countSafe
}

func Part2(fileLoc string) int {
	err, fileSlice := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}
	var countSafe int = 0

	for _, lvlString := range fileSlice {
		lvlInt := utils.StringToInts(lvlString, " ")
		if safeAfterRemove(lvlInt) {
			countSafe++
		}
	}

	return countSafe
}

// removes different index of slice to check if safe
func safeAfterRemove(lvlInt []int) bool {
	fmt.Println(lvlInt)
	if checkSafe(lvlInt) {
		return true
	}
	for i := 0; i < len(lvlInt); i++ {
		temp := make([]int, len(lvlInt[0:i]))
		copy(temp, lvlInt[0:i])
		if i != len(lvlInt)-1 {
			temp = append(temp, lvlInt[i+1:]...)
		}
		fmt.Println(temp)
		if checkSafe(temp) {
			return true
		}
	}
	fmt.Println("-------")
	return false
}

// checks if safe - if sorted then check difference acceptable
func checkSafe(lvlInt []int) bool {
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
			return true
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
			return true
		}
	}

	return false
}
