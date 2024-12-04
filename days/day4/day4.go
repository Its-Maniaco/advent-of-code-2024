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
				ans += checkVertical(lines, "XMAS", i, j)
				ans += checkDiagonal(lines, "XMAS", i, j)
			}
		}
	}
	fmt.Println(ans)
}

func Part2(fileLoc string) {
	err, lines := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	ans := 0

	for i, line := range lines {
		for j, ch := range line {
			if string(ch) == "M" {
				ans += checkDiagonal(lines, "MAS", i, j)
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

func checkVertical(mat []string, cmp string, r int, c int) int {
	count := 0
	var check string

	// above
	for k := r; k >= r-3 && k >= 0; k-- {
		check = check + string(mat[k][c])
	}
	count = increaseIfMatch(check, cmp, count)

	//below
	check = ""
	for k := r; k <= r+3 && k < len(mat); k++ {
		check = check + string(mat[k][c])
	}
	count = increaseIfMatch(check, cmp, count)

	return count
}

func checkDiagonal(mat []string, cmp string, r int, c int) int {
	count := 0

	count += checkDiagTopLeft(mat, cmp, r, c)
	count += checkDiagTopRight(mat, cmp, r, c)
	count += checkDiagBottomLeft(mat, cmp, r, c)
	count += checkDiagBottomRight(mat, cmp, r, c)

	return count
}

func checkDiagTopLeft(mat []string, cmp string, r int, c int) int {
	check := ""
	for i, j := r, c; i >= 0 && j >= 0 && i >= r-3 && j >= c-3; i, j = i-1, j-1 {
		check = check + string(mat[i][j])
	}
	return increaseIfMatch(check, cmp, 0)
}

func checkDiagTopRight(mat []string, cmp string, r int, c int) int {
	check := ""
	for i, j := r, c; i >= 0 && j < len(mat[0]) && i >= r-3 && j <= c+3; i, j = i-1, j+1 {
		check = check + string(mat[i][j])
	}
	return increaseIfMatch(check, cmp, 0)
}

func checkDiagBottomLeft(mat []string, cmp string, r int, c int) int {
	check := ""
	for i, j := r, c; i < len(mat) && j >= 0 && i <= r+3 && j >= c-3; i, j = i+1, j-1 {
		check = check + string(mat[i][j])
	}
	return increaseIfMatch(check, cmp, 0)
}

func checkDiagBottomRight(mat []string, cmp string, r int, c int) int {
	check := ""
	for i, j := r, c; i < len(mat) && j < len(mat[0]) && i <= r+3 && j <= c+3; i, j = i+1, j+1 {
		check = check + string(mat[i][j])
	}
	return increaseIfMatch(check, cmp, 0)
}

func increaseIfMatch(s string, cmp string, count int) int {
	if s == cmp {
		return count + 1
	}
	return count
}
