package main

import (
	"github.com/jyggen/advent-of-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(5, solvePartOne(parseInput("R2, L3")))
	assert.Equal(2, solvePartOne(parseInput("R2, R2, R2")))
	assert.Equal(12, solvePartOne(parseInput("R5, L5, R5, R3")))
	assert.Equal(307, solvePartOne(parseInput(util.ReadFile("input"))))
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(4, solvePartTwo(parseInput("R8, R4, R4, R8")))
	assert.Equal(165, solvePartTwo(parseInput(util.ReadFile("input"))))
}
