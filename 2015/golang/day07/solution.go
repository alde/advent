package day07

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/alde/advent/2015/golang/shared"
)

func Solve() {
	shared.InputMustExist("./day07/input.txt", 2015, 7)
	result := shared.Result{
		Title: "Some Assembly Required",
		Day:   7,
	}

	result.Parts = []shared.Part{
		part1("./day07/input.txt"),
		part2("./day07/input.txt"),
	}

	shared.PrettyPrint(result)
}

func part1(input string) shared.Part {
	lines := shared.ReadLines(input)

	start := time.Now()
	registers := map[string]uint16{}
	wires := prepareWires(lines)

	for output := range wires {
		emulateCircuit(output, registers, wires)
	}

	return shared.Part{
		Duration: time.Since(start),
		Result:   int(registers["a"]),
	}
}

func part2(input string) shared.Part {
	lines := shared.ReadLines(input)

	start := time.Now()
	registers := map[string]uint16{}
	wires := prepareWires(lines, "a -> b")

	for output := range wires {
		emulateCircuit(output, registers, wires)
	}

	return shared.Part{
		Duration: time.Since(start),
		Result:   int(registers["a"]),
	}
}

func prepareWires(lines shared.LineGenerator, override ...string) map[string]string {
	wires := map[string]string{}
	for line := range lines {
		if len(override) > 0 && strings.HasSuffix(line, " -> b") {
			line = "3176 -> b"
		}
		parts := strings.Split(line, " -> ")
		output := parts[1]
		wires[output] = line
	}
	return wires
}

func emulateCircuit(input string, registers map[string]uint16, instructions map[string]string) uint16 {
	if value, ok := registers[input]; ok {
		return value
	}
	if value, err := strconv.Atoi(input); err == nil {
		return uint16(value)
	}
	if instruction, ok := instructions[input]; ok {
		processInstruction(instruction, registers, instructions)
		return registers[input]
	}
	return 0
}

func processInstruction(line string, registers map[string]uint16, instructions map[string]string) {
	var left, right, output string
	if strings.Contains(line, " AND ") {
		fmt.Sscanf(line, "%s AND %s -> %s", &left, &right, &output)
		registers[output] = emulateCircuit(left, registers, instructions) & emulateCircuit(right, registers, instructions)
	} else if strings.Contains(line, " OR ") {
		var right string
		fmt.Sscanf(line, "%s OR %s -> %s", &left, &right, &output)
		registers[output] = emulateCircuit(left, registers, instructions) | emulateCircuit(right, registers, instructions)
	} else if strings.Contains(line, " LSHIFT ") {
		var right uint16
		fmt.Sscanf(line, "%s LSHIFT %d -> %s", &left, &right, &output)
		registers[output] = emulateCircuit(left, registers, instructions) << right
	} else if strings.Contains(line, " RSHIFT ") {
		var right uint16
		fmt.Sscanf(line, "%s RSHIFT %d -> %s", &left, &right, &output)
		registers[output] = emulateCircuit(left, registers, instructions) >> right
	} else if strings.Contains(line, "NOT ") {
		fmt.Sscanf(line, "NOT %s -> %s", &left, &output)
		registers[output] = ^emulateCircuit(left, registers, instructions)
	} else {
		fmt.Sscanf(line, "%s -> %s", &left, &output)
		registers[output] = emulateCircuit(left, registers, instructions)
	}
}
