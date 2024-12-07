package day6

import (
	"fmt"
	"log"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	err, fs := utils.File2DSlice(fileLoc)
	if err != nil {
		log.Fatal(err.Error())
	}

	si, sj := findStart(fs)
	if si == -1 || sj == -1 {
		log.Fatal("No start found")
	}

	fmt.Println(bruteTravel(fs, si, sj))

}

func findStart(fS [][]string) (int, int) {
	for i, row := range fS {
		for j, ch := range row {
			if ch == "^" {
				fmt.Printf("Start: %v %v\n", i, j)
				return i, j
			}
		}
	}
	return -1, -1
}

func bruteTravel(fs [][]string, si, sj int) int {
	count := 0
	for {
		steps := 0
		steps, si, sj = travelUp(fs, si, sj)
		count += steps
		if si == -1 {
			break
		}

		steps, si, sj = travelRight(fs, si, sj)
		count += steps
		if sj == len(fs) {
			break
		}

		steps, si, sj = travelDown(fs, si, sj)
		count += steps
		if si == len(fs) {
			break
		}

		steps, si, sj = travelLeft(fs, si, sj)
		count += steps
		if sj == -1 {
			break
		}
		utils.PrintGrid(fs)
		fmt.Println("---------")
	}
	return count
}

// return travel steps and the row,col below '#'
func travelUp(fs [][]string, si, sj int) (int, int, int) {
	count := 0
	for i := si; i > 0; i-- {
		if i == 0 {
			break
		}
		if fs[i][sj] != "X" {
			fs[i][sj] = "X"
			count++
		}
		if fs[i-1][sj] == "#" {
			return count, i, sj
		}
	}
	return count + 1, -1, sj // reached edge
}

func travelDown(fs [][]string, si, sj int) (int, int, int) {
	count := 0
	for i := si; i < len(fs); i++ {
		if i == len(fs)-1 {
			break
		}
		if fs[i][sj] != "X" {
			fs[i][sj] = "X"
			count++
		}
		if fs[i+1][sj] == "#" {
			return count, i, sj
		}
	}
	return count + 1, len(fs), sj
}

func travelRight(fs [][]string, si, sj int) (int, int, int) {
	count := 0
	for j := sj; j < len(fs); j++ {
		if j == len(fs)-1 {
			break
		}
		if fs[si][j] != "X" {
			fs[si][j] = "X"
			count++
		}
		if fs[si][j+1] == "#" {
			return count, si, j
		}
	}
	return count + 1, si, len(fs)
}

func travelLeft(fs [][]string, si, sj int) (int, int, int) {
	count := 0
	for j := sj; j > 0; j-- {
		if j == 0 {
			break
		}
		if fs[si][j] != "X" {
			fs[si][j] = "X"
			count++
		}
		if fs[si][j-1] == "#" {
			return count, si, j
		}
	}
	return count + 1, si, -1
}

/* TODO: Optimize
func travelVertical(fs [][]string, si, sj int) (int, int) {
	rLmt, rSub := 0, 1
	if fs[si][sj] == "V" {
		rLmt, rSub = len(fs), -1
	}
	count := 0
	for i := si; i >= rLmt; i-- {
		if i == 0 {
			return count, -1
		}
		if fs[i-rSub][sj] == "#" {
			return count, i
		}
		count++
	}
	return count, -1
}
*/
