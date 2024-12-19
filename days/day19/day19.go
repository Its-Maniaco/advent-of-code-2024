package day19

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

	ptrn, dsgn := getPatternDesign(fs)
	if ptrn == nil || dsgn == nil {
		log.Fatal("Empty desgin or pattern.")
	}

	count := 0
	fmt.Println(dsgn)
	fmt.Println(ptrn)
	for _, row := range dsgn {
		if checkPossible(row, ptrn, "") {
			count++
		}
	}

	fmt.Println("Ans: ", count)
}

func checkPossible(s string, ptrn []string, tmp string) bool {
	if tmp == s {
		return true
	}

	for i := 0; i < len(ptrn); i++ {
		tmp2 := tmp + ptrn[i]
		if strings.HasPrefix(s, tmp2) {
			if checkPossible(s, ptrn, tmp2) {
				return true
			}
		}
	}
	return false
}

func getPatternDesign(fs []string) ([]string, []string) {
	ptrn := strings.Split(fs[0], ", ")
	dsn := []string{}
	for i := 2; i < len(fs); i++ {
		dsn = append(dsn, fs[i])
	}
	return ptrn, dsn
}
