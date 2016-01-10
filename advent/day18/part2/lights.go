package main

import (
	"fmt"
	"io/ioutil"
)

//Location defines coordinates on a grid
type Location struct {
	x, y int
}

const (
	//GRIDX Maximum grid dimension X
	GRIDX = 100
	//GRIDY Maximum grid dimension Y
	GRIDY = 100
	//LIMIT defines maximum iteration count on grid
	LIMIT = 100
)

var (
	lightGrid = make([][]bool, GRIDX)
	corners   = []Location{
		{x: -1, y: -1},
		{x: 0, y: -1},
		{x: 1, y: -1},
		{x: 1, y: 0},
		{x: 1, y: 1},
		{x: 0, y: 1},
		{x: -1, y: 1},
		{x: -1, y: 0},
	}
)

func init() {
	fileContent, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	for i := range lightGrid {
		lightGrid[i] = make([]bool, GRIDY)
	}
	x := 0
	y := 0
	for _, v := range fileContent {
		if v == '#' {
			lightGrid[x][y] = true
		}
		y++
		if v == '\n' {
			x++
			y = 0
		}
	}

}

//animate Animate the lights, Conway's Game of Life style
//This is returning a grid because otherwise it was overwriting the wrong way.
//Since slices are mutable it was leaving my Grid in an inconsistante state.
func animate(grid [][]bool) [][]bool {
	innerGrid := make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		innerGrid[i] = make([]bool, len(grid))
	}

	for i := range grid {
		for j := range grid[i] {
			innerGrid[i][j] = animateLightAt(grid, i, j)
		}
	}
	return innerGrid
}

func animateLightAt(grid [][]bool, x, y int) bool {
	onCount := 0
	currentLight := grid[x][y]
	for _, v := range corners {
		newX := x + v.x
		newY := y + v.y
		if (newX >= 0 && newX < GRIDX) && (newY >= 0 && newY < GRIDY) {
			if grid[newX][newY] {
				onCount++
			}
		}
	}

	switch {
	case x == 0 && y == 0:
		return true
	case x == GRIDX-1 && y == GRIDY-1:
		return true
	case x == 0 && y == GRIDY-1:
		return true
	case x == GRIDX-1 && y == 0:
		return true
	}

	if currentLight {
		if onCount == 2 || onCount == 3 {
			return true
		}
	}

	if !currentLight {
		if onCount == 3 {
			return true
		}
	}
	return false
}

func countOnLights() (count int) {
	for i := 0; i < GRIDX; i++ {
		for j := 0; j < GRIDY; j++ {
			if lightGrid[i][j] {
				count++
			}
		}
	}

	return
}

func main() {
	for i := 0; i < LIMIT; i++ {
		lightGrid = animate(lightGrid)
	}
	fmt.Println("Lights which are on:", countOnLights())
}
