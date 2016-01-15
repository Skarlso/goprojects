package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/fighterlyt/permutation"
)

var presents []int

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	content = bytes.TrimSpace(content)
	presents := convertToIntSlice(strings.Split(string(content), "\n"))

	presentGroups, _ := permutation.NewPerm(presents, nil)
	// fmt.Println(presentGroups)
	for i, err := presentGroups.Next(); err == nil; i, err = presentGroups.Next() {
		fmt.Println(i.([]int))
	}
	// presentGroups := itertools.Permutations(presents, len(presents))
	// for _, v := range presentGroups {
	// 	fmt.Println(v)
	// }
}

func getQuantumEntanglement(in []int) (qe int) {
	for _, v := range in {
		qe *= v
	}

	return
}

func convertToIntSlice(in []string) (out []int) {
	for _, v := range in {
		i, _ := strconv.Atoi(string(v))
		out = append(out, i)
	}

	return
}
