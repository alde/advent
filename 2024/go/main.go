package main

import (
	"flag"
	"fmt"
	"strconv"

	"alde.nu/advent/2024/day1"
)

func main() {
	target := flag.String("target", "all", "specify which day to run, or all")
	flag.Parse()

	days := []func(){
		day1.Solve,
	}

	if *target == "all" {
		for _, d := range days {
			d()
		}
		return
	}

	day, err := strconv.Atoi(*target)
	if err != nil {
		fmt.Printf("day must be either 'all' or an integer. Got %+v", target)
	} else {
		days[day]()
	}
}
