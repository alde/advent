package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SolvePartOne(t *testing.T) {
	testInput := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	res, err := SolvePartOne(testInput)
	assert.NoError(t, err)
	assert.Equal(t, 18, res)
}

func Test_SolvePartTwo(t *testing.T) {
	testInput := []string{
		"MMMSXXMASM",
		"MSAMXMSMSA",
		"AMXSXMAAMM",
		"MSAMASMSMX",
		"XMASAMXAMM",
		"XXAMMXXAMA",
		"SMSMSASXSS",
		"SAXAMASAAA",
		"MAMMMXMMMM",
		"MXMXAXMASX",
	}
	res, err := SolvePartTwo(testInput)
	assert.NoError(t, err)
	assert.Equal(t, 9, res)
}

func Test_ReformatInput(t *testing.T) {
	input := []string{
		"BTXYER",
		"LSAMXT",
		"DAOUAF",
		"XMASTS",
		"OXYNGE",
	}

	expected := []string{
		"..X...",
		".SAMX.",
		".A..A.",
		"XMAS.S",
		".X....",
	}

	actual := reformat(input)

	assert.Equal(t, expected, actual)
}

func Test_Padding(t *testing.T) {
	input := []string{
		"..X...",
		".SAMX.",
		".A..A.",
		"XMAS.S",
		".X....",
	}
	expected := []string{
		"............",
		"............",
		"............",
		".....X......",
		"....SAMX....",
		"....A..A....",
		"...XMAS.S...",
		"....X.......",
		"............",
		"............",
		"............",
	}
	actual := padding(input)

	assert.Equal(t, expected, actual)
}

func Test_FindXes(t *testing.T) {
	input := []string{
		"............",
		"............",
		"............",
		".....X......",
		"....SAMX....",
		"....A..A....",
		"...XMAS.S...",
		"....X.......",
		"............",
		"............",
		"............",
	}

	expected := []Coord{
		{5, 3},
		{7, 4},
		{3, 6},
		{4, 7},
	}
	actual := findXes(input)

	assert.Equal(t, expected, actual)
}

func Test_FindAs(t *testing.T) {
	input := []string{
		"............",
		"............",
		"............",
		".....X......",
		"....SAMX....",
		"....A..A....",
		"...XMAS.S...",
		"....X.......",
		"............",
		"............",
		"............",
	}

	expected := []Coord{
		{5, 4},
		{4, 5}, {7, 5},
		{5, 6},
	}
	actual := findAs(input)

	assert.Equal(t, expected, actual)
}

func Test_FindXMas(t *testing.T) {
	input := []string{
		".........",
		".........",
		".........",
		"...M.M...",
		"....A....",
		"...S.S...",
		".........",
		".........",
		".........",
	}

	expected := map[Coord]int{
		{4, 4}: 2,
	}

	actual := findXMas(input, []Coord{{4, 4}})

	assert.Equal(t, expected, actual)
}
