package day06

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ assert.Assertions
var _ testing.T

func Test_processInstruction(t *testing.T) {
	grid := map[pair]int{}
	tdata := []struct {
		input    string
		expected map[pair]int
	}{
		{
			"turn on 0,0 through 2,2",
			map[pair]int{
				{0, 0}: 1,
				{0, 1}: 1,
				{0, 2}: 1,
				{1, 0}: 1,
				{1, 1}: 1,
				{1, 2}: 1,
				{2, 0}: 1,
				{2, 1}: 1,
				{2, 2}: 1,
			},
		},
		{
			"turn off 0,0 through 1,1",
			map[pair]int{
				{0, 0}: 0,
				{0, 1}: 0,
				{0, 2}: 1,
				{1, 0}: 0,
				{1, 1}: 0,
				{1, 2}: 1,
				{2, 0}: 1,
				{2, 1}: 1,
				{2, 2}: 1,
			},
		},
		{
			"toggle 0,0 through 3,3",
			map[pair]int{
				{0, 0}: 1,
				{0, 1}: 1,
				{0, 2}: 0,
				{0, 3}: 1,
				{1, 0}: 1,
				{1, 1}: 1,
				{1, 2}: 0,
				{1, 3}: 1,
				{2, 0}: 0,
				{2, 1}: 0,
				{2, 2}: 0,
				{2, 3}: 1,
				{3, 0}: 1,
				{3, 1}: 1,
				{3, 2}: 1,
				{3, 3}: 1,
			},
		},
	}

	for i, td := range tdata {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			processInstruction(td.input, grid)
			assert.Equal(t, td.expected, grid)
		})
	}
}

func Test_processAdvancedInstructions(t *testing.T) {
	grid := map[pair]int{}
	tdata := []struct {
		input    string
		expected map[pair]int
	}{
		{
			"turn on 0,0 through 2,2",
			map[pair]int{
				{0, 0}: 1,
				{0, 1}: 1,
				{0, 2}: 1,
				{1, 0}: 1,
				{1, 1}: 1,
				{1, 2}: 1,
				{2, 0}: 1,
				{2, 1}: 1,
				{2, 2}: 1,
			},
		},
		{
			"turn off 0,0 through 1,1",
			map[pair]int{
				{0, 0}: 0,
				{0, 1}: 0,
				{0, 2}: 1,
				{1, 0}: 0,
				{1, 1}: 0,
				{1, 2}: 1,
				{2, 0}: 1,
				{2, 1}: 1,
				{2, 2}: 1,
			},
		},
		{
			"toggle 0,0 through 3,3",
			map[pair]int{
				{0, 0}: 2,
				{0, 1}: 2,
				{0, 2}: 3,
				{0, 3}: 2,
				{1, 0}: 2,
				{1, 1}: 2,
				{1, 2}: 3,
				{1, 3}: 2,
				{2, 0}: 3,
				{2, 1}: 3,
				{2, 2}: 3,
				{2, 3}: 2,
				{3, 0}: 2,
				{3, 1}: 2,
				{3, 2}: 2,
				{3, 3}: 2,
			},
		},
	}

	for i, td := range tdata {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			processAdvancedInstructions(td.input, grid)
			assert.Equal(t, td.expected, grid)
		})
	}
}

func Test_CountLights(t *testing.T) {
	tdata := []struct {
		input    map[pair]int
		expected int
	}{
		{
			map[pair]int{
				{0, 0}: 1,
				{0, 1}: 1,
				{0, 2}: 0,
				{0, 3}: 1,
				{1, 0}: 1,
				{1, 1}: 1,
				{1, 2}: 0,
				{1, 3}: 1,
				{2, 0}: 0,
				{2, 1}: 0,
				{2, 2}: 0,
				{2, 3}: 1,
				{3, 0}: 1,
				{3, 1}: 1,
				{3, 2}: 1,
				{3, 3}: 1,
			},
			11,
		},
	}
	for i, td := range tdata {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			actual := countLights(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}
