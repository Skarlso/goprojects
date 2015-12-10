package main

import (
	"fmt"
	"strconv"
)

var operationPermutation = [][]string{{"+", "-", " ", " ", "-", "+", "+", "-"},
	{"-", "+", " "},
	{"-", " ", "+"},
	{" ", "-", "+"},
	{"+", "-", "+"},
	{" ", " ", " "},
}

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func filterArray(op []string) []int {
	var returnInt []int
	var minuesedInt []int
	//Absolute value the ones which has negatives
	minuesedInt = append(minuesedInt, numbers[0])
	for i, v := range op {
		if v == "-" {
			minuesedInt = append(minuesedInt, -numbers[i+1])
		} else {
			minuesedInt = append(minuesedInt, numbers[i+1])
		}
	}
	// fmt.Println("Minused:", minuesedInt)
	skip := false
	index := 0
	//Filter out the spaces and bring the numbers together
	for _, v := range op {
		if skip {
			skip = false
			index++
			continue
		}
		if v == " " {
			num := strconv.Itoa(minuesedInt[index]) + strconv.Itoa(minuesedInt[index+1])
			intNum, _ := strconv.Atoi(num)
			returnInt = append(returnInt, intNum)
			skip = true
			index++
		} else {
			returnInt = append(returnInt, minuesedInt[index])
			index++
		}
	}

	//Add in leftovers
	if index < len(minuesedInt) {
		for i := index; i < len(minuesedInt); i++ {
			returnInt = append(returnInt, minuesedInt[i])
		}
	}
	// fmt.Println("return: ", returnInt)
	return returnInt
}

func doOperation() {
	var combinationsOfHundred [][]string
	//After filter there is only a sum needed here since our array is sanitised
	for _, v := range operationPermutation {
		sum := 0
		numberToWorkOn := filterArray(v)
		fmt.Println("NumToWorkOn:", numberToWorkOn)
		for _, v := range numberToWorkOn {
			sum += v
		}
		fmt.Println("sum:", sum)
		if sum == 37 {
			combinationsOfHundred = append(combinationsOfHundred, v)
		}
	}

	fmt.Println(combinationsOfHundred)
}

func main() {
	doOperation()
}
