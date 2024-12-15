package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SolvePartOne(t *testing.T) {
	testData := []string{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"}

	res, err := SolvePartOne(testData)
	assert.NoError(t, err)
	assert.Equal(t, 161, res)
}

func Test_SolvePartTwo(t *testing.T) {
	testData := []string{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"}

	res, err := SolvePartTwo(testData)
	assert.NoError(t, err)
	assert.Equal(t, 48, res)
}

func Test_ParseMultis(t *testing.T) {
	tdata := []struct {
		input    string
		expected [][]int
	}{
		{"mul(2,4)", [][]int{{2, 4}}},
		{"xmul(2,4)%&mul[3,7]", [][]int{{2, 4}}},
		{"xmul(2,4)%&mul(3,7)", [][]int{{2, 4}, {3, 7}}},
	}
	for _, td := range tdata {
		res, err := parseMultis(td.input)
		assert.NoError(t, err)
		assert.Equal(t, td.expected, res)

	}
}

func Test_ParseMultis2(t *testing.T) {
	tdata := []struct {
		input    string
		expected [][]int
	}{
		{"don't()mul(2,4)", [][]int{}},
		{"xmul(2,4)%&mul[3,7]", [][]int{{2, 4}}},
		{"don't()xmul(2,4)%do()&mul(3,7)", [][]int{{3, 7}}},
	}
	for _, td := range tdata {
		res, err := parseMultis2(td.input)
		assert.NoError(t, err)
		assert.Equal(t, td.expected, res)

	}
}
