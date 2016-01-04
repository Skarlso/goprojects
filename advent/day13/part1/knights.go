package main

import (
	"bufio"
	"fmt"
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
	file, _ := os.Open("test_input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		like, _ := strconv.Atoi(split[3]) //If lose -> * -1
		if split[2] == "lose" {
			like *= -1
		}
		// table[split[0]] = append(table[split[0]], Person{strings.Trim(split[10], "."), like})
		table[split[0]] = append(table[split[0]], map[string]int{strings.Trim(split[10], "."): like})
		if !arrayutils.ContainsString(keys, split[0]) {
			keys = append(keys, split[0])
		}
	}

	fmt.Println(table)
	generatePermutation(keys, len(keys))
	// fmt.Println(seatingCombinations)
	// getMinDistances()
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

// func getMinDistances() {
// 	max := 0
// 	for _, v := range seatingCombinations {
// 		dis := 0
// 		for i := range v {
// 			//Check if the next item is in the connections of the first item.
// 			//If yes, retrieve its distance. That distance will be to THIS item.
// 			if i+1 < len(v) {
// 				dis += getLikeForTargetConnect(v[i], v[i+1])
// 			}
// 		}
// 		if dis > max {
// 			max = dis
// 		}
// 	}
// 	fmt.Println("Maximum like:", max)
// }
//
// func getLikeForTargetConnect(name string, neighbour string) int {
// 	neighbours := table[name]
// 	for _, t := range neighbours {
// 		if t.name == neighbour {
// 			return t.like
// 		}
// 	}
// 	return 0
// }
