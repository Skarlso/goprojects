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

var seatingCombinations = make([][]string, 0)
var table = make(map[string][]map[string]int)
var keys = make([]string, 0)

//Person a person
type Person struct {
	// neighbour *Person
	name string
	like int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		like, _ := strconv.Atoi(split[3]) //If lose -> * -1
		if split[2] == "lose" {
			like *= -1
		}
		table[split[0]] = append(table[split[0]], map[string]int{strings.Trim(split[10], "."): like})
		if !arrayutils.ContainsString(keys, split[0]) {
			keys = append(keys, split[0])
		}
	}
	generatePermutation(keys, len(keys))
	fmt.Println("Best seating efficiency:", calculateSeatingEfficiancy())
}

func generatePermutation(s []string, n int) {
	if n == 1 {
		news := make([]string, len(s))
		copy(news, s)
		seatingCombinations = append(seatingCombinations, news)
	}
	for i := 0; i < n; i++ {
		s[i], s[n-1] = s[n-1], s[i]
		generatePermutation(s, n-1)
		s[i], s[n-1] = s[n-1], s[i]
	}
}

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
			leftLike := getLikeForTargetConnect(v[i], v[left])
			rightLike := getLikeForTargetConnect(v[i], v[right])
			calculatedOrder += leftLike + rightLike
		}
		if calculatedOrder > bestSeating {
			bestSeating = calculatedOrder
		}
	}

	return bestSeating
}

func getLikeForTargetConnect(name string, neighbour string) int {
	neighbours := table[name]
	for _, t := range neighbours {
		if v, ok := t[neighbour]; ok {
			return v
		}
	}
	return 0
}
