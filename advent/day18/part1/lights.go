package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ajstarks/svgo"
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

//Fill in the grid
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
	//Make a copy of the main grid
	innerGrid := make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		innerGrid[i] = make([]bool, len(grid))
	}

	//Switch on|off lights based on the given rules
	for i := range grid {
		for j := range grid[i] {
			innerGrid[i][j] = animateLightAt(grid, i, j)
		}
	}
	//Return the new grid with the correct values
	return innerGrid
}

//animateLightAt changes a light according to the game rules
func animateLightAt(grid [][]bool, x, y int) bool {
	onCount := 0
	currentLight := grid[x][y]

	//Collect the number of turned on lights around x,y.
	for _, v := range corners {
		newX := x + v.x
		newY := y + v.y
		if (newX >= 0 && newX < GRIDX) && (newY >= 0 && newY < GRIDY) {
			if grid[newX][newY] {
				onCount++
			}
		}
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

//countOnLights counts the 'on' state Lights
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

func drawState(n int, grid [][]bool) {
	f, _ := os.Create(fmt.Sprintf("step%d.svg", n))
	defer f.Close()
	width := 100
	height := 100
	canvas := svg.New(f)
	canvas.Start(width, height)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] {
				canvas.Rect(i, j, 10, 10, "fill:red;stroke:red;")
			} else {
				canvas.Rect(i, j, 10, 10, "fill:black;stroke:black;")
			}
		}
	}
	canvas.End()
}

//main main function
func main() {
	//Step the grid 'LIMIT' times
	for i := 0; i < LIMIT; i++ {
		lightGrid = animate(lightGrid)
		drawState(i, lightGrid)
	}
	fmt.Println("Lights which are on:", countOnLights())
}
