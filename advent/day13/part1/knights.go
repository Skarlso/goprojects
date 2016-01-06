package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/skarlso/goutils/arrayutils"
)

//Permutated variations on seating orders
var seatingCombinations = make([][]int, 0)

//Representation of the Table. This will be handled as a circular slice
var table = make(map[int][]map[int]int)

//A list of separate keys which will be used for the permutations
var keys = make([]int, 0)

//A mapping between names and 'int' based IDs
var nameMapping = make(map[string]int)

//Person representation of a Person who has a like and name
type Person struct {
	// neighbour *Person
	name string
	like int
}

func main() {
	CalculatePerfectSeating()
}

//CalculatePerfectSeating returns the perfect seating order based on Love/Hate relations
func CalculatePerfectSeating() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	id := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		trimmedNeighbour := strings.Trim(split[10], ".")
		like, _ := strconv.Atoi(split[3]) //If lose -> * -1
		//Save the mappings of names between IDs. We need to check if it exists already
		//because if so, do not increment the ID.
		if _, ok := nameMapping[split[0]]; !ok {
			nameMapping[split[0]] = id
			id++
		}
		//We need to save the ID of the neigbour as well
		if _, ok := nameMapping[trimmedNeighbour]; !ok {
			nameMapping[trimmedNeighbour] = id
			id++
		}
		if split[2] == "lose" {
			like *= -1
		}
		//Save the connections that A Single person can have. This will be used to look up like relation.
		table[nameMapping[split[0]]] = append(table[nameMapping[split[0]]], map[int]int{nameMapping[trimmedNeighbour]: like})

		//Save the keys separatly so we can permutate them.
		if !arrayutils.ContainsInt(keys, nameMapping[split[0]]) {
			keys = append(keys, nameMapping[split[0]])
		}
	}
	generatePermutation(keys, len(keys))
	fmt.Println("Best seating efficiency:", calculateSeatingEfficiancy())
}

//GeneratePermutation generate possible permutations of the IDs which correspond to names
func generatePermutation(s []int, n int) {
	if n == 1 {
		news := make([]int, len(s))
		copy(news, s)
		seatingCombinations = append(seatingCombinations, news)
	}
	for i := 0; i < n; i++ {
		s[i], s[n-1] = s[n-1], s[i]
		generatePermutation(s, n-1)
		s[i], s[n-1] = s[n-1], s[i]
	}
}

//CalculateSeatingEfficiancy will return the best seating number
func calculateSeatingEfficiancy() int {
	bestSeating := math.MinInt64
	for _, v := range seatingCombinations {
		calculatedOrder := 0

		for i := range v {
			left := (i - 1) % len(v)
			//This is to work around the fact that in Go
			//modulo of a negative number will not return a positive number.
			//So -1 % 4 will not return 3 but -1. In that case we add length.
			if left < 0 {
				left += len(v)
			}
			right := (i + 1) % len(v)
			//Basically the first element will have a neighbour which is the last and the last's
			//neighbour is the first.
			calculatedOrder += getLikeForNeighbour(v[i], v[left]) + getLikeForNeighbour(v[i], v[right])
		}
		if calculatedOrder > bestSeating {
			bestSeating = calculatedOrder
		}
	}

	return bestSeating
}

//GetLikeForNeighbour Gets the like number for a neighbour for an ID
func getLikeForNeighbour(name int, neighbour int) int {
	neighbours := table[name]
	for _, t := range neighbours {
		if v, ok := t[neighbour]; ok {
			return v
		}
	}
	return 0
}
