package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ajstarks/svgo"
)

//Location Defines a Santa whether he is human or not
type Location struct {
	x, y int
}

//HumanSanta it's a humanoid Santa
type HumanSanta struct {
	x, y int
}

//RobotSanta It's a robotic Santa
type RobotSanta struct {
	x, y int
}

//Move moves a human santa
func (hs *HumanSanta) Move(direction byte) {
	switch direction {
	case '^':
		hs.y++
	case 'v':
		hs.y--
	case '<':
		hs.x--
	case '>':
		hs.x++
	}
}

//Move moves a robot santa
func (rs *RobotSanta) Move(direction byte) {
	switch direction {
	case '^':
		rs.y++
	case 'v':
		rs.y--
	case '<':
		rs.x--
	case '>':
		rs.x++
	}
}

//DeliverPresents counts how many houses got presents this year
func main() {
	//Read in file input
	fileContent, err := ioutil.ReadFile("present_input.txt")
	if err != nil {
		panic(err)
	}

	f, _ := os.Create("day1.svg")
	defer f.Close()
	width := 120
	height := 120
	canvas := svg.New(f)
	canvas.Start(width, height)
	defer canvas.End()

	hs := &HumanSanta{30, 30}
	rs := &RobotSanta{30, 30}
	l := Location{30, 30}
	//Playing field
	var field = make(map[Location]bool, 0)
	field[l] = true
	//Moving Santa
	for i, v := range fileContent {
		if i&1 == 1 {
			origX := hs.x
			origY := hs.y
			hs.Move(v)
			canvas.Line(origX, origY, hs.x, hs.y, "fill: red; stroke: blue; stroke-width: 1")
			l = Location{hs.x, hs.y}
			field[l] = true
		} else {
			origX := rs.x
			origY := rs.y
			rs.Move(v)
			canvas.Line(origX, origY, rs.x, rs.y, "fill: blue; stroke: red; stroke-width: 1")
			l = Location{rs.x, rs.y}
			field[l] = true
		}
	}
	//All the houses which got visited will be true
	var housesVisited int
	for _, v := range field {
		if v {
			housesVisited++
		}
	}
	fmt.Println(housesVisited)
}
