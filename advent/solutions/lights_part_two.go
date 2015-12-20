package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	//ON on
	ON = iota
	//OFF off
	OFF
	//TOGGLE toggle
	TOGGLE
)

var lightgridV2 = make([][]int, GRIDX)

//TurnOnTheLightsV2 turn on the lights
func TurnOnTheLightsV2() {
	inFile, _ := os.Open("solutions/lights_input.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for i := range lightgridV2 {
		lightgridV2[i] = make([]int, GRIDY)
	}

	for scanner.Scan() {
		switches := strings.Split(scanner.Text(), " ")
		switch switches[0] {
		case "turn":
			if switches[1] == "on" {
				from := strings.Split(switches[2], ",")
				to := strings.Split(switches[4], ",")
				// fmt.Println("On:", from, to)
				handleLightV2(from, to, ON)
			} else {
				from := strings.Split(switches[2], ",")
				to := strings.Split(switches[4], ",")
				// fmt.Println("Off:", from, to)
				handleLightV2(from, to, OFF)
			}
		case "toggle":
			from := strings.Split(switches[1], ",")
			to := strings.Split(switches[3], ",")
			// fmt.Println("toggle", from, to)
			handleLightV2(from, to, TOGGLE)
		}
	}

	countLightsV2()
}

func handleLightV2(from, to []string, action int) {
	fromX, _ := strconv.Atoi(from[0])
	fromY, _ := strconv.Atoi(from[1])

	toX, _ := strconv.Atoi(to[0])
	toY, _ := strconv.Atoi(to[1])

	for i := fromX; i <= toX; i++ {
		for j := fromY; j <= toY; j++ {

			switch action {
			case ON:
				lightgridV2[i][j]++
			case OFF:
				if lightgridV2[i][j] > 0 {
					lightgridV2[i][j]--
				}
			case TOGGLE:
				lightgridV2[i][j] += 2
			}
		}
	}
}

func countLightsV2() {
	lightcount := 0
	for i := 0; i < len(lightgridV2); i++ {
		for j := 0; j < len(lightgridV2[i]); j++ {
			lightcount += lightgridV2[i][j]
		}
	}

	fmt.Println("Lights on:", lightcount)
}
