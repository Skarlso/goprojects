package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/skarlso/goutils/arrayutils"
)

var replacements = make(map[string][]string)
var endResult = "CRnCaCaCaSiRnBPTiMgArSiRnSiRnMgArSiRnCaFArTiTiBSiThFYCaFArCaCaSiThCaPBSiThSiThCaCaPTiRnPBSiThRnFArArCaCaSiThCaSiThSiRnMgArCaPTiBPRnFArSiThCaSiRnFArBCaSiRnCaPRnFArPMgYCaFArCaPTiTiTiBPBSiThCaPTiBPBSiRnFArBPBSiRnCaFArBPRnSiRnFArRnSiRnBFArCaFArCaCaCaSiThSiThCaCaPBPTiTiRnFArCaPTiBSiAlArPBCaCaCaCaCaSiRnMgArCaSiThFArThCaSiThCaSiRnCaFYCaSiRnFYFArFArCaSiRnFYFArCaSiRnBPMgArSiThPRnFArCaSiRnFArTiRnSiRnFYFArCaSiRnBFArCaSiRnTiMgArSiThCaSiThCaFArPRnFArSiRnFArTiTiTiTiBCaCaSiRnCaCaFYFArSiThCaPTiBPTiBCaSiThSiRnMgArCaF"

//init Loads in the strings from the input file
func init() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var mol string
		var replace string
		split := strings.Split(line, "=>")
		if len(split) > 1 {
			mol = strings.Trim(split[0], " ")
			replace = strings.Trim(split[1], " ")
			replacements[mol] = append(replacements[mol], replace)
		}
	}
}

//allIndiciesForString finds all the indexes for a given string
func allIndiciesForString(s, in string) (indicies []int) {
	index := strings.Index(in, s)
	offset := 0
	for index > -1 {
		indicies = append(indicies, index+offset)
		//Offset is there to save how long the string was before it was cut to tail
		offset += len(in[:index+len(s)])
		in = in[index+len(s):]
		index = strings.Index(in, s)
	}

	return
}

//replace does the replacing
func replace() {
	steps := 0
	keys := make([]string, 0, len(replacements))
	for k := range replacements {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	keys = arrayutils.ReverseString(keys)
	startingPoints := replacements["e"]
	for _, stp := range startingPoints {
		molecule := stp
		for _, k := range keys {
			if k == "e" {
				continue
			}

			for _, com := range replacements[k] {
				molecule = strings.Replace(molecule, k, com, -1)
				fmt.Println(molecule)
				if molecule == endResult {
					fmt.Println("Sortest:", steps)
				} else {
					steps++
				}
			}
		}

	}
	fmt.Println("Steps:", steps)
}

func main() {
	replace()
}
