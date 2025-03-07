package day02

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tdata := []struct {
		input    string
		expected [3]int
	}{
		{input: "1x1x10", expected: [3]int{1, 1, 10}},
	}

	for _, td := range tdata {

		t.Run(fmt.Sprintf("Test %v", td.input), func(t *testing.T) {
			a, b, c := parse(td.input)
			assert.Equal(t, td.expected, [3]int{a, b, c})
		})
	}
}

func TestWrapping(t *testing.T) {
	tdata := []struct {
		l        int
		w        int
		h        int
		expected int
	}{
		{l: 2, w: 3, h: 4, expected: 58},
		{l: 1, w: 1, h: 10, expected: 43},
	}

	for i, td := range tdata {

		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			actual := wrappingPaper(td.l, td.w, td.h)
			assert.Equal(t, td.expected, actual)
		})
	}
}
func TestRibbon(t *testing.T) {
	tdata := []struct {
		l        int
		w        int
		h        int
		expected int
	}{
		{l: 2, w: 3, h: 4, expected: 34},
		{l: 1, w: 1, h: 10, expected: 14},
	}

	for i, td := range tdata {

		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			actual := ribbon(td.l, td.w, td.h)
			assert.Equal(t, td.expected, actual)
		})
	}
}
