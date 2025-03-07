package day01

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Parse(t *testing.T) {
	tdata := []struct {
		input    string
		expected int
	}{
		{input: "(())", expected: 0},
		{input: "()()", expected: 0},
		{input: "(((", expected: 3},
		{input: "(()(()(", expected: 3},
		{input: "))(((((", expected: 3},
		{input: "())", expected: -1},
		{input: "))(", expected: -1},
		{input: ")))", expected: -3},
		{input: ")())())", expected: -3},
	}

	for _, td := range tdata {

		t.Run(fmt.Sprintf("Test %s", td.input), func(t *testing.T) {
			actual := parse(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}

func Test_FindBasement(t *testing.T) {
	tdata := []struct {
		input    string
		expected int
	}{
		{input: "(()))", expected: 5},
		{input: ")", expected: 1},
	}

	for _, td := range tdata {

		t.Run(fmt.Sprintf("Test %s", td.input), func(t *testing.T) {
			actual := findBasement(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}
