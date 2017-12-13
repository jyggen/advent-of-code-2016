package main

import (
	"github.com/jyggen/advent-of-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, 24, solvePartOne(parseInput("0: 3\n1: 2\n4: 4\n6: 4")))
	assert.Equal(t, 1704, solvePartOne(parseInput(util.ReadFile("input"))))
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(t, 10, solvePartTwo(parseInput("0: 3\n1: 2\n4: 4\n6: 4")))
	assert.Equal(t, 3970918, solvePartTwo(parseInput(util.ReadFile("input"))))
}
