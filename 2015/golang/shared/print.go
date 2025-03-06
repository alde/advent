package shared

import (
	"fmt"
	"time"
)

type Part struct {
	Duration time.Duration
	Result   int
}

type Result struct {
	Day   int
	Title string
	Parts []Part
}

func PrettyPrint(result Result) {
	fmt.Printf("%02d - %s\n", result.Day, result.Title)
	for i, p := range result.Parts {
		fmt.Printf("Part %d : %d (%s)\n", i+1, p.Result, p.Duration)
	}
}
