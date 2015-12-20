package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//LightCoordinates The location of lights which will be increased
type LightCoordinates struct {
	x, y int
}

var lightgridV2 = make(map[LightCoordinates]int, 1000)

//TurnOnTheLightsV2 turn on the lights
func TurnOnTheLightsV2() {
	inFile, _ := os.Open("solutions/lights_input.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		switches := strings.Split(scanner.Text(), " ")
		switch switches[0] {
		case "turn":
			if switches[1] == "on" {
				from := strings.Split(switches[2], ",")
				to := strings.Split(switches[4], ",")
				// fmt.Println("On:", from, to)
				turnOnOffV2(from, to, true)
			} else {
				from := strings.Split(switches[2], ",")
				to := strings.Split(switches[4], ",")
				// fmt.Println("Off:", from, to)
				turnOnOffV2(from, to, false)
			}
		case "toggle":
			from := strings.Split(switches[1], ",")
			to := strings.Split(switches[3], ",")
			// fmt.Println("toggle", from, to)
			toggleV2(from, to)
		}
	}

	countLightsV2()
}

func turnOnOffV2(from, to []string, on bool) {
	fromX, _ := strconv.Atoi(from[0])
	fromY, _ := strconv.Atoi(from[1])

	toX, _ := strconv.Atoi(to[0])
	toY, _ := strconv.Atoi(to[1])

	for i := fromX; i <= toX; i++ {
		for j := fromY; j <= toY; j++ {
			l := LightCoordinates{i, j}
			if on {
				lightgridV2[l]++
			} else {
				if lightgridV2[l] > 0 {
					lightgridV2[l]--
				}
			}
		}
	}
}

func toggleV2(from, to []string) {
	fromX, _ := strconv.Atoi(from[0])
	fromY, _ := strconv.Atoi(from[1])

	toX, _ := strconv.Atoi(to[0])
	toY, _ := strconv.Atoi(to[1])

	for i := fromX; i <= toX; i++ {
		for j := fromY; j <= toY; j++ {
			l := LightCoordinates{i, j}
			lightgridV2[l] += 2
		}
	}
}

func countLightsV2() {
	lightcount := 0
	for _, v := range lightgridV2 {
		lightcount += v
	}

	fmt.Println("Lights on:", lightcount)
}
