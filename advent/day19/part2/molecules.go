package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	fmt.Println(replacements)
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

func replaceReq(s string, step int) {
	// fmt.Println("Current step:", s)
	for k := range replacements {
		indexes := allIndiciesForString(k, s)
		// fmt.Println(indexes)
		for _, i := range indexes {
			for _, rep := range replacements[k] {
				head := s[:i]
				tail := s[i+len(k):]
				s = head + rep + tail
				if s != endResult {
					fmt.Println("Current step:", s)
					step++
					replaceReq(s, step)
				} else {
					fmt.Println("Found: ", step)
					return
				}
			}
		}
	}
}

func main() {
	replaceReq("e", 0)
}
