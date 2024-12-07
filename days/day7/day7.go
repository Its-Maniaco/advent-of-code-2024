package day7

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	err, fs := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}
	ans := 0
	count := 0 // count tracks no. of solutions for a line, made if part 2 might have asked

	for _, line := range fs {
		cmp, nums := parseInput(line)
		count = 0
		log.Println("Checking for ", line)
		recur(cmp, nums, nums[0], 1, &count)
		if count > 0 {
			ans += cmp
		}
	}
	fmt.Println(ans)
}

func recur(cmp int, nums []int, sum int, s int, count *int) int {
	/*base condition: all nums have been accounted*/
	if s == len(nums) {
		if sum == cmp {
			return 1
		}
		return 0
	}

	/* we can only pick the next immediate number from nums, i.e i<s+1
	i<len(nums) will skip numbers in between for pnc in sign
	*/
	for i := s; i < s+1; i++ {
		if sum+nums[i] <= cmp {
			*count = *count + recur(cmp, nums, sum+nums[i], s+1, count)
		}
		if sum*nums[i] <= cmp {
			*count = *count + recur(cmp, nums, sum*nums[i], s+1, count)
		}
	}
	return 0
}

func parseInput(line string) (int, []int) {
	var output []int
	tmp := strings.Split(line, ":")

	a, err := strconv.Atoi(tmp[0])
	if err != nil {
		log.Fatal("could not parse input")
	}

	output = append(output, utils.StringToInts(tmp[1], " ")...)
	return a, output
}
