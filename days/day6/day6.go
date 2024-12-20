package day6

import (
	"fmt"
	"log"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part(fileLoc string) {
	err, fs := utils.File2DSlice(fileLoc)
	if err != nil {
		log.Fatal(err.Error())
	}

	si, sj := findStart(fs)
	if si == -1 || sj == -1 {
		log.Fatal("No start found")
	}

	// track indices where guard has visited
	vis := [][]int{}

	//fmt.Println("Part 1 ans: ", bruteTravel(fs, &vis, si, sj))
	fmt.Println("Part 2 ans: ", loopCount(fs, &vis, si, sj))

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

func bruteTravel(fs [][]string, vis *[][]int, si, sj int) int {
	uniqcount := 0
	totalcount := 0
	for {
		steps := 0
		cnt := 0
		steps, cnt, si, sj = travelUp(fs, vis, si, sj)
		uniqcount += steps
		totalcount += cnt
		if si == -1 {
			break
		}

		steps, cnt, si, sj = travelRight(fs, vis, si, sj)
		uniqcount += steps
		totalcount += cnt
		if sj == len(fs) {
			break
		}

		steps, cnt, si, sj = travelDown(fs, vis, si, sj)
		uniqcount += steps
		totalcount += cnt
		if si == len(fs) {
			break
		}

		steps, cnt, si, sj = travelLeft(fs, vis, si, sj)
		uniqcount += steps
		totalcount += cnt
		if sj == -1 {
			break
		}

		// utils.PrintGrid(fs)
		// fmt.Println("---------")

		// stuck in infinite loop
		if totalcount >= len(fs)*len(fs) {
			return -1
		}
	}
	return uniqcount
}

func loopCount(fs [][]string, vis *[][]int, si, sj int) int {
	fscpy := utils.CopyGrid(fs)
	bruteTravel(fscpy, vis, si, sj)
	count := 0
	visLoop := *vis

	for _, pst := range visLoop {
		if pst[0] == si && pst[1] == sj {
			continue
		}
		fscpy2 := utils.CopyGrid(fs)
		fscpy2[pst[0]][pst[1]] = "#"
		//utils.PrintGrid(fscpy2)
		if bruteTravel(fscpy2, nil, si, sj) == -1 {
			fmt.Printf("Obstacle at: (%v,%v)\n", pst[0], pst[1])
			count++
		}
	}
	return count
}

// return travel steps and the row,col below '#'
func travelUp(fs [][]string, vis *[][]int, si, sj int) (int, int, int, int) {
	uniqcount := 0
	totalcount := 0
	for i := si; i > 0; i-- {
		totalcount++
		if i == 0 {
			break
		}
		if fs[i][sj] != "X" {
			fs[i][sj] = "X"
			uniqcount++
			if vis != nil {
				*vis = append(*vis, []int{i, sj})
			}
		}
		if fs[i-1][sj] == "#" {
			return uniqcount, totalcount, i, sj
		}
	}
	return uniqcount + 1, totalcount, -1, sj // reached edge
}

func travelDown(fs [][]string, vis *[][]int, si, sj int) (int, int, int, int) {
	uniqcount := 0
	totalcount := 0
	for i := si; i < len(fs); i++ {
		totalcount++
		if i == len(fs)-1 {
			break
		}
		if fs[i][sj] != "X" {
			fs[i][sj] = "X"
			uniqcount++
			if vis != nil {
				*vis = append(*vis, []int{i, sj})
			}
		}
		if fs[i+1][sj] == "#" {
			return uniqcount, totalcount, i, sj
		}
	}
	return uniqcount + 1, totalcount, len(fs), sj
}

func travelRight(fs [][]string, vis *[][]int, si, sj int) (int, int, int, int) {
	uniqcount := 0
	totalcount := 0
	for j := sj; j < len(fs); j++ {
		totalcount++
		if j == len(fs)-1 {
			break
		}
		if fs[si][j] != "X" {
			fs[si][j] = "X"
			uniqcount++
			if vis != nil {
				*vis = append(*vis, []int{si, j})
			}
		}
		if fs[si][j+1] == "#" {
			return uniqcount, totalcount, si, j
		}
	}
	return uniqcount + 1, totalcount, si, len(fs)
}

func travelLeft(fs [][]string, vis *[][]int, si, sj int) (int, int, int, int) {
	uniqcount := 0
	totalcount := 0
	for j := sj; j > 0; j-- {
		totalcount++
		if j == 0 {
			break
		}
		if fs[si][j] != "X" {
			fs[si][j] = "X"
			uniqcount++
			if vis != nil {
				*vis = append(*vis, []int{si, j})
			}
		}
		if fs[si][j-1] == "#" {
			return uniqcount, totalcount, si, j
		}
	}
	return uniqcount + 1, totalcount, si, -1
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
