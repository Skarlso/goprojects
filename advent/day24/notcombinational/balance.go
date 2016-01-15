package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// var presents = []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
var presents []int
var presentGroups = make([][]int, 0)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content = bytes.TrimSpace(content)
	presents := convertToIntSlice(strings.Split(string(content), "\n"))
	// presents = arrayutils.ReverseInt(presents)
	//Part 1
	limit := sum(presents...) / 3
	fmt.Println(limit)
	//Part 2
	// limit := sum(presents...) / 4

	//TODO: Later on try here to just reverse the array and start pulling numbers together
	//Since the smallest product is needed and the smallest number of presents
	//I could start from the biggest number and if they give out the limit,
	//just store the product and remove those numbers from the list and go on
	//to the next numbers which give the same as LIMIT.
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
