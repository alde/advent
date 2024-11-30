package day1

import (
	"time"

	"alde.nu/advent/2024/utils"
)

func Solve() {
	s := utils.NewSolution(1, "test title", "./input.txt")
	s.Solve(1, SolvePartOne)
	s.Solve(2, SolvePartTwo)

	s.Print()
}

func SolvePartOne(input <-chan *utils.LineData) int {
	return 1
}

func SolvePartTwo(input <-chan *utils.LineData) int {
	time.Sleep(1 * time.Second)
	return 2
}
