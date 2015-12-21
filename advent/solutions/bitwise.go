package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//Equation represents an Equation
//Will be stored in a map for which the key is the result.
type Equation struct {
	operands  []string
	operation string
	result    int
	hasValue  bool
}

var equations = make(map[string]*Equation, 0)
var circuit = make([]string, 0)

//BobbyTable Create help for wiring.
func BobbyTable() {

	file, _ := os.Open("solutions/bitwise_input2.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		circuit = append(circuit, line)
		split := strings.Split(line, " ")
		var operands []string
		var operation string
		switch len(split) {
		case 3:
			hasValue := false
			res, err := strconv.Atoi(split[0])
			if err == nil {
				hasValue = true
			}
			operands = append(operands, split[0])
			equations[split[len(split)-1]] = &Equation{operands: operands, operation: "SET", result: res, hasValue: hasValue}
		case 4:
			operands = append(operands, split[1])
			operation = split[0]
			equations[split[len(split)-1]] = &Equation{operands, operation, 0, false}
		case 5:
			operands = append(operands, split[0])
			operands = append(operands, split[2])
			operation = split[1]
			equations[split[len(split)-1]] = &Equation{operands, operation, 0, false}
		}

	}

	result := value(equations["a"])
	fmt.Println("a:", result)
}

func value(in *Equation) (result int) {
	if in.hasValue {
		return in.result
	}
	switch in.operation {
	case "SET":
		result = value(equations[in.operands[0]])
	case "AND":
		op1 := 0
		op2 := 0
		if v, err := strconv.Atoi(in.operands[0]); err == nil {
			op1 = v
		} else {
			op1 = value(equations[in.operands[0]])
		}
		if v, err := strconv.Atoi(in.operands[1]); err == nil {
			op2 = v
		} else {
			op2 = value(equations[in.operands[1]])
		}
		result = op1 & op2
	case "OR":
		result = value(equations[in.operands[0]]) | value(equations[in.operands[1]])
	case "NOT":
		result = ^value(equations[in.operands[0]])
	case "LSHIFT":
		shiftBy, _ := strconv.Atoi(in.operands[1])
		result = value(equations[in.operands[0]]) << uint(shiftBy)
	case "RSHIFT":
		shiftBy, _ := strconv.Atoi(in.operands[1])
		result = value(equations[in.operands[0]]) >> uint(shiftBy)
	}

	in.hasValue = true
	in.result = result
	return
}
