package day05

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ assert.Assertions
var _ testing.T

func Test_isNice(t *testing.T) {
	tdata := []struct {
		input    string
		expected bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for i, td := range tdata {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			actual := isNice(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}

func Test_BetterIsNice(t *testing.T) {
	tdata := []struct {
		input    string
		expected bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for i, td := range tdata {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			actual := betterIsNice(td.input)
			assert.Equal(t, td.expected, actual)
		})
	}
}
