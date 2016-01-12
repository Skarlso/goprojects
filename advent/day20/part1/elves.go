package main

import "fmt"

const (
	//LIMIT is the limit for a house's present. More than enough presents...
	LIMIT = 34000000
)

func main() {
	houses := make([]int, LIMIT/10)
	for i := 1; i < LIMIT/10; i++ {
		for j := i; j < LIMIT/10; j += i {
			houses[j] += i * 10
		}
	}

	for i, v := range houses {
		if v >= LIMIT {
			fmt.Println(i)
			break
		}
	}
}
