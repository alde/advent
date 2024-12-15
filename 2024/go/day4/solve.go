package day4

import (
	"strings"

	"alde.nu/advent/2024/utils"
)

func Solve() {
	s := utils.NewSolution(4, "Ceres Search", "./day4/input.txt")
	s.Solve(1, SolvePartOne)
	s.Solve(2, SolvePartTwo)

	s.Print()
}

func SolvePartOne(input []string) (int, error) {
	result := 0

	pad := padding(input)
	exes := findXes(pad)
	validWords := findValidWords(pad, exes)
	for _, v := range validWords {
		result += v
	}

	return result, nil
}

func SolvePartTwo(input []string) (int, error) {
	result := 0

	pad := padding(input)
	exes := findAs(pad)
	validWords := findXMas(pad, exes)

	result = len(validWords)

	return result, nil
}

func reformat(input []string) []string {
	result := []string{}
	for _, row := range input {
		newRow := ""
		for _, ch := range row {
			if ch == 'X' || ch == 'M' || ch == 'A' || ch == 'S' {
				newRow += string(ch)
				continue
			}
			newRow += "."
		}
		result = append(result, newRow)
	}

	return result
}

func padding(input []string) []string {
	rowLenght := len(input[0]) + 6
	blankRows := []string{
		strings.Repeat(".", rowLenght),
		strings.Repeat(".", rowLenght),
		strings.Repeat(".", rowLenght),
	}
	paddedInput := []string{}

	for r := range input {
		paddedInput = append(paddedInput, strings.Join([]string{"...", input[r], "..."}, ""))
	}

	paddedInput = append(blankRows, paddedInput...)
	paddedInput = append(paddedInput, blankRows...)

	return paddedInput
}

type Coord struct {
	X, Y int
}

func findXes(input []string) []Coord {
	coords := []Coord{}
	for r := range input {
		for c := range input[r] {
			if input[r][c] == 'X' {
				coords = append(coords, Coord{c, r})
			}
		}
	}

	return coords
}

func findAs(input []string) []Coord {
	coords := []Coord{}
	for r := range input {
		for c := range input[r] {
			if input[r][c] == 'A' {
				coords = append(coords, Coord{c, r})
			}
		}
	}

	return coords
}

func findXMas(grid []string, startPoints []Coord) []Coord {
	seen := map[Coord]int{}

	for _, point := range startPoints {
		// Check Diagonal UL to DR
		// M..
		// .A.
		// ..S
		if grid[point.Y-1][point.X-1] == 'M' && grid[point.Y+1][point.X+1] == 'S' {
			seen[point] += 1
		}
		// Check Diagonal DR to UL
		// S..
		// .A.
		// ..M
		if grid[point.Y+1][point.X+1] == 'M' && grid[point.Y-1][point.X-1] == 'S' {
			seen[point] += 1
		}
		// Check Diagonal DL to UR
		// ..S
		// .A.
		// M..
		if grid[point.Y+1][point.X-1] == 'M' && grid[point.Y-1][point.X+1] == 'S' {
			seen[point] += 1
		}
		// Check Diagonal UR to DL
		// ..M
		// .A.
		// S..
		if grid[point.Y-1][point.X+1] == 'M' && grid[point.Y+1][point.X-1] == 'S' {
			seen[point] += 1
		}
	}

	// Valid finds are where the count of a point is 2

	valid := []Coord{}

	for c, s := range seen {
		if s == 2 {
			valid = append(valid, c)
		}
	}

	return valid
}

func findValidWords(grid []string, startPoints []Coord) map[Coord]int {
	seen := map[Coord]int{}
	for _, point := range startPoints {
		// Check above
		if grid[point.Y-1][point.X] == 'M' && grid[point.Y-2][point.X] == 'A' && grid[point.Y-3][point.X] == 'S' {
			seen[point] += 1
		}
		// Check down
		if grid[point.Y+1][point.X] == 'M' && grid[point.Y+2][point.X] == 'A' && grid[point.Y+3][point.X] == 'S' {
			seen[point] += 1
		}
		// Check left
		if grid[point.Y][point.X-1] == 'M' && grid[point.Y][point.X-2] == 'A' && grid[point.Y][point.X-3] == 'S' {
			seen[point] += 1
		}
		// Check right
		if grid[point.Y][point.X+1] == 'M' && grid[point.Y][point.X+2] == 'A' && grid[point.Y][point.X+3] == 'S' {
			seen[point] += 1
		}
		// Check Diagonal Up Left
		if grid[point.Y-1][point.X-1] == 'M' && grid[point.Y-2][point.X-2] == 'A' && grid[point.Y-3][point.X-3] == 'S' {
			seen[point] += 1
		}
		// Check Diagonal Up Right
		if grid[point.Y-1][point.X+1] == 'M' && grid[point.Y-2][point.X+2] == 'A' && grid[point.Y-3][point.X+3] == 'S' {
			seen[point] += 1
		}
		// Check Diagonal Down Right
		if grid[point.Y+1][point.X+1] == 'M' && grid[point.Y+2][point.X+2] == 'A' && grid[point.Y+3][point.X+3] == 'S' {
			seen[point] += 1
		}
		// Check Diagonal Down Left
		if grid[point.Y+1][point.X-1] == 'M' && grid[point.Y+2][point.X-2] == 'A' && grid[point.Y+3][point.X-3] == 'S' {
			seen[point] += 1
		}
	}

	return seen
}
