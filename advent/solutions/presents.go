package solutions

import (
	"fmt"
	"io/ioutil"
)

//location of a house
type location struct {
	x, y int
}

//DeliverPresents counts how many houses got presents this year
func DeliverPresents() {
	//Read in file input
	fileContent, err := ioutil.ReadFile("solutions/present_input.txt")
	if err != nil {
		panic(err)
	}

	//Playing field
	field := make(map[location]bool, 0)

	//First House init location
	l := location{0, 0}

	//Get a new field with the first house as starting point which got visited
	field[l] = true

	//Begin from 0 and then move around
	newX := 0
	newY := 0

	//Moving Santa
	for _, v := range fileContent {

		switch v {
		case '^':
			newY++
		case 'v':
			newY--
		case '<':
			newX--
		case '>':
			newX++
		}

		newLocation := location{newX, newY}
		field[newLocation] = true
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
