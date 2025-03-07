package day01

import (
	"log/slog"
	"os"
	"time"

	"github.com/alde/advent/2015/golang/shared"
)

func Solve() {
	result := shared.Result{
		Title: "Not Quite Lisp",
		Day:   1,
	}

	result.Parts = []shared.Part{
		part1("./day01/input.txt"),
		part2("./day01/input.txt"),
	}

	shared.PrettyPrint(result)

}

func part1(input string) shared.Part {

	lines, err := shared.ReadLines(input)
	if err != nil {
		slog.Error("failed to read file")
		os.Exit(1)
	}
	start := time.Now()

	result := 0
	for line := range lines {
		result += parse(line)
	}
	return shared.Part{
		Duration: time.Since(start),
		Result:   result,
	}
}

func part2(input string) shared.Part {

	lines, err := shared.ReadLines(input)
	if err != nil {
		slog.Error("failed to read file")
		os.Exit(1)
	}
	start := time.Now()

	result := 0
	for line := range lines {
		result += findBasement(line)
	}
	return shared.Part{
		Duration: time.Since(start),
		Result:   result,
	}
}

func parse(input string) int {
	moves := 0

	for _, r := range input {
		if r == '(' {
			moves += 1
		}
		if r == ')' {
			moves -= 1
		}
	}

	return moves
}

func findBasement(input string) int {
	floor := 0

	for i, r := range input {
		if r == '(' {
			floor += 1
		}
		if r == ')' {
			floor -= 1
		}
		if floor < 0 {
			return i + 1
		}
	}
	return -1
}
