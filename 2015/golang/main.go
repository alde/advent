package main

import (
	"fmt"

	"github.com/alde/advent/2015/golang/day01"
	"github.com/alde/advent/2015/golang/day02"
	"github.com/alde/advent/2015/golang/day03"
	"github.com/alde/advent/2015/golang/day04"
)

type Solution struct {
	Fn func()
}

func main() {
	solutions := []func(){
		day01.Solve,
		day02.Solve,
		day03.Solve,
		day04.Solve,
	}

	for _, solution := range solutions {
		solution()
		fmt.Println()
	}

}
