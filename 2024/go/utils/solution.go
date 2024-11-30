package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
)

type Solution struct {
	Day     int
	Title   string
	Lines   <-chan *LineData
	Results map[int]*Part
}

type Part struct {
	Result  int
	Elapsed time.Duration
}

func NewSolution(day int, title string, input string) *Solution {
	l := LazyReadLines(input)
	return &Solution{
		Day:     day,
		Title:   title,
		Lines:   l,
		Results: make(map[int]*Part),
	}
}

func (s *Solution) Solve(part int, f func(<-chan *LineData) int) {
	result, elapsed := WithTimer(func() int { return f(s.Lines) })
	s.Results[part] = &Part{
		Result:  result,
		Elapsed: elapsed,
	}
}

func (s *Solution) Print() {
	s.printHeader()
	s.printResults()
	// s.Table()
}

func (s *Solution) printHeader() {
	fmt.Printf("%d : %s\n", s.Day, s.Title)
}

func (s *Solution) printResults() {
	for i, r := range s.Results {
		fmt.Printf("\t⭐ [%12s]\tPart %d\t%v\n", r.Elapsed, i, r.Result)
	}
}

func (s *Solution) Table() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Part", "Elapsed", "Result"})
	for i, r := range s.Results {
		t.AppendRow(table.Row{i, r.Elapsed, r.Result})
	}
	t.Render()
}
