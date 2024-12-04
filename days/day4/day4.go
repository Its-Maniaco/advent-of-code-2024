package day4

import (
	"fmt"
	"log"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	err, lines := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	ans := 0

	// i index of line
	for i, line := range lines {
		var flag bool = false // to prevent retrigger for checkInLine for each X
		for j, ch := range line {
			if string(ch) == "X" {
				if flag == false {
					ans += checkInLine(line)
					flag = true
				}
				ans += checkVertical(lines, i, j)
				ans += checkDiagonal(lines, i, j)
			}
		}
	}
	fmt.Println(ans)
}

func checkInLine(line string) int {
	count := 0
	count += strings.Count(line, "XMAS")
	count += strings.Count(line, "SAMX")
	return count
}

func checkVertical(mat []string, r int, c int) int {
	count := 0
	var check string

	// above
	for k := r; k >= r-3 && k >= 0; k-- {
		check = check + string(mat[k][c])
	}
	count = increaseIfMatch(check, count)

	//below
	check = ""
	for k := r; k <= r+3 && k < len(mat); k++ {
		check = check + string(mat[k][c])
	}
	count = increaseIfMatch(check, count)

	return count
}

func checkDiagonal(mat []string, r int, c int) int {
	count := 0

	// top left
	check := ""
	for i, j := r, c; i >= 0 && j >= 0 && i >= r-3 && j >= c-3; i, j = i-1, j-1 {
		check = check + string(mat[i][j])
	}
	count = increaseIfMatch(check, count)

	// top right
	check = ""
	for i, j := r, c; i >= 0 && j < len(mat[0]) && i >= r-3 && j <= c+3; i, j = i-1, j+1 {
		check = check + string(mat[i][j])
	}
	count = increaseIfMatch(check, count)

	// bottom left
	check = ""
	for i, j := r, c; i < len(mat) && j >= 0 && i <= r+3 && j >= c-3; i, j = i+1, j-1 {
		check = check + string(mat[i][j])
	}
	count = increaseIfMatch(check, count)

	// bottom right
	check = ""
	for i, j := r, c; i < len(mat) && j < len(mat[0]) && i <= r+3 && j <= c+3; i, j = i+1, j+1 {
		check = check + string(mat[i][j])
	}
	count = increaseIfMatch(check, count)

	return count
}

func increaseIfMatch(s string, count int) int {
	if s == "XMAS" {
		return count + 1
	}
	return count
}
