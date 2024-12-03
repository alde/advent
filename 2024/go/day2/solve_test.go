package day2

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
	assert.Equal(t, 2, res)
}

func Test_SolvePartTwo(t *testing.T) {
	testData, _ := utils.ReadLines(testInput)

	res, err := SolvePartTwo(testData)
	assert.NoError(t, err)
	assert.Equal(t, 4, res)
}

func Test_isSafeReport(t *testing.T) {
	assert.True(t, isSafeReport("1 2 3 4 5", false))
	assert.True(t, isSafeReport("5 4 3 2 1", false))
	assert.True(t, isSafeReport("7 6 4 2 1", false))
	assert.False(t, isSafeReport("1 2 7 8 9", false))
	assert.True(t, isSafeReport("1 2 7", true))
}

func Test_safeReports(t *testing.T) {
	testData, _ := utils.ReadLines(testInput)
	assert.Equal(t, 2, safeReports(testData, false))
	assert.Equal(t, 4, safeReports(testData, true))
}
