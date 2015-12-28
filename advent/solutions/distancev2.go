package solutions

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Node Represents a Location
type Node struct {
	distance   int
	name       string
	connection *Node
}

var nodeConnections = make(map[Node][]Node)

//CalculateDistanceV2 Calculate shortest distance for given routes
func CalculateDistanceV2() {
	file, _ := os.Open("solutions/distance_input2.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		distance, _ := strconv.Atoi(split[4])
		parent := Node{distance: distance, name: split[0]}
		connection := Node{distance: distance, name: split[2]}
		parent.connection = &connection
		connection.connection = &parent
		nodeConnections[parent] = append(nodeConnections[parent], connection)
	}
	// fmt.Println("Locations:", nodeConnections)
	traverseNodes()
}

//Generates possibles routes between locations which are connected
func traverseNodes() {
}
