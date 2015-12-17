package solutions

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func WrappingPaper() {
	//2*l*w + 2*w*h + 2*h*l
	//l*w*h
	p := make([]int, 3)
	areas := make([]int, 3)
	sum := 0

	inFile, _ := os.Open("solutions/wrapping_input.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		pString := strings.Split(scanner.Text(), "x")
		p[0], _ = strconv.Atoi(pString[0])
		p[1], _ = strconv.Atoi(pString[1])
		p[2], _ = strconv.Atoi(pString[2])
		areas[0] = p[0] * p[1]
		areas[1] = p[1] * p[2]
		areas[2] = p[0] * p[2]

		sum += (2 * areas[0]) + (2 * areas[1]) + (2 * areas[2])
		sort.Ints(areas)
		sum += areas[0]
	}
	fmt.Println("Sum:", sum)
}
