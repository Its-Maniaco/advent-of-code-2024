package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

type direction struct {
	Dx, Dy int
	symbol string
}

var dirs = []direction{
	{-1, 0, "^"},
	{1, 0, "v"},
	{0, -1, "<"},
	{0, 1, ">"},
}

func Part1(fileLoc string) {
	keypad := returnKeypad()
	keypadMap := buildMap(keypad)

	dirpad := returnDirpad()
	dirpadMap := buildMap(dirpad)

	err, fs := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	ans := 0

	for _, row := range fs {
		// get command for final keypad robo
		keypadRobo := roboCommand(row, keypadMap, "A")
		//get command for keypad controlling robo
		keypadRoboController := roboCommand(keypadRobo, dirpadMap, "A")
		// get commands for robo being controlled by human
		firstRobo := roboCommand(keypadRoboController, dirpadMap, "A")
		fmt.Println(firstRobo)
		num, err := strconv.Atoi(row[:len(row)-1])
		if err != nil {
			num = 0
			log.Printf("Error parsing %v into number: \n", num)
		}
		ans += num * len(firstRobo)
	}

	fmt.Println("Ans: ", ans)
}

// returns complete combo for robot on keypad
func roboCommand(s string, padMap map[string]map[string]string, start string) string {
	res := ""

	for _, char := range s {
		charString := string(char)
		res += padMap[start][charString]
		start = charString
		// for pressing that button
		res += "A"
	}
	return res
}

// key: number to reach, val: shortest part to reach key from that location
func buildMap(grid [][]string) map[string]map[string]string {
	res := make(map[string]map[string]string)

	rows := len(grid)
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == "X" {
				continue
			}
			key := grid[i][j]
			res[key] = bfs(grid, i, j)
		}
	}
	return res
}

func bfs(grid [][]string, startX, startY int) map[string]string {
	rows := len(grid)
	cols := len(grid[0])

	type pnt struct {
		x, y int
	}

	// queue
	q := []pnt{}
	q = append(q, pnt{startX, startY})

	// inner map
	pathMap := make(map[string]string)
	pathMap[grid[startX][startY]] = ""

	// visited
	vis := make(map[string]bool)
	vis[grid[startX][startY]] = true

	for len(q) > 0 {
		curr := q[0]
		q = q[1:] // pop first element
		for _, dir := range dirs {
			nx, ny := curr.x+dir.Dx, curr.y+dir.Dy
			if nx >= 0 && nx < rows && ny >= 0 && ny < cols {
				visPnt := grid[nx][ny]
				if !vis[visPnt] {
					vis[visPnt] = true
					q = append(q, pnt{nx, ny})
					pathMap[visPnt] = pathMap[grid[curr.x][curr.y]] + dir.symbol
				}
			}

		}
	}
	return pathMap
}

func returnKeypad() [][]string {
	keypad := [][]string{
		{"7", "8", "9"},
		{"4", "5", "6"},
		{"1", "2", "3"},
		{"X", "0", "A"},
	}

	return keypad
}
func returnDirpad() [][]string {
	dirpad := [][]string{
		{"X", "^", "A"},
		{"<", "v", ">"},
	}

	return dirpad
}
