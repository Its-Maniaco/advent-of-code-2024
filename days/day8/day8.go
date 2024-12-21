package day8

import (
	"fmt"
	"log"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	fmt.Println("Started")
	err, fs := utils.File2DSlice(fileLoc)
	if err != nil {
		log.Fatal(err.Error())
	}

	nodeMap := getNodeMap(fs)
	nodeLocs := make(map[string][][]int)
	ans := 0
	for _, v := range nodeMap {
		ans += getAntiNode(fs, nodeLocs, v)
	}

	fmt.Println(ans)
}

// get total number of antinodes and sets them in process
func getAntiNode(grid [][]string, nodelocs map[string][][]int, pst [][]int) int {
	count := 0
	for i := 0; i < len(pst)-1; i++ {
		for j := i + 1; j < len(pst); j++ {
			r1, c1 := pst[i][0], pst[i][1]
			r2, c2 := pst[j][0], pst[j][1]
			fmt.Printf("Checking: (%v,%v),(%v,%v)\n", r1, c1, r2, c2)
			count += setAntinode(grid, nodelocs, r1, c1, r2, c2)
		}
	}
	return count
}

// checks and sets the 2 possible indices for antinode
func setAntinode(grid [][]string, nodelocs map[string][][]int, r1, c1, r2, c2 int) int {
	count := 0
	for i := 0; ; i++ {
		rD, cD := i*intAbs(r2-r1), i*intAbs(c2-c1) // differences between indices

		// if point 1 is right and up relative to point 2
		pr1, pc1 := r1-rD, c1+cD // possible row
		pr2, pc2 := r2+rD, c2-cD // possible column
		// if point 1 is left and up relative to point 2
		if c1 < c2 {
			pr1, pc1 = r1-rD, c1-cD
			pr2, pc2 = r2+rD, c2+cD
		}

		if (pr1 < 0 || pr1 >= len(grid) || pc1 < 0 || pc1 >= len(grid)) && (pr2 < 0 || pr2 >= len(grid) || pc2 < 0 || pc2 >= len(grid)) {
			break
		}

		fmt.Printf("Trying#: (%v,%v)\n", pr1, pc1)

		count += setGrid(grid, nodelocs, grid[r1][c1], pr1, pc1)
		fmt.Printf("Trying#: (%v,%v)\n", pr2, pc2)
		count += setGrid(grid, nodelocs, grid[r1][c1], pr2, pc2)
	}

	return count
}

// sets antinode for indices
func setGrid(grid [][]string, nodelocs map[string][][]int, nodelocKey string, r, c int) int {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return 0
	}

	if grid[r][c] != "#" {
		grid[r][c] = "#"
		nodelocs[nodelocKey] = append(nodelocs[nodelocKey], []int{r, c})
		utils.PrintGrid(grid)
		fmt.Println("---------")
		return 1
	}
	return 0
}

// get inital map of all characters
func getNodeMap(grid [][]string) map[string][][]int {
	mp := make(map[string][][]int)
	for i, row := range grid {
		for j, ch := range row {
			if ch != "." {
				mp[ch] = append(mp[ch], []int{i, j})
			}
		}
	}
	return mp
}

func intAbs(a int) int {
	if a > 0 {
		return a
	}
	return a * -1
}
