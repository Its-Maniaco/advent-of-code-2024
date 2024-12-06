package day5

import (
	"fmt"
	"log"
	"slices"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	err, lines := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	ordering := make(map[int][]int)
	flag := false // false for first section
	ansP1 := 0
	ansP2 := 0
	for _, line := range lines {
		fmt.Println(line)
		if line == "" {
			flag = true
		}
		if !flag {
			nums := utils.StringToInts(line, "|")
			fmt.Println(nums)
			k := nums[1]
			v := nums[0]
			setOrdering(ordering, k, v)
		} else {
			// second section
			nums := utils.StringToInts(line, ",")
			valid, mid := verifyUpdateValidity(ordering, nums)
			ansP1 += mid
			if !valid {
				ansP2 += getMidFromIncorrect(ordering, nums)
			}
		}
	}

	fmt.Println("Day 5 Part 1: ", ansP1)
	fmt.Println("Day 5 Part 2: ", ansP2)
}

func setOrdering(mp map[int][]int, k int, v int) {
	mp[k] = append(mp[k], v)
}

// returns true if the whole line is valid
func verifyUpdateValidity(mp map[int][]int, nums []int) (bool, int) {
	for i, num := range nums {
		for j := 0; j < i; j++ {
			if !slices.Contains(mp[num], nums[j]) {
				return false, 0
			}
		}
	}
	return true, nums[len(nums)/2]
}

// only need that element for which half elements are on left
func getMidFromIncorrect(mp map[int][]int, nums []int) int {
	for i, num := range nums {
		count := 0
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}
			if slices.Contains(mp[num], nums[j]) {
				count++
			}
		}
		if count == len(nums)/2 {
			return num
		}
	}
	return 0
}
