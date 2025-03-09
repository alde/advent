package day07

import (
	"fmt"
	"strings"
	"time"

	"github.com/alde/advent/2015/golang/shared"
)

func Solve() {
	shared.InputMustExist("./day07/input.txt", 2015, 7)
	result := shared.Result{
		Title: "Probably a Fire Hazard",
		Day:   7,
	}

	result.Parts = []shared.Part{
		part1("./day07/input.txt"),
		// part2("./day07/input.txt"),
	}

	shared.PrettyPrint(result)
}

type pair struct {
	x, y int
}

func part1(input string) shared.Part {
	lines := shared.ReadLines(input)

	start := time.Now()

	for line := range lines {
		_ = line
	}
	return shared.Part{
		Duration: time.Since(start),
		Result:   0,
	}
}

func part2(input string) shared.Part {
	// lines := shared.ReadLines(input)

	start := time.Now()
	// grid := map[pair]int{}

	// for line := range lines {
	// 	processAdvancedInstructions(line, grid)
	// }
	return shared.Part{
		Duration: time.Since(start),
		Result:   0,
	}
}

func processInstruction(line string, registers map[string]uint16) {
	if strings.Contains(line, " AND ") {
		var left string
		var right string
		var output string
		fmt.Sscanf(line, "%s AND %s -> %s", &left, &right, &output)
		vleft := registers[left]
		vright := registers[right]
		registers[output] = vleft & vright

	} else if strings.Contains(line, " OR ") {
		var left string
		var right string
		var output string
		fmt.Sscanf(line, "%s OR %s -> %s", &left, &right, &output)
		vleft := registers[left]
		vright := registers[right]
		registers[output] = vleft | vright

	} else if strings.Contains(line, " LSHIFT ") {
		var left string
		var right uint16
		var output string
		fmt.Sscanf(line, "%s LSHIFT %d -> %s", &left, &right, &output)
		vleft := registers[left]
		registers[output] = vleft << 2

	} else if strings.Contains(line, "RSHIFT ") {
		var left string
		var right uint16
		var output string
		fmt.Sscanf(line, "%s RSHIFT %d -> %s", &left, &right, &output)
		vleft := registers[left]
		registers[output] = vleft >> 2
	} else if strings.Contains(line, "NOT ") {
		var left string
		var output string
		fmt.Sscanf(line, "NOT %s -> %s", &left, &output)
		vleft := registers[left]
		registers[output] = ^vleft
	} else {
		var left uint16
		var output string
		fmt.Sscanf(line, "%d -> %s", &left, &output)
		registers[output] = left
	}

}
