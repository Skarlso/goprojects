package solutions

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/skarlso/goutils/arrayutils"
)

var connections = make(map[string][]Targets)
var routes = make([][]string, 0)
var keys = make([]string, 0)

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
		distance, _ := strconv.Atoi(split[4])
		//Adding both way of the connection
		connections[split[0]] = append(connections[split[0]], Targets{split[2], distance})
		connections[split[2]] = append(connections[split[2]], Targets{split[0], distance})
		if !arrayutils.ContainsString(keys, split[0]) {
			keys = append(keys, split[0])
		}
		if !arrayutils.ContainsString(keys, split[2]) {
			keys = append(keys, split[2])
		}
	}
	generatePermutation(keys, len(keys))
	getMinDistances()
}

func generatePermutation(s []string, n int) {
	if n == 1 {
		//Need this, because underneath slices all point to the same array.
		//So doing routes = append(routes, s) here would only result in the same
		//data.
		news := make([]string, len(s))
		copy(news, s)
		routes = append(routes, news)
	}
	for i := 0; i < n; i++ {
		s[i], s[n-1] = s[n-1], s[i]
		generatePermutation(s, n-1)
		s[i], s[n-1] = s[n-1], s[i]
	}
}

func getMinDistances() {
	min := math.MaxInt32
	for _, v := range routes {
		dis := 0
		for i := range v {
			//Check if the next item is in the connections of the first item.
			//If yes, retrieve its distance. That distance will be to THIS item.
			if i+1 < len(v) {
				dis += getDistanceForTargetConnect(v[i], v[i+1])
			}
		}
		if dis < min {
			min = dis
		}
	}
	fmt.Println("Minimum distance:", min)
}

func getDistanceForTargetConnect(name string, conn string) int {
	targets := connections[name]
	for _, t := range targets {
		if t.name == conn {
			return t.distance
		}
	}
	return 0
}
