package main

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
		//houseNumber and all of it's dividents * 10.
		// 6's dividents -> 6, 3, 2, 1 -> 6 + 3 + 2 + 1 -> 6 + 6 -> 12
		houseNumber++
	}
}

// func main() {
// 	houseNumber := 1
// 	done := make(chan bool)
// 	for {
// 		houses[houseNumber] = 0
// 		go visitHouses(houseNumber, done)
// 		houseNumber++
// 		if <-done {
// 			break
// 		}
// 	}
// 	// fmt.Println("houses:", houses)
// 	// fmt.Println("elves:", elves)
//
// }
//
// func visitHouses(houseNumber int, done chan bool) {
// 	newElf := Elf{houseNumber}
// 	elves = append(elves, newElf)
// 	for _, e := range elves {
// 		if houseNumber%e.elfNumber == 0 {
// 			houses[houseNumber] += e.elfNumber * 10
// 			// fmt.Println("House number and value:", houseNumber, houses[houseNumber])
// 			if houses[houseNumber] >= LIMIT {
// 				fmt.Println("Found house with limit: ", houseNumber)
// 				done <- true
// 			}
// 		}
// 	}
// 	done <- false
// }
