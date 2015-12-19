package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//TurnOnTheLights turn on the lights
func TurnOnTheLights() {
	inFile, _ := os.Open("solutions/lights_input.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	niceCounter := 0
	for scanner.Scan() {
		//scanner.Text()
		switches := strings.Split(scanner.Text(), " ")
		fmt.Println(switches)
	}
	fmt.Println(niceCounter)
}
