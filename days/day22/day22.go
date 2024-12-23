package main

import (
	"fmt"
	"log"
	"math"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	err, nums := utils.LineInt(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	// list of last required secrets for each
	finalSecrets := []int{}

	memo := make(map[int]int)
	for _, num := range nums {
		sec := num
		//newSecret := true
		for i := 1; i <= 2000; {
			sec, _ = nextSecret(sec, memo)
			// if !newSecret {
			// 	continue
			// }
			i++
		}
		fmt.Printf("Calculating secrets for: %v & last secret: %v\n", num, sec)
		//utils.AppendIntegerToFile("myoutput.txt", sec)
		finalSecrets = append(finalSecrets, sec)
	}
	//fmt.Println(finalSecrets)
	ans := addSecets(finalSecrets)
	fmt.Println("Ans: ", ans)
}

// returns next secret and flag if secret is unique till that point
func nextSecret(sec int, memo map[int]int) (int, bool) {
	tmp := sec
	//check if cached
	val, ok := memo[sec]
	if ok {
		return val, false
	}

	// step 1
	sec = secretProcess(sec, 64)

	//step 2
	sec = secretProcess(sec, 0.03125)

	//step 3
	sec = secretProcess(sec, 2048)

	memo[tmp] = sec
	return sec, true
}

func secretProcess(sec int, mul float64) int {
	// mul/divide step
	mul = mul * float64(sec)
	mul = math.Floor(float64(mul))
	intMul := int(mul)

	//mix
	sec = sec ^ intMul

	//prune
	sec = sec % 16777216

	return sec
}

func addSecets(secrets []int) int {
	ans := 0
	for _, num := range secrets {
		ans += (num)
	}

	return ans
}
