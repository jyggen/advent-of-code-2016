package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	input := `ULL
RRDDD
LURDL
UUUUD`

	assert.Equal(t, "1985", solvePartOne(input))
}

func TestSolvePartTwo(t *testing.T) {
	input := `ULL
RRDDD
LURDL
UUUUD`

	assert.Equal(t, "5DB3", solvePartTwo(input))
}
