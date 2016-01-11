package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/skarlso/goutils/arrayutils"
)

var molecule = "CRnCaCaCaSiRnBPTiMgArSiRnSiRnMgArSiRnCaFArTiTiBSiThFYCaFArCaCaSiThCaPBSiThSiThCaCaPTiRnPBSiThRnFArArCaCaSiThCaSiThSiRnMgArCaPTiBPRnFArSiThCaSiRnFArBCaSiRnCaPRnFArPMgYCaFArCaPTiTiTiBPBSiThCaPTiBPBSiRnFArBPBSiRnCaFArBPRnSiRnFArRnSiRnBFArCaFArCaCaCaSiThSiThCaCaPBPTiTiRnFArCaPTiBSiAlArPBCaCaCaCaCaSiRnMgArCaSiThFArThCaSiThCaSiRnCaFYCaSiRnFYFArFArCaSiRnFYFArCaSiRnBPMgArSiThPRnFArCaSiRnFArTiRnSiRnFYFArCaSiRnBFArCaSiRnTiMgArSiThCaSiThCaFArPRnFArSiRnFArTiTiTiTiBCaCaSiRnCaCaFYFArSiThCaPTiBPTiBCaSiThSiRnMgArCaF"
var combinations []string
var replacements = make(map[string][]string)

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

func allIndiciesForString(s, in string) (indicies []int) {
	index := strings.Index(in, s)
	offset := 0
	for index > -1 {
		indicies = append(indicies, index+offset)
		offset += len(in[:index]) + 1
		in = in[index+len(s):]
		index = strings.Index(in, s)
	}

	return
}

func replace() {
	for k, v := range replacements {
		indexes := allIndiciesForString(k, molecule)
		for _, i := range indexes {
			head := molecule[:i]
			tail := molecule[i+len(k):]
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
