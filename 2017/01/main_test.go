package main

import (
	"github.com/jyggen/advent-of-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, 3, solvePartOne(parseInput("1122")))
	assert.Equal(t, 4, solvePartOne(parseInput("1111")))
	assert.Equal(t, 0, solvePartOne(parseInput("1234")))
	assert.Equal(t, 9, solvePartOne(parseInput("91212129")))
	assert.Equal(t, 1253, solvePartOne(parseInput(util.ReadFile("input"))))
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(t, 6, solvePartTwo(parseInput("1212")))
	assert.Equal(t, 0, solvePartTwo(parseInput("1221")))
	assert.Equal(t, 4, solvePartTwo(parseInput("123425")))
	assert.Equal(t, 12, solvePartTwo(parseInput("123123")))
	assert.Equal(t, 4, solvePartTwo(parseInput("12131415")))
	assert.Equal(t, 1278, solvePartTwo(parseInput(util.ReadFile("input"))))
}
