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
		f := isNicePartTwo
		if f(scanner.Text()) {
			niceCounter++
			fmt.Println(scanner.Text())
		}
	}
	fmt.Println(niceCounter)
}

func isNicePartOne(s string) bool {
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

func isNicePartTwo(s string) bool {
	pairFound := false
	letterBetweenMatch := false

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) {
			pair := string(s[i]) + string(s[i+1])
			if strings.Contains(s[i+2:], pair) {
				pairFound = true
			}
		}
	}

	for i, v := range s {
		if i+2 < len(s) {
			if string(v) == string(s[i+2]) {
				letterBetweenMatch = true
			}
		}
	}

	return pairFound && letterBetweenMatch
}
