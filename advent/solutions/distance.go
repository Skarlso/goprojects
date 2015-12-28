package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Skarlso/goutils/arrayutils"
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
	file, _ := os.Open("solutions/distance_input2.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		distance, _ := strconv.Atoi(split[4])
		connections[split[0]] = append(connections[split[0]], Targets{split[2], distance})
		if !arrayutils.ContainsString(keys, split[0]) {
			keys = append(keys, split[0])
		}
		if !arrayutils.ContainsString(keys, split[2]) {
			keys = append(keys, split[2])
		}
	}
	fmt.Println(connections)
	generatePermutation(keys, len(keys))
	fmt.Println(routes)
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

//Gather connections based on the possible routes.
//Need another way to save the distance data in context to connected routes.
func getMinDistances() {
	mindis := 99999999
	fmt.Println(mindis)
}
