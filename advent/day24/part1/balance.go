package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/ntns/goitertools/itertools"
	"github.com/skarlso/goutils/arrayutils"
)

// var presents = []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
var presentGroups = make([][]int, 0)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content = bytes.TrimSpace(content)
	presents := convertToIntSlice(strings.Split(string(content), "\n"))
	presents = arrayutils.ReverseInt(presents)
	//Part 1
	// limit := sum(presents...) / 3

	//Part 2
	limit := sum(presents...) / 4

	for i := 0; i < len(presents); i++ {
		//My recursive permuation was simply not fast enough
		presCombinations := itertools.Combinations(presents, i)
		for _, v := range presCombinations {
			if sum(v...) == limit {
				presentGroups = append(presentGroups, v)
			}

		}
	}

	var smallestPresentCountGroups [][]int
	smallestPresentCount := math.MaxInt64
	for _, v := range presentGroups {
		if len(v) <= smallestPresentCount {
			smallestPresentCount = len(v)
			smallestPresentCountGroups = append(smallestPresentCountGroups, v)
		}
	}

	if len(smallestPresentCountGroups) > 1 {
		smallestQe := math.MaxInt64
		for _, v := range smallestPresentCountGroups {
			vQe := getQuantumEntanglement(v)
			if vQe < smallestQe {
				smallestQe = vQe
			}
		}
		fmt.Println("Smallest qe", smallestQe)
	} else {
		fmt.Println("Smallest qe: ", getQuantumEntanglement(smallestPresentCountGroups[0]))
	}
}

func getQuantumEntanglement(in []int) int {
	qe := 1
	for _, v := range in {
		qe *= v
	}

	return qe
}

func convertToIntSlice(in []string) (out []int) {
	for _, v := range in {
		i, _ := strconv.Atoi(string(v))
		out = append(out, i)
	}

	return
}

func sum(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}
