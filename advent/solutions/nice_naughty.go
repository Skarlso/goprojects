package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//NiceOrNaughty returns how many nice strings there are in a file
func NiceOrNaughty() {
	inFile, _ := os.Open("solutions/nice_naughty_input.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	niceCounter := 0
	for scanner.Scan() {
		if isNice(scanner.Text()) {
			niceCounter++
		}
	}
	fmt.Println(niceCounter)
}

func isNice(s string) bool {
	//Negative lookahead is not supported in Go
	//Backreference is not supported either...
	//Which leaves me with cycles.
	naughty := []string{"ab", "cd", "pq", "xy"}
	for _, v := range naughty {
		if strings.Contains(s, v) {
			return false
		}
	}

	var nice = "aeiou"
	vowelcount := 0
	var previousChar rune
	doubleChar := false
	for _, v := range s {
		if strings.Contains(nice, string(v)) {
			vowelcount++
		}
		if previousChar == v {
			doubleChar = true
		}
		previousChar = v
	}
	if vowelcount < 3 || !doubleChar {
		return false
	}

	return true
}
