package day11

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	err, fs := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	in := utils.StringToInts(fs[0], " ")
	input := []uint64{}
	//int to uint64
	for _, n := range in {
		input = append(input, uint64(n))
	}

	input = blinkNTimes(input, 75)
	fmt.Println("Number of stones after 25 blinks: ", len(input))

}

func blinkNTimes(nums []uint64, n uint64) []uint64 {
	for i := uint64(0); i < n; i++ {
		nums = blink(nums)
	}
	return nums
}

// change in whole array on 1 blink
func blink(nums []uint64) []uint64 {
	var output []uint64
	for _, n := range nums {
		digitcnt := digitCount(n)
		if n == 0 {
			output = append(output, 1)
		} else if digitcnt%2 == 0 {
			// split in left and right
			left, right, err := splitNum(n)
			if err != nil {
				log.Fatal(err)
			}
			output = append(output, left, right)

		} else {
			output = append(output, n*2024)
		}
	}
	return output
}

func digitCount(n uint64) int {
	count := 1
	for n > 9 {
		count++
		n /= 10
	}
	return count
}

// split an even digit number into 2 parts
func splitNum(n uint64) (uint64, uint64, error) {
	digicnt := int(digitCount(n))
	s := fmt.Sprintf("%v", n)
	sl, sr := s[:digicnt/2], s[digicnt/2:]
	l, err := strconv.Atoi(sl)
	if err != nil {
		return 0, 0, err
	}
	r, err := strconv.Atoi(sr)
	if err != nil {
		return 0, 0, err
	}

	return uint64(l), uint64(r), nil
}
