package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, 0, solvePartOne("5  10  25"))
}

func TestSolvePartTwo(t *testing.T) {
	input := `101  301  501
102  302  502
103  303  503
201  401  601
202  402  602
203  403  603`

	assert.Equal(t, 6, solvePartTwo(input))
}
