package main

import (
	"fmt"
	"sort"
)

type sortableInt []int

func (s sortableInt) Len() int {
	return len(s)
}
func (s sortableInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortableInt) Less(i, j int) bool {
	return s[i] > s[j]
}

func arrange() {
	toSort := []int{999, 23, 322, 42, 500}
	// var biggestNumber string
	sort.Sort(sortableInt(toSort))
	// biggestNumber = strings.Join(list4, "")
	fmt.Println(toSort)
}

//
// func main() {
// 	arrange()
// }
