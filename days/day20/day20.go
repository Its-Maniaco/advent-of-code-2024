package main

import (
	"fmt"
	"log"
	"math"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	err, fs := utils.File2DSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	start, end := findSE(fs)
	path := createPath(fs, start, end)
	count := 0
	for i := 0; i < len(path)-1; i++ {
		for j := i + 1; j < len(path); j++ {
			dist := distanceBetweenPoints(path[i], path[j])
			// dist is the time cheat is active for
			if dist <= 2 {
				stepDiff := path[i][2] - path[j][2]
				if stepDiff < 0 {
					stepDiff *= -1
				}
				// steps saved if cheat is activated
				stepSaved := stepDiff - 2
				if stepSaved >= 100 {
					count++
				}
			}
		}
	}
	fmt.Println("Count: ", count)
}

func Part2(fileLoc string) {
	err, fs := utils.File2DSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	start, end := findSE(fs)
	path := createPath(fs, start, end)
	count := 0
	for i := 0; i < len(path)-1; i++ {
		for j := i + 1; j < len(path); j++ {
			dist := distanceBetweenPoints(path[i], path[j])
			if dist <= 20 {
				stepDiff := path[i][2] - path[j][2]
				if stepDiff < 0 {
					stepDiff *= -1
				}
				stepSaved := stepDiff - dist
				if stepSaved >= 100 {
					count++
				}
			}
		}
	}
	fmt.Println("Count: ", count)
}

// identify route from S to E and steps from start it takes to reach that point
func createPath(grid [][]string, start, end []int) [][]int {
	path := [][]int{}

	cpy := utils.CopyGrid(grid)
	dir := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	step := 1

	// add start point
	path = append(path, []int{start[0], start[1], 0})

	i, j := start[0], start[1]
	for !(i == end[0] && j == end[1]) {
		for _, k := range dir {
			if i+k[0] >= 0 && i+k[0] < len(cpy) && j+k[1] >= 0 && j+k[1] < len(cpy[1]) {
				if cpy[i+k[0]][j+k[1]] == "." || cpy[i+k[0]][j+k[1]] == "E" {
					cpy[i][j] = "O"
					i += k[0]
					j += k[1]

					path = append(path, []int{i, j, step})
					step++
					break
				}
			}
		}
	}

	/*
		sort.Slice(path, func(i, j int) bool {
			d1 := distance(path[i])
			d2 := distance(path[j])
			if d1 == d2 {
				return path[i][2] < path[j][2] // Sort by step if distances are equal
			}
			return d1 < d2
		})
	*/

	return path
}

// find start and end points
func findSE(grid [][]string) ([]int, []int) {
	start := make([]int, 2)
	end := make([]int, 2)
	for i, row := range grid {
		for j, ch := range row {
			if ch == "S" {
				start[0] = i
				start[1] = j
			} else if ch == "E" {
				end[0] = i
				end[1] = j
			}
		}
	}

	return start, end
}

// calc linear distance between 2 points
func distanceBetweenPoints(p1, p2 []int) int {
	i := p1[0] - p2[0]
	if i < 0 {
		i *= -1
	}
	j := p1[1] - p2[1]
	if j < 0 {
		j *= -1
	}

	return i + j
}

// for sorting
func distance(a []int) float64 {
	return math.Sqrt(float64(a[0]*a[0] + a[1]*a[1]))
}
