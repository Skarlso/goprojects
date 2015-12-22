package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var connections = make(map[string][]Targets)

//Targets locations which Santa would like to visit
type Targets struct {
	name     string
	distance int
}

//CalculateDistance Calculate shortest distance for given routes
func CalculateDistance() {
	file, _ := os.Open("solutions/distance_input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		distance, _ := strconv.Atoi(split[len(split)-1])
		connections[split[0]] = append(connections[split[0]], Targets{split[2], distance})
	}

	fmt.Println(connections)
}
