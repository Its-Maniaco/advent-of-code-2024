package day1

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1() {
	err, fileSlice := utils.LineSlice("days/day1/Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	leftList, rightList := getSortedLists(fileSlice)
	ans := 0

	for i := 0; i < len(leftList); i++ {
		diff := leftList[i] - rightList[i]
		if diff < 0 {
			ans -= diff
		} else {
			ans += diff
		}
	}

	fmt.Println("Day 1 Puzzle 1 ans: ", ans)
}

func Part2() {
	err, fileSlice := utils.LineSlice("days/day1/Input.txt")
	if err != nil {
		log.Fatal(err)
	}

	leftList, rightList := getSortedLists(fileSlice)
	ans := linearSearch(leftList, rightList)

	fmt.Println("Day 1 Puzzle 2 ans: ", ans)
}

func getSortedLists(fileSlice []string) (leftList, rightList []int) {

	for _, line := range fileSlice {
		lineInt := strings.Split(line, "   ")
		a, _ := strconv.Atoi(lineInt[0])
		leftList = append(leftList, a)
		b, _ := strconv.Atoi(lineInt[1])
		rightList = append(rightList, b)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)
	return
}

// helper for puzzle 2
func linearSearch(list1, list2 []int) int {
	ans := 0
	for i := 0; i < len(list1); i++ {
		count := 0
		for j := 0; j < len(list2); j++ {
			if list1[i] < list2[j] {
				break
			} else if list1[i] == list2[j] {
				count++
			}
		}
		count = count * list1[i]
		ans += count
	}
	return ans
}
