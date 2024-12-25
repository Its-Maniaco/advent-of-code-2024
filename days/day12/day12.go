package main

import (
	"fmt"
	"log"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

var directions = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func Part(fileLoc string) {
	err, grid := utils.File2DSlice(fileLoc)
	if err != nil {
		log.Fatal(err)

	}

	vis := make([][]bool, len(grid))
	for i := range vis {
		vis[i] = make([]bool, len(grid[0]))
	}
	groups := [][][]string{}
	totalPrice := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if !vis[i][j] {
				fmt.Println(grid[i][j]) //, "vis before bfs", vis)
				grp, parameter, area := bfs(grid, vis, i, j)
				groups = append(groups, grp)
				fmt.Println("A: ", area, " P: ", parameter)
				totalPrice += area * parameter
			}
		}
		// utils.PrintGrid(groups[i])
	}

	fmt.Println("Part 1: ", totalPrice)

}

// output pattern, area, parameter
func bfs(grid [][]string, vis [][]bool, startX, startY int) ([][]string, int, int) {
	// group returns the patterns
	group := make([][]string, len(grid))
	for i := 0; i < len(grid); i++ {
		group[i] = make([]string, len(grid))
	}
	group[startX][startY] = grid[startX][startY]

	parameter, area := 0, 1

	queue := [][2]int{{startX, startY}}
	vis[startX][startY] = true
	for len(queue) > 0 {
		currX, currY := queue[0][0], queue[0][1]
		queue = queue[1:]

		for _, dir := range directions {
			nx, ny := currX+dir[0], currY+dir[1]
			// valid index
			if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid) {
				// same group and not visited
				if grid[nx][ny] == grid[currX][currY] && !vis[nx][ny] {
					queue = append(queue, [2]int{nx, ny})
					vis[nx][ny] = true
					group[nx][ny] = grid[nx][ny]
					area++
				} else if grid[nx][ny] != grid[startX][startY] {
					parameter++
				}
			} else {
				parameter++
			}
		}
	}
	return group, parameter, area
}
