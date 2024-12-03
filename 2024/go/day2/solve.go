package day2

import (
	"strings"

	"alde.nu/advent/2024/utils"
)

func Solve() {
	s := utils.NewSolution(2, "Red-Nosed Reports", "./day2/input.txt")
	s.Solve(1, SolvePartOne)
	s.Solve(2, SolvePartTwo)

	s.Print()
}

func SolvePartOne(input []string) (int, error) {
	result := safeReports(input, false)
	return result, nil
}

func SolvePartTwo(input []string) (int, error) {
	result := safeReports(input, true)

	return result, nil
}

func safeReports(input []string, dampener bool) int {
	safe := 0
	for _, row := range input {
		if isSafeReport(row, dampener) {
			safe += 1
		}
	}
	return safe
}

func isSafeReport(row string, dampener bool) bool {
	report := makeReport(row)

	// Check if the report is safe without using the dampener
	if isStrictlySafe(report) {
		return true
	}

	// If dampener is active, try removing one element at a time
	if dampener {
		for i := 0; i < len(report); i++ {
			// Create a new report with the i-th element removed
			modifiedReport := append([]int{}, report[:i]...)         // Elements before i
			modifiedReport = append(modifiedReport, report[i+1:]...) // Elements after i

			// Check if the modified report is safe
			if isStrictlySafe(modifiedReport) {
				return true
			}
		}
	}

	return false
}

func isStrictlySafe(report []int) bool {
	r := report[0]
	isIncreasing := true
	for i, a := range report {
		if i == 0 {
			continue
		}
		if i == 1 {
			if a > r {
				isIncreasing = true
			} else {
				isIncreasing = false
			}
		}

		if !compareStep(r, a, isIncreasing) {
			return false
		}
		r = a
	}
	return true
}

func compareStep(a, b int, isIncreasing bool) bool {
	if a == b {
		return false
	}
	if isIncreasing {
		return a < b && abs(a-b) <= 3
	}
	return a > b && abs(a-b) <= 3
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func makeReport(row string) []int {
	report := []int{}
	for _, s := range strings.Split(row, " ") {
		report = append(report, utils.ToInt(s))
	}
	return report
}
