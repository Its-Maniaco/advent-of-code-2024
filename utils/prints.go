package utils

import "fmt"

func PrintGrid(grid [][]string) {
	for _, row := range grid {
		for _, n := range row {
			fmt.Printf("%v ", n)
		}
		fmt.Println()
	}
}

func Print2DSlice(grid [][]int) {
	for _, row := range grid {
		for _, n := range row {
			fmt.Printf("%v ", n)
		}
		fmt.Println()
	}
}
