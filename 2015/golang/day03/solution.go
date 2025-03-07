package day03

import (
	"time"

	"github.com/alde/advent/2015/golang/shared"
)

func Solve() {
	result := shared.Result{
		Title: "Perfectly Spherical Houses in a Vacuum",
		Day:   3,
	}

	result.Parts = []shared.Part{
		part1("./day03/input.txt"),
		part2("./day03/input.txt"),
	}

	shared.PrettyPrint(result)
}

func part1(input string) shared.Part {
	lines := shared.ReadLines(input)

	start := time.Now()
	result := 0

	for line := range lines {
		result += atLeastOneGift(line)
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
		result += roboSantaDeliveries(line)
	}
	return shared.Part{
		Duration: time.Since(start),
		Result:   result,
	}
}

type coord struct {
	x, y int
}

func atLeastOneGift(input string) int {
	houses := map[coord]int{}

	current := coord{0, 0}

	// Start with a package in the middle
	houses[current] = 1

	for _, dir := range input {
		switch dir {
		case '<':
			current.x -= 1
		case '>':
			current.x += 1
		case '^':
			current.y -= 1
		case 'v':
			current.y += 1
		default:
			continue
		}
		if _, ok := houses[current]; !ok {
			houses[current] = 1
		} else {
			houses[current] += 1
		}
	}

	return len(houses)
}

func roboSantaDeliveries(input string) int {
	houses := map[coord]int{}

	santa := coord{0, 0}
	robot := coord{0, 0}

	// Start with a package in the middle
	houses[santa] = 1

	for i, dir := range input {
		robotTurn := i%2 != 0
		switch dir {
		case '<':
			if robotTurn {
				robot.x -= 1
			} else {
				santa.x -= 1
			}
		case '>':
			if robotTurn {
				robot.x += 1
			} else {
				santa.x += 1
			}
		case '^':
			if robotTurn {
				robot.y -= 1
			} else {
				santa.y -= 1
			}
		case 'v':
			if robotTurn {
				robot.y += 1
			} else {
				santa.y += 1
			}
		default:
			continue
		}
		current := santa
		if robotTurn {
			current = robot
		}
		if _, ok := houses[current]; !ok {
			houses[current] = 1
		} else {
			houses[current] += 1
		}
	}

	return len(houses)
}
