package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ assert.Assertions
var _ testing.T

func Test_FindLowestCollision(t *testing.T) {
	assert.Equal(t, 609043, firstNumber("abcdef", 5))
	assert.Equal(t, 1048970, firstNumber("pqrstuv", 5))
}
