package main

import "fmt"

var list1 = []string{"a", "b", "c"}
var list2 = []string{"1", "2", "3"}

func combine() {
	var comb []string
	for i, v := range list1 {
		comb = append(comb, v)
		comb = append(comb, list2[i])
	}
	fmt.Println(comb)
}

// func main() {
// 	combine()
// }
