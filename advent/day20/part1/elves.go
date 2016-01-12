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
	elfNumber int
}

func main() {
	houseNumber := 1
	for {
		houses[houseNumber] = 0
		newElf := Elf{houseNumber}
		elves = append(elves, newElf)
		for _, e := range elves {
			if houseNumber%e.elfNumber == 0 {
				houses[houseNumber] += e.elfNumber * 10
				if houses[houseNumber] >= LIMIT {
					fmt.Println("Found house with limit: ", houseNumber)
					return
				}
			}
		}
		houseNumber++
	}
	// fmt.Println("houses:", houses)
	// fmt.Println("elves:", elves)

}

// func limitReached() bool {
// 	for k, v := range houses {
// 		if v >= LIMIT {
// 			fmt.Println("Limit reached, with key: ", k)
// 			return true
// 		}
// 	}
//
// 	return false
// }
