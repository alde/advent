package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"alde.nu/advent/2024/day1"
	"alde.nu/advent/2024/day2"
	"alde.nu/advent/2024/day3"
	"alde.nu/advent/2024/day4"
)

func main() {
	target := flag.String("target", "all", "specify which day to run, or all")
	flag.Parse()

	days := map[int]func(){
		1: day1.Solve,
		2: day2.Solve,
		3: day3.Solve,
		4: day4.Solve,
	}

	if *target == "all" {
		for _, d := range days {
			d()
		}
		return
	}

	day, err := strconv.Atoi(*target)
	if err != nil {
		fmt.Printf("day must be either 'all' or an integer. Got '%v'\n", *target)
		os.Exit(2)
	}

	dayFunc, ok := days[day]
	if !ok {
		fmt.Printf("day %d has not been added to the list of solved days\n", day)
		os.Exit(2)
	}
	dayFunc()
}
