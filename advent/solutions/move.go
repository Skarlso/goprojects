package solutions

import (
	"fmt"
	"io/ioutil"
)

//MoveSanta Moves santa to a designated floor by reading the input file
func MoveSanta() {
	fileContent, err := ioutil.ReadFile("solutions/input.txt")
	if err != nil {
		panic(err)
	}
	depth := 0
	firstFound := false
	firstChar := -1
	for i, v := range fileContent {
		switch v {
		case '(':
			depth++
		case ')':
			depth--
		}
		if depth < 0 && !firstFound {
			firstChar = i + 1
			firstFound = true
		}
	}
	fmt.Println("Depth:", depth)
	fmt.Println("First Basement:", firstChar)
}
