package day02

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alde/advent/2015/golang/shared"
)

const Title = "I Was Told There Would Be No Math"

func Solve() {
	result := shared.Result{
		Title: Title,
		Day:   2,
	}

	result.Parts = []shared.Part{
		part1("./day02/input.txt"),
		part2("./day02/input.txt"),
	}

	shared.PrettyPrint(result)

}

func parse(input string) (int, int, int) {
	ints := [3]int{}
	parts := strings.Split(input, "x")
	for i, p := range parts {
		ints[i], _ = strconv.Atoi(p)
	}
	return ints[0], ints[1], ints[2]
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
		result += wrappingPaper(parse(line))
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
		result += ribbon(parse(line))
	}
	return shared.Part{
		Duration: time.Since(start),
		Result:   result,
	}
}

func wrappingPaper(l, w, h int) int {
	a := l * w
	b := w * h
	c := h * l

	return 2*a + 2*b + 2*c + min(a, b, c)
}

func ribbon(l, w, h int) int {
	smallestSide := 2*l + 2*w
	if 2*w+2*h < smallestSide {
		smallestSide = 2*w + 2*h
	}
	if 2*l+2*h < smallestSide {
		smallestSide = 2*l + 2*h
	}

	return smallestSide + (l * w * h)
}

func min(i ...int) int {
	smallest := 10000
	for _, c := range i {
		if c < smallest {
			smallest = c
		}
	}

	return smallest
}
