package solutions

import (
	"fmt"
	"io/ioutil"
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
func DeliverPresents() {
	//Read in file input
	fileContent, err := ioutil.ReadFile("solutions/present_input.txt")
	if err != nil {
		panic(err)
	}

	hs := &HumanSanta{0, 0}
	rs := &RobotSanta{0, 0}
	l := Location{0, 0}
	//Playing field
	var field = make(map[Location]bool, 0)
	field[l] = true
	//Moving Santa
	for i, v := range fileContent {
		if i&1 == 1 {
			hs.Move(v)
			l = Location{hs.x, hs.y}
			field[l] = true
		} else {
			rs.Move(v)
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
