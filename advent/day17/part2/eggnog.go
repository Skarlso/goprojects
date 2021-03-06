package main

import (
	"fmt"

	"github.com/Skarlso/goitertools/itertools"
)

//LIMIT the liter limit we are looking for
const LIMIT = 150

var input = []int{33, 14, 18, 20, 45, 35, 16, 35, 1, 13, 18, 13, 50, 44, 48, 6, 24, 41, 30, 42}
var combinationCount int
var combinationsAndContainerCount = make(map[int]int)

func main() {
	for i := 0; i < len(input); i++ {
		//My recursive permuation was simply not fast enough
		contCombinations := itertools.Combinations(input, i)
		for _, v := range contCombinations {
			if sum(v...) == LIMIT {
				combinationsAndContainerCount[len(v)]++
				combinationCount++
			}
		}
	}

	fmt.Println("Combination count:", combinationCount)
	fmt.Println("Smallest Number of Containers which add up to LIMIT:", combinationsAndContainerCount)
}

func sum(nums ...int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}
