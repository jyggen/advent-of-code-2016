package main

import (
	//"github.com/jyggen/advent-of-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, 3, solvePartOne(parseInput("5-8\n0-2\n4-7")))
	//assert.Equal(t, 1956, solve(util.ReadFile("input"), 40))
}

func TestSolvePartTwo(t *testing.T) {
}
