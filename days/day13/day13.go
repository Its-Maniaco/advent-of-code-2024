package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
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
		fmt.Printf("Line %v\n\t\tA: %v,%v\n\t\tB:%v,%v\n\t\tPrize:%v,%v\n", i, aX, aY, bX, bY, pX, pY)
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

	cost := math.MaxInt
	// press A
	for cntA := 0; cntA <= prizex/ax && cntA <= prizey/ay && cntA < 100; cntA++ {
		// count B presses
		Xleft, Yleft := prizex-cntA*ax, prizey-cntA*ay
		if Xleft%bx == 0 && Yleft%by == 0 {
			cntB := Xleft / bx
			if cntB >= 0 && cntB <= 100 && cntB*by == Yleft {
				fmt.Printf("\t\tPress A: %v times & B: %v times\n", cntA, cntB)
				cost = int(math.Min(float64(cost), float64(cntA*3+cntB*1)))
			}
		}
	}

	if cost == math.MaxInt {
		return -1
	}
	return cost
}
