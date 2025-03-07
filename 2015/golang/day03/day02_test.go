package day03

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ assert.Assertions
var _ testing.T

func Test_PresentsDelivered(t *testing.T) {
	tdata := []struct {
		input    string
		expected int
	}{
		{input: ">", expected: 2},
		{input: "^>v<", expected: 4},
		{input: "^v^v^v^v^v", expected: 2},
	}

	for i, td := range tdata {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			actual := atLeastOneGift(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}

func Test_RoboSantaDeliveries(t *testing.T) {
	tdata := []struct {
		input    string
		expected int
	}{
		{input: "^v", expected: 3},
		{input: "^>v<", expected: 3},
		{input: "^v^v^v^v^v", expected: 11},
	}

	for i, td := range tdata {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			actual := roboSantaDeliveries(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}
