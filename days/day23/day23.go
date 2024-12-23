package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/Its-Maniaco/advent-of-code-2024/utils"
)

func Part1(fileLoc string) {
	err, fs := utils.LineSlice(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	lan := buildMap(fs)

	validConns := connLanFind(lan)

	// add those connections which have a "t" in start
	director := [][]string{}
	for _, row := range validConns {
		for _, s := range row {
			if string(s[0]) == "t" {
				director = append(director, row)
				break
			}
		}
	}
	fmt.Println("Ans: ", len(director))
}

// TODO: try without 2 way keys
func buildMap(fs []string) map[string][]string {
	lan := make(map[string][]string, len(fs))
	for _, row := range fs {
		s := strings.Split(row, "-")
		lan[s[0]] = append(lan[s[0]], s[1])
		lan[s[1]] = append(lan[s[1]], s[0])
	}

	for k := range lan {
		sort.Strings(lan[k])
	}

	return lan
}

// get all valid 3 way connections
func connLanFind(lan map[string][]string) [][]string {
	lans := [][]string{}
	lansString := []string{}
	// k,v are base map key value
	for k, v := range lan {
		ss := []string{}
		// k2 is connection for a computer in base map
		for _, k2 := range v {
			// k3 is key that will be used to check if an computer that is present in lan[k2]
			// is also in k
			for _, k3 := range lan[k2] {
				if k3 == k {
					continue
				}
				// check if this key is also in k2
				for _, k1 := range v {
					if k2 == k1 {
						continue
					}
					// common connection fouind
					if k1 == k3 {
						ss = append(ss, k, k2, k3)
						sort.Strings(ss)
						ssString := connString(ss)
						if !checkConnStringExists(lansString, ssString) {
							lans = append(lans, ss)
							lansString = append(lansString, ssString)
						}
						ss = nil
					}
				}
			}
		}
	}

	return lans
}

// convert a connection into string
func connString(s []string) string {
	ss := fmt.Sprintf("%v-%v-%v", s[0], s[1], s[2])
	return ss
}

func checkConnStringExists(lans []string, s string) bool {
	for _, v := range lans {
		if s == v {
			return true
		}
	}
	return false
}
