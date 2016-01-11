package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var molecule = "CRnCaCaCaSiRnBPTiMgArSiRnSiRnMgArSiRnCaFArTiTiBSiThFYCaFArCaCaSiThCaPBSiThSiThCaCaPTiRnPBSiThRnFArArCaCaSiThCaSiThSiRnMgArCaPTiBPRnFArSiThCaSiRnFArBCaSiRnCaPRnFArPMgYCaFArCaPTiTiTiBPBSiThCaPTiBPBSiRnFArBPBSiRnCaFArBPRnSiRnFArRnSiRnBFArCaFArCaCaCaSiThSiThCaCaPBPTiTiRnFArCaPTiBSiAlArPBCaCaCaCaCaSiRnMgArCaSiThFArThCaSiThCaSiRnCaFYCaSiRnFYFArFArCaSiRnFYFArCaSiRnBPMgArSiThPRnFArCaSiRnFArTiRnSiRnFYFArCaSiRnBFArCaSiRnTiMgArSiThCaSiThCaFArPRnFArSiRnFArTiTiTiTiBCaCaSiRnCaCaFYFArSiThCaPTiBPTiBCaSiThSiRnMgArCaF"
var combinations []string
var replacements = make(map[string]string)

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
			fmt.Println(split)
			mol = strings.Trim(split[0], " ")
			replace = strings.Trim(split[1], " ")
			replacements[mol] = replace
		}
	}
}

func replace() {
	// for i, v := range molecule {
	// 	split := strings.Split(molecule, "")
	// 	split[i] = replacements[string(v)]
	// 	newComb := strings.Join(split, "")
	// 	if !arrayutils.ContainsString(combinations, newComb) {
	// 		combinations = append(combinations, newComb)
	// 	}
	// }
	for k, v := range replacements {
		for i, v := range strings.Split(molecule, "") {
			if v == k {
				//replace...
			}
			// newComb := strings.Join(split, "")
			// if !arrayutils.ContainsString(combinations, newComb) {
			// 	combinations = append(combinations, newComb)
			// }
		}
	}
}

func main() {
	replace()
	fmt.Println("Maximum number of combinations:", len(combinations))
}
