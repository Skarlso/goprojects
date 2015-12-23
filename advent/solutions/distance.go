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

	fmt.Println("Distance for AlphaCentauri and Tambi", getPathDistanceBetweenPoints([]string{"AlphaCentauri", "Tambi"}))
	generatePossibleRoutes()
}

//Retrieves distances for a route
//Route is represented by a sorted list
func getPathDistanceBetweenPoints(loc []string) int {
	dis := connections[loc[0]][0].distance
	return dis
}

//Generates possibles routes between locations which are connected
func generatePossibleRoutes() {
	//Lookup a location and then lookup that locations locations and connect it.
	//And then lookup that locations locations locations..... Recursive? Binary tree?
	//Begin small -> Locations for AlphaCentauri
	fmt.Println("Locations for AlphaCentauri:", connections["AlphaCentauri"])
	var possibleRoute []string
	possibleRoute = append(possibleRoute, "AlphaCentauri")
	for _, v := range connections["AlphaCentauri"] {
		fmt.Printf("Inner locations for: %s:%v\n", v.name, connections[v.name])
		if r, ok := connections[v.name]; ok {
			possibleRoute = append(possibleRoute, r[0].name)
		}
	}
	fmt.Println("Possible route combination:", possibleRoute)
}
