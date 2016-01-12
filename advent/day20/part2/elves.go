package main

import "fmt"

const (
	//LIMIT is the limit for a house's present. More than enough presents...
	LIMIT = 34000000
)

func main() {
	houses := make([]int, LIMIT)
	for i := 1; i < LIMIT; i++ {
		houseNum := 0
		for j := i; j < LIMIT; j += i {
			houseNum++
			if houseNum < 50 {
				houses[j] += i * 11
			}
		}
	}

	for i, v := range houses {
		if v >= LIMIT {
			fmt.Println(i)
			break
		}
	}
}
