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
	//ON Defines a light which is on
	ON = 1
	//OFF Defines a light which is off
	OFF = 0
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
		y = y + 1
		if v == '\n' {
			x = x + 1
			y = 0
		}
	}

}

//animate Animate the lights, Conway's Game of Life style
func animate() {
	for i := 0; i < LIMIT; i++ {
		for x := 0; x < GRIDX; x++ {
			for y := 0; y < GRIDY; y++ {
				animateLightAt(x, y)
			}
		}
	}
}

func animateLightAt(x, y int) {
	onCount := 0
	currentLight := lightGrid[x][y]
	for _, v := range corners {
		newX := x + v.x
		newY := y + v.y
		if (newX >= 0 && newX < GRIDX) && (newY >= 0 && newY < GRIDY) {
			if lightGrid[newX][newY] {
				onCount++
			}
		}
	}
	// fmt.Println("On Count:", onCount)
	if onCount == 3 && !currentLight {
		lightGrid[x][y] = true
	}

	if (onCount != 2 && onCount != 3) && currentLight {
		lightGrid[x][y] = false
	}

	// I leave this here because I liked it, but realized that wasn't that easy
	// if onCount >= 2 && onCount <= 3 {
	// 	lightGrid[x][y] &^= 1
	// }
	//
	// if onCount == 3 {
	// 	lightGrid[x][y] |= 1
	// }
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
	animate()
	fmt.Println("Lights which are on:", countOnLights())
}
