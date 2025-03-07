package day06

import (
	"fmt"
	"strings"
	"time"

	"github.com/alde/advent/2015/golang/shared"
)

func Solve() {
	shared.InputMustExist("./day06/input.txt", 2015, 6)
	result := shared.Result{
		Title: "Probably a Fire Hazard",
		Day:   6,
	}

	result.Parts = []shared.Part{
		part1("./day06/input.txt"),
		part2("./day06/input.txt"),
	}

	shared.PrettyPrint(result)
}

type pair struct {
	x, y int
}

func part1(input string) shared.Part {
	lines := shared.ReadLines(input)

	start := time.Now()
	grid := map[pair]int{}

	for line := range lines {
		processInstruction(line, grid)
	}
	return shared.Part{
		Duration: time.Since(start),
		Result:   countLights(grid),
	}
}

func part2(input string) shared.Part {
	lines := shared.ReadLines(input)

	start := time.Now()
	grid := map[pair]int{}

	for line := range lines {
		processAdvancedInstructions(line, grid)
	}
	return shared.Part{
		Duration: time.Since(start),
		Result:   countLights(grid),
	}
}

func processInstruction(line string, grid map[pair]int) {
	if strings.HasPrefix(line, "toggle") {
		var rowstart, colstart, rowend, colend int
		fmt.Sscanf(line, "toggle %d,%d through %d,%d", &rowstart, &colstart, &rowend, &colend)
		for r := rowstart; r <= rowend; r++ {
			for c := colstart; c <= colend; c++ {
				if grid[pair{r, c}] == 0 {
					grid[pair{r, c}] = 1
				} else {
					grid[pair{r, c}] = 0
				}
			}
		}
	}
	if strings.HasPrefix(line, "turn on") {
		var rowstart, colstart, rowend, colend int
		fmt.Sscanf(line, "turn on %d,%d through %d,%d", &rowstart, &colstart, &rowend, &colend)
		for r := rowstart; r <= rowend; r++ {
			for c := colstart; c <= colend; c++ {
				grid[pair{r, c}] = 1
			}
		}
	}
	if strings.HasPrefix(line, "turn off") {
		var rowstart, colstart, rowend, colend int
		fmt.Sscanf(line, "turn off %d,%d through %d,%d", &rowstart, &colstart, &rowend, &colend)
		for r := rowstart; r <= rowend; r++ {
			for c := colstart; c <= colend; c++ {
				grid[pair{r, c}] = 0
			}
		}
	}
}

func processAdvancedInstructions(line string, grid map[pair]int) {
	if strings.HasPrefix(line, "toggle") {
		var rowstart, colstart, rowend, colend int
		fmt.Sscanf(line, "toggle %d,%d through %d,%d", &rowstart, &colstart, &rowend, &colend)
		for r := rowstart; r <= rowend; r++ {
			for c := colstart; c <= colend; c++ {
				grid[pair{r, c}] += 2
			}
		}
	}
	if strings.HasPrefix(line, "turn on") {
		var rowstart, colstart, rowend, colend int
		fmt.Sscanf(line, "turn on %d,%d through %d,%d", &rowstart, &colstart, &rowend, &colend)
		for r := rowstart; r <= rowend; r++ {
			for c := colstart; c <= colend; c++ {
				grid[pair{r, c}] += 1
			}
		}
	}
	if strings.HasPrefix(line, "turn off") {
		var rowstart, colstart, rowend, colend int
		fmt.Sscanf(line, "turn off %d,%d through %d,%d", &rowstart, &colstart, &rowend, &colend)
		for r := rowstart; r <= rowend; r++ {
			for c := colstart; c <= colend; c++ {
				grid[pair{r, c}] -= 1
				if grid[pair{r, c}] < 0 {
					grid[pair{r, c}] = 0
				}
			}
		}
	}
}

func countLights(grid map[pair]int) int {
	count := 0
	for _, val := range grid {
		count += val
	}
	return count
}
