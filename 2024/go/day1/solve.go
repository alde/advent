package day1

import (
	"container/heap"
	"regexp"
	"strconv"

	"alde.nu/advent/2024/utils"
)

func Solve() {
	s := utils.NewSolution(1, "Historian Hysteria", "./day1/input.txt")
	s.Solve(1, SolvePartOne)
	s.Solve(2, SolvePartTwo)

	s.Print()
}

func makeHeaps(testData []string) (*utils.MinHeap, *utils.MinHeap) {
	left := &utils.MinHeap{}
	right := &utils.MinHeap{}
	re := regexp.MustCompile(`\s+`)
	for _, line := range testData {
		if len(line) == 0 {
			break
		}

		parts := re.Split(line, -1)
		l, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])

		heap.Push(left, l)
		heap.Push(right, r)
	}

	return left, right
}

func SolvePartOne(input []string) (int, error) {
	result := 0
	left, right := makeHeaps(input)
	for left.Len() > 0 && right.Len() > 0 {
		l := heap.Pop(left).(int)
		r := heap.Pop(right).(int)

		result += abs(r - l)
	}

	return result, nil
}

func SolvePartTwo(input []string) (int, error) {
	result := 0
	left, right := makeHeaps(input)
	occurances := countOccurances(right)

	for left.Len() > 0 {
		l := heap.Pop(left).(int)
		result += l * occurances[l]
	}
	return result, nil
}

func countOccurances(h *utils.MinHeap) map[int]int {
	count := map[int]int{}
	for h.Len() > 0 {
		r := heap.Pop(h).(int)
		_, ok := count[r]
		if !ok {
			count[r] = 1
		} else {
			count[r] += 1
		}
	}
	return count
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
