package day05

import (
	"strings"
	"time"

	"github.com/alde/advent/2015/golang/shared"
)

func Solve() {
	shared.InputMustExist("./day05/input.txt", 2015, 5)
	result := shared.Result{
		Title: "Doesn't He Have Intern-Elves For This?",
		Day:   5,
	}

	result.Parts = []shared.Part{
		part1("./day05/input.txt"),
		part2("./day05/input.txt"),
	}

	shared.PrettyPrint(result)
}

func part1(input string) shared.Part {
	lines := shared.ReadLines(input)

	start := time.Now()
	result := 0

	for line := range lines {
		if isNice(line) {
			result += 1
		}
	}
	return shared.Part{
		Duration: time.Since(start),
		Result:   result,
	}
}

func part2(input string) shared.Part {
	lines := shared.ReadLines(input)

	start := time.Now()
	result := 0

	for line := range lines {
		if betterIsNice(line) {
			result += 1
		}
	}
	return shared.Part{
		Duration: time.Since(start),
		Result:   result,
	}
}

func isNice(input string) bool {
	vowelCount := 0
	hasRepeat := false
	for i, current := range input {
		switch current {
		case 'a':
			fallthrough
		case 'e':
			fallthrough
		case 'i':
			fallthrough
		case 'o':
			fallthrough
		case 'u':
			vowelCount += 1
		}

		// if we're at the end, don't keep checking - we'll go out of bounds on the next check
		if i >= len(input)-1 {
			break
		}

		peekAhead := rune(input[i+1])
		if current == 'a' && peekAhead == 'b' {
			return false
		}
		if current == 'c' && peekAhead == 'd' {
			return false
		}
		if current == 'p' && peekAhead == 'q' {
			return false
		}
		if current == 'x' && peekAhead == 'y' {
			return false
		}

		if current == rune(peekAhead) {
			hasRepeat = true
		}
	}

	return hasRepeat && vowelCount >= 3
}

func betterIsNice(input string) bool {
	hasRepeatedPair := false
	i := 1

	for i < len(input) {
		a, b := input[i-1], input[i]
		pair := string([]byte{a, b})
		if strings.Contains(input[i+1:], pair) {
			hasRepeatedPair = true
		}
		i++
	}

	hasRepeatWithPadding := false
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+2] {
			hasRepeatWithPadding = true
			break
		}
	}

	return hasRepeatedPair && hasRepeatWithPadding
}
