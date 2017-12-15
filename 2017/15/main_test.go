package main

import (
	"github.com/jyggen/advent-of-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, 588, solvePartOne(parseInput("Generator A starts with 65\nGenerator B starts with 8921")))
	assert.Equal(t, 573, solvePartOne(parseInput(util.ReadFile("input"))))
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(t, 309, solvePartTwo(parseInput("Generator A starts with 65\nGenerator B starts with 8921")))
	assert.Equal(t, 294, solvePartTwo(parseInput(util.ReadFile("input"))))
}
