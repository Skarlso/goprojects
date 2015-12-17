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
	for _, v := range fileContent {
		switch v {
		case '(':
			depth++
		case ')':
			depth--
		}

	}
	fmt.Println("Depth:", depth)
}
