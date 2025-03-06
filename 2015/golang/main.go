package main

import (
	"fmt"

	"github.com/alde/advent/2015/golang/day01"
)

type Solution struct {
	Fn   func()
	Name string
}

func main() {
	solutions := make([]*Solution, 0, 25)
	solutions = append(solutions, &Solution{
		Fn:   day01.Solve,
		Name: day01.Title,
	})

	fmt.Printf("%d of %d days completed\n\n", len(solutions), cap(solutions))

	solutions[0].Fn()
}
