package main

import (
	"bufio"
	"fmt"
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

var registers = make(map[string]int, 0)
var instructions = make([]Instruction, 0)

func init() {
	file, _ := os.Open("input.txt")
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
	registers["a"] = 1
	i := 0
	for i < len(instructions) {
		currentInst := instructions[i]
		switch currentInst.inst {
		case "jio":
			if registers[currentInst.register] == 1 {
				i += currentInst.offset
			} else {
				i++
			}
		case "jie":
			if registers[currentInst.register]&1 == 0 {
				i += currentInst.offset
			} else {
				i++
			}
		case "jmp":
			i += currentInst.offset
		case "tpl":
			registers[currentInst.register] *= 3
			i++
		case "inc":
			registers[currentInst.register]++
			i++
		case "hlf":
			registers[currentInst.register] /= 2
			i++
		}
	}

	fmt.Println("Value of register b is:", registers["b"])
}
