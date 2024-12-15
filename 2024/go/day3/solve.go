package day3

import (
	"strconv"
	"strings"

	"alde.nu/advent/2024/utils"
)

func Solve() {
	s := utils.NewSolution(3, "Mull It Over", "./day3/input.txt")
	s.Solve(1, SolvePartOne)
	s.Solve(2, SolvePartTwo)

	s.Print()
}

func SolvePartOne(input []string) (int, error) {
	result := 0
	multis, err := parseMultis(strings.Join(input, ""))
	if err != nil {
		panic(err)
	}
	for _, m := range multis {
		result += m[0] * m[1]
	}

	return result, nil
}

func SolvePartTwo(input []string) (int, error) {
	result := 0
	multis, err := parseMultis2(strings.Join(input, ""))
	if err != nil {
		panic(err)
	}
	for _, m := range multis {
		result += m[0] * m[1]
	}

	return result, nil
}

func parseMul(s string, from int, to int) ([]int, error) {
	parts := strings.Split(s[from:to+1], ",")

	p1, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	p2, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return []int{p1, p2}, nil
}

func enable(str string, index int) bool {
	if len(str) < index+3 {
		return false
	}
	return str[index] == 'd' && str[index+1] == 'o' && str[index+2] == '(' && str[index+3] == ')'
}

func disable(str string, index int) bool {
	if len(str) < index+6 {
		return false
	}
	return str[index] == 'd' && str[index+1] == 'o' &&
		str[index+2] == 'n' && str[index+3] == '\'' &&
		str[index+4] == 't' && str[index+5] == '(' &&
		str[index+6] == ')'
}

func checkIfMul(str string, index int) (bool, int, int) {
	if len(str) < index+4 {
		return false, -1, -1
	}
	startsRight := str[index] == 'm' && str[index+1] == 'u' && str[index+2] == 'l' && str[index+3] == '('
	if !startsRight {
		return false, -1, -1
	}

	end := findEnd(str, index)
	if end == -1 {
		return false, -1, -1
	}

	return true, index + 4, end - 1
}

func findEnd(str string, index int) int {
	for i := index + 4; i < len(str); i++ {
		if str[i] >= 48 && str[i] <= 71 {
			continue
		}
		if str[i] == ',' {
			continue
		}
		if str[i] == ')' {
			return i
		}
		// not a valid item between parens
		return -1
	}
	return -1
}

func parseMultis(s string) ([][]int, error) {
	res := [][]int{}
	for i := 0; i < len(s); i++ {
		if s[i] == 'm' {
			isMul, start, end := checkIfMul(s, i)
			if isMul {
				p, err := parseMul(s, start, end)
				if err != nil {
					panic(err)
				}
				res = append(res, p)

				i = end
				continue
			}
		}
	}

	return res, nil
}

func parseMultis2(s string) ([][]int, error) {
	res := [][]int{}
	enabled := true
	for i := 0; i < len(s); i++ {
		if s[i] == 'd' {
			if enabled {
				enabled = !disable(s, i)
			} else {
				enabled = enable(s, i)
			}
		}
		if s[i] == 'm' {
			isMul, start, end := checkIfMul(s, i)
			if isMul {
				p, err := parseMul(s, start, end)
				if err != nil {
					panic(err)
				}
				if enabled {
					res = append(res, p)
				}
				i = end
				continue
			}
		}
	}

	return res, nil
}
