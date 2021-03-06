package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	//GRIDX Maximum grid dimension X
	GRIDX = 1000
	//GRIDY Maximum grid dimension Y
	GRIDY = 1000
)

var lightgrid = make([][]bool, GRIDX)

//TurnOnTheLights turn on the lights
func TurnOnTheLights() {
	inFile, _ := os.Open("solutions/lights_input.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for i := range lightgrid {
		lightgrid[i] = make([]bool, GRIDY)
	}

	for scanner.Scan() {
		switches := strings.Split(scanner.Text(), " ")
		switch switches[0] {
		case "turn":
			if switches[1] == "on" {
				from := strings.Split(switches[2], ",")
				to := strings.Split(switches[4], ",")
				// fmt.Println("On:", from, to)
				handleLight(from, to, ON)
			} else {
				from := strings.Split(switches[2], ",")
				to := strings.Split(switches[4], ",")
				// fmt.Println("Off:", from, to)
				handleLight(from, to, OFF)
			}
		case "toggle":
			from := strings.Split(switches[1], ",")
			to := strings.Split(switches[3], ",")
			// fmt.Println("toggle", from, to)
			handleLight(from, to, TOGGLE)
		}
	}

	countLights()
}

func handleLight(from, to []string, action int) {
	fromX, _ := strconv.Atoi(from[0])
	fromY, _ := strconv.Atoi(from[1])

	toX, _ := strconv.Atoi(to[0])
	toY, _ := strconv.Atoi(to[1])

	for i := fromX; i <= toX; i++ {
		for j := fromY; j <= toY; j++ {

			switch action {
			case ON:
				lightgrid[i][j] = true
			case OFF:
				lightgrid[i][j] = false
			case TOGGLE:
				lightgrid[i][j] = !lightgrid[i][j]
			}
		}
	}
}

func countLights() {
	lightcount := 0
	for i := 0; i < len(lightgrid); i++ {
		for j := 0; j < len(lightgrid[i]); j++ {
			if lightgrid[i][j] {
				lightcount++
			}
		}
	}

	fmt.Println("Lights on:", lightcount)
}
