package day14

import (
	"fmt"
	"log"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	err, fs := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	getFinalGrid(fs, 101, 103)
	// utils.Print2DSlice(grid)
	// fmt.Println("ANS: ", calcQuadrant(grid))
}

func getFinalGrid(fs []string, col, row int) {
	for l := 1; l <= 100; l++ {
		fmt.Println(l, "th second")
		grid := make([][]int, row)
		for i := range grid {
			grid[i] = make([]int, col)
		}
		for _, r := range fs {
			x, y, vx, vy := extractPosVel(strings.Split(r, " "))
			for i := 0; i < l; i++ {
				x, y = move(row, col, x, y, vx, vy)
			}
			grid[y][x]++
		}
		utils.Write2DSliceToFileDay14("days/day14/Output.txt", grid)
	}
}

// get initial pos and velocity
func extractPosVel(val []string) (int, int, int, int) {
	pst := strings.TrimPrefix(val[0], "p=")
	x, y := utils.StringToInts(pst, ",")[0], utils.StringToInts(pst, ",")[1]
	vel := strings.TrimPrefix(val[1], "v=")
	vx, vy := utils.StringToInts(vel, ",")[0], utils.StringToInts(vel, ",")[1]

	return x, y, vx, vy
}

// returns pst after moving once
func move(row, col, x, y, vx, vy int) (int, int) {
	fx := x + vx
	if vx > 0 {
		for fx >= col {
			fx -= col
		}
	} else {
		for fx < 0 {
			fx += col
		}
	}

	fy := y + vy
	if vy > 0 {
		for fy >= row {
			fy -= row
		}
	} else {
		for fy < 0 {
			fy += row
		}
	}

	// row, col
	return fx, fy
}

// calculate quadrant
func calcQuadrant(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])
	grad1, grad2, grad3, grad4 := 0, 0, 0, 0
	for i, row := range grid {
		for j, val := range row {
			if i == r/2 || j == c/2 {
				continue
			}
			// upper half
			if i < r/2 {
				// grad1
				if j < c/2 {
					grad1 += val
				} else {
					grad2 += val
				}
			} else {
				// grad3
				if j < c/2 {
					grad3 += val
				} else {
					grad4 += val
				}
			}
		}
	}
	return grad1 * grad2 * grad3 * grad4
}
