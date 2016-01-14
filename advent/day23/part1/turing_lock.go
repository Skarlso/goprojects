package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Instruction defines a turing instruction
type Instruction struct {
	inst     string
	register string
	offset   int
}

var instructions = make([]Instruction, 0)

func init() {
	file, _ := os.Open("test_input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		//Hardcoded instruction handling for greater justice.
		var i Instruction
		switch split[0] {
		case "jio", "jie":
			register := strings.Trim(split[1], ",")
			offset, _ := strconv.Atoi(split[2])
			i = Instruction{inst: split[0], register: register, offset: offset}
		case "jmp":
			offset, _ := strconv.Atoi(split[1])
			i = Instruction{inst: split[0], offset: offset}
		case "tpl", "inc", "hlf":
			register := strings.Trim(split[1], ",")
			i = Instruction{inst: split[0], register: register}
		}
		instructions = append(instructions, i)
	}
}

func main() {

}
