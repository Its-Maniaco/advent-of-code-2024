package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part(fileLoc string, flag bool) {
	err, fs := utils.LineSlice(fileLoc)
	if err != nil {
		fmt.Println(err)
	}
	totalCost := 0

	// Get cost and prize
	for i := 0; i <= len(fs)-3; i = i + 4 {
		// cost of A
		t := strings.Split(fs[i], ", ")
		aX, _ := strconv.Atoi(t[0][11:])
		aY, _ := strconv.Atoi(t[1][1:])
		//cost of B
		t = strings.Split(fs[i+1], ", ")
		bX, _ := strconv.Atoi(t[0][11:])
		bY, _ := strconv.Atoi(t[1][1:])
		//prize
		t = strings.Split(fs[i+2], ", ")
		pX, _ := strconv.Atoi(t[0][9:])
		pY, _ := strconv.Atoi(t[1][2:])
		if flag {
			pX += 10000000000000
			pY += 10000000000000
		}
		//fmt.Printf("Line %v\n\t\tA: %v,%v\n\t\tB:%v,%v\n\t\tPrize:%v,%v\n", i, aX, aY, bX, bY, pX, pY)
		cost := calcTokens([2]int{aX, aY}, [2]int{bX, bY}, [2]int{pX, pY})
		if cost != -1 {
			totalCost += cost
		}
	}

	fmt.Println("Minimum Cost: ", totalCost)
}

func calcTokens(a, b, prize [2]int) int {
	ax, ay := a[0], a[1]
	bx, by := b[0], b[1]
	prizex, prizey := prize[0], prize[1]

	/*
		ax*cntA + bx*cntB = prizex
		ay*cntA + by*cntB = prizey
	*/

	cntA := float64(((by * prizex) - (bx * prizey))) / float64((ax*by)-(ay*bx))
	// if not whole number
	if math.Round(cntA) != cntA {
		return -1
	}
	cntB := float64(prizey-(ay*int(cntA))) / float64(by)
	if math.Round(cntB) != cntB {
		return -1
	}
	cost := int(3*cntA + cntB)
	return cost
}
