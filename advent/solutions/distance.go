package solutions

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var connections = make(map[string][]Targets)
var routes = make(map[string]int)

//Targets locations which Santa would like to visit
type Targets struct {
	name     string
	distance int
	// visited bool
}

//CalculateDistance Calculate shortest distance for given routes
func CalculateDistance() {
	file, _ := os.Open("solutions/distance_input2.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		distance, _ := strconv.Atoi(split[4])
		connections[split[0]] = append(connections[split[0]], Targets{split[2], distance})
		// connections[split[2]] = append(connections[split[2]], Targets{split[0], distance})
	}
	generatePossibleRoutes()
}

// func generatePossibleConnectionsForLocation(loc string) {
// 	for _, l := range connections[loc] {
// 		if _, ok := connections[l.name]; ok {
// 			generatePossibleConnectionsForLocation(l.name)
// 		}
// 	}
// }

//Generates possibles routes between locations which are connected
func generatePossibleRoutes() {

}
