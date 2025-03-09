package day07

import (
	"fmt"
	"testing"

	"github.com/alde/advent/2015/golang/shared"
	"github.com/stretchr/testify/assert"
)

var _ assert.Assertions
var _ testing.T

func Test_processInstruction(t *testing.T) {
	grid := map[string]uint16{}
	tdata := []struct {
		input            string
		expectedRegister string
		expectedValue    uint16
		expected         map[string]uint16
	}{
		{
			"123 -> x", "x", 123, map[string]uint16{"x": 123},
		},
		{
			"456 -> y", "y", 456, map[string]uint16{"x": 123, "y": 456},
		},
		{
			"x AND y -> d", "d", 72, map[string]uint16{"x": 123, "y": 456, "d": 72},
		},
		{
			"x OR y -> e", "e", 507, map[string]uint16{"x": 123, "y": 456, "d": 72, "e": 507},
		},
		{
			"x LSHIFT 2 -> f", "f", 492, map[string]uint16{"x": 123, "y": 456, "d": 72, "e": 507, "f": 492},
		},
		{
			"y RSHIFT 2 -> g", "g", 114, map[string]uint16{"x": 123, "y": 456, "d": 72, "e": 507, "f": 492, "g": 114},
		},
		{
			"NOT x -> h", "h", 65412, map[string]uint16{"x": 123, "y": 456, "d": 72, "e": 507, "f": 492, "g": 114, "h": 65412},
		},
		{
			"NOT y -> i", "i", 65079, map[string]uint16{"x": 123, "y": 456, "d": 72, "e": 507, "f": 492, "g": 114, "h": 65412, "i": 65079},
		},
	}

	for _, td := range tdata {
		t.Run(fmt.Sprintf("Test %s", td.input), func(t *testing.T) {
			processInstruction(td.input, grid, map[string]string{})
			assert.Equal(t, td.expectedValue, grid[td.expectedRegister])
			assert.Equal(t, td.expected, grid)
		})
	}
}

func Test_prepareWires(t *testing.T) {
	input := shared.Stream([]string{
		"123 -> x",
		"456 -> y",
		"x AND y -> d",
		"x OR y -> e",
		"x LSHIFT 2 -> f",
		"y RSHIFT 2 -> g",
		"NOT x -> h",
		"NOT y -> i",
	})
	wires := prepareWires(input)
	assert.Equal(t, map[string]string{
		"x": "123 -> x",
		"y": "456 -> y",
		"d": "x AND y -> d",
		"e": "x OR y -> e",
		"f": "x LSHIFT 2 -> f",
		"g": "y RSHIFT 2 -> g",
		"h": "NOT x -> h",
		"i": "NOT y -> i",
	}, wires)
}

func Test_emulateCircuit(t *testing.T) {
	grid := map[string]uint16{}
	input := shared.Stream([]string{
		"NOT x -> h",
		"NOT y -> i",
		"x AND y -> d",
		"x LSHIFT 2 -> f",
		"x OR y -> e",
		"y RSHIFT 2 -> g",
		"123 -> x",
		"456 -> y",
	})
	wires := prepareWires(input)

	for output := range wires {
		emulateCircuit(output, grid, wires)
	}

	assert.Equal(t, uint16(72), grid["d"])
	assert.Equal(t, uint16(65079), grid["i"])
}
