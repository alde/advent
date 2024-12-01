package day1

import (
	"testing"

	"alde.nu/advent/2024/utils"
	"github.com/stretchr/testify/assert"
)

var testInput = "./input_test.txt"

func Test_SolvePartOne(t *testing.T) {
	testData, _ := utils.ReadLines(testInput)

	res, err := SolvePartOne(testData)
	assert.NoError(t, err)
	assert.Equal(t, 11, res)
}

func Test_SolvePartTwo(t *testing.T) {
	testData, _ := utils.ReadLines(testInput)

	res, err := SolvePartTwo(testData)
	assert.NoError(t, err)
	assert.Equal(t, 31, res)
}
