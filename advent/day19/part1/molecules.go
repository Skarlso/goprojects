package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/skarlso/goutils/arrayutils"
)

var molecule = "CRnCaCaCaSiRnBPTiMgArSiRnSiRnMgArSiRnCaFArTiTiBSiThFYCaFArCaCaSiThCaPBSiThSiThCaCaPTiRnPBSiThRnFArArCaCaSiThCaSiThSiRnMgArCaPTiBPRnFArSiThCaSiRnFArBCaSiRnCaPRnFArPMgYCaFArCaPTiTiTiBPBSiThCaPTiBPBSiRnFArBPBSiRnCaFArBPRnSiRnFArRnSiRnBFArCaFArCaCaCaSiThSiThCaCaPBPTiTiRnFArCaPTiBSiAlArPBCaCaCaCaCaSiRnMgArCaSiThFArThCaSiThCaSiRnCaFYCaSiRnFYFArFArCaSiRnFYFArCaSiRnBPMgArSiThPRnFArCaSiRnFArTiRnSiRnFYFArCaSiRnBFArCaSiRnTiMgArSiThCaSiThCaFArPRnFArSiRnFArTiTiTiTiBCaCaSiRnCaCaFYFArSiThCaPTiBPTiBCaSiThSiRnMgArCaF"

// var molecule = "HOHOHOHO"
var combinations []string
var replacements = make(map[string][]string)

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

//replace does the replacing
func replace() {
	for k, v := range replacements {
		//Get all the indexes for a Key
		indexes := allIndiciesForString(k, molecule)
		for _, i := range indexes {
			//Save the head up to the index
			head := molecule[:i]
			//Save the tail from the index + lenght of the searched key
			tail := molecule[i+len(k):]

			//Create a string for all the replacement possbilities
			for _, com := range v {
				newMol := head + com + tail
				if !arrayutils.ContainsString(combinations, newMol) {
					combinations = append(combinations, newMol)
				}
			}
		}
	}
}

func main() {
	replace()
	fmt.Println("Maximum number of combinations:", len(combinations))
}
