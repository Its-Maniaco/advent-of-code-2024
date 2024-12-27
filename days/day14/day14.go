package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	err, fs := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	grid := getFinalGrid(fs, 101, 103)
	//utils.Print2DSlice(grid)
	fmt.Println("ANS: ", calcQuadrant(grid))
}

// pattern can form anytime i.e do not wait for everyone to move once and then checking
func Part2(fileLoc string) {
	err, fs := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}
	// delete output file if it exists
	err = os.Remove("Output.txt")
	if err != nil {
		log.Println("Error deleting file: ", err)
	} else {
		log.Println("File deleted.")
	}

	col, row := 101, 103
	grid := make([][]int, row)
	for k := range grid {
		grid[k] = make([]int, col)
	}

	// to store last position for each robo
	loc := make(map[int][2]int, len(fs))
	// store initial loc of each robo in map
	for i, r := range fs {
		x, y, _, _ := extractPosVel(strings.Split(r, " "))
		loc[i] = [2]int{x, y}
		grid[y][x]++
	}

	//utils.Print2DSlice(grid)

	for i := 0; i < 9000; i++ {
		// move each robot sequentially
		for roboNum, r := range fs {
			// get velocity & last location for that robo
			_, _, vx, vy := extractPosVel(strings.Split(r, " "))
			x, y := loc[roboNum][0], loc[roboNum][1]

			nx, ny := move(row, col, x, y, vx, vy)
			loc[roboNum] = [2]int{nx, ny}
			// add new location to grid
			grid[y][x]--
			grid[ny][nx]++
			//write output to file
			if search(grid) {

				err = utils.Write2DSliceToFile("Output.txt", grid, i)
				if err != nil {
					log.Println("Error creating ouput file.")
				}
			}
		}

	}
}

func getFinalGrid(fs []string, col, row int) [][]int {
	grid := make([][]int, row)
	for i := range grid {
		grid[i] = make([]int, col)
	}
	for _, r := range fs {
		//fmt.Println("r> ", r)
		x, y, vx, vy := extractPosVel(strings.Split(r, " "))
		for i := 0; i < 100; i++ {
			x, y = move(row, col, x, y, vx, vy)
		}
		grid[y][x]++
	}
	return grid
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

func search(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		// look for 5 continous non-zero number
		for j := 0; j < len(grid[0])-1-5; j++ {
			if grid[i][j] != 0 && grid[i][j+1] != 0 && grid[i][j+2] != 0 && grid[i][j+3] != 0 &&
				grid[i][j+4] != 0 {
				return true
			}
		}
	}
	return false
}
