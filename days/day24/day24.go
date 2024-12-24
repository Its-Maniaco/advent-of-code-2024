package main

import (
	"fmt"
	"log"
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	err, fs := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	gates, ops := seperate(fs)

	// do all the operations
	for len(ops) > 0 {
		tmp := []string{}
		for _, row := range ops {
			t := parseOps(row)
			a, op, b, key := t[0], t[1], t[2], t[3]
			//fmt.Printf("%v %v %v %v\n", a, op, b, key)
			va, ok := gates[a]
			if !ok {
				tmp = append(tmp, row)
				continue
			}
			vb, ok := gates[b]
			if !ok {
				tmp = append(tmp, row)
				continue
			}

			if op == "AND" {
				gates[key] = va & vb
			} else if op == "OR" {
				gates[key] = va | vb
			} else {
				gates[key] = va ^ vb
			}
		}
		/*
			Wrong: copy does not allocate or reinitialize the destination slice;
			it assumes the destination slice already has sufficient capacity.

			ops = nil
			copy(ops, tmp)
		*/
		ops = append([]string{}, tmp...)

	}
	fmt.Println("Maps: ", gates)
	binZ := getKeyVals(gates)

	ans := convToDecimal(binZ)
	fmt.Println("Part 1: ", ans)
}

// parse input
func seperate(fs []string) (map[string]int, []string) {
	start := map[string]int{}
	operations := []string{}
	flag := false
	for _, row := range fs {
		if !flag {
			if row == "" {
				flag = true
				continue
			} else {
				tmp := strings.Split(row, ": ")
				k, vs := tmp[0], tmp[1]
				v, err := strconv.Atoi(vs)
				if err != nil {
					log.Println("Error parsing input: ", row)
					v = 0
				}
				start[k] = v
			}
		} else {
			operations = append(operations, row)
		}
	}
	return start, operations
}

// parse operations
func parseOps(row string) []string {
	t := strings.Split(row, " ")
	a, op, b, key := t[0], t[1], t[2], t[4]
	return []string{a, op, b, key}
}

// extract gate values starting with Z in sorted order
func getKeyVals(gates map[string]int) []int {
	// extract keys starting with "z"
	keys := []string{}
	for k := range gates {
		if strings.HasPrefix(k, "z") {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	vals := []int{}
	for _, k := range keys {
		vals = append(vals, gates[k])
	}
	return vals
}

func convToDecimal(vals []int) int {
	l := len(vals)
	n := 0
	slices.Reverse(vals)
	for i, v := range vals {
		n += v * int(math.Pow(2, float64(l-i-1)))
	}

	return n
}
