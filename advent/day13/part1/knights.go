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
	fmt.Println(seatingCombinations)
	fmt.Println("Connection for Alice -> Bob", getLikeForTargetConnect("Alice", "Bob"))
	fmt.Println("Connection for Alice -> David", getLikeForTargetConnect("Alice", "David"))
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

func getLikeForTargetConnect(name string, neighbour string) int {
	neighbours := table[name]
	for _, t := range neighbours {
		if v, ok := t[neighbour]; ok {
			return v
		}
	}
	return 0
}
