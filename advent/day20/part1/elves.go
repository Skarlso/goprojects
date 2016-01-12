package main

import "fmt"

const (
	//LIMIT is the limit for a house's present. More than enough presents...
	LIMIT = 34000000
)

var houses = make(map[int]int, 0)
var elves = make([]Elf, 0)

//Elf represents either a pointed eared wise a** or a little fellow delivering presents
type Elf struct {
	multiplier int
}

func main() {
	i := 1
	for {
		newElf := Elf{i}
		elves = append(elves, newElf)
		for _, e := range elves {
			houses[e.multiplier] += e.multiplier * 10
		}
		i++

		if limitReached() {
			break
		}
	}

}

func limitReached() bool {
	for k, v := range houses {
		if v == LIMIT {
			fmt.Println("Limit reached, with key: ", k)
			return true
		}
	}

	return false
}
