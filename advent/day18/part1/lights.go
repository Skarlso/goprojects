package main

import (
	"fmt"
	"io/ioutil"
)

const (
	//GRIDX Maximum grid dimension X
	GRIDX = 100
	//GRIDY Maximum grid dimension Y
	GRIDY = 100
	//ON Defines a light which is on
	ON = 1
	//OFF Defines a light which is off
	OFF = 0
)

var lightGrid = make([][]int, GRIDX)

func init() {
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	for i := range lightGrid {
		lightGrid[i] = make([]int, GRIDY)
	}

	x := 0
	y := 0
	for _, v := range fileContent {
		if v == '#' {
			lightGrid[x][y] = 1
		}
		y = y + 1
		if v == '\n' {
			x = x + 1
			y = 0
		}
	}

}

func main() {
	fmt.Println("Grid Is Loaded:", lightGrid)
}
