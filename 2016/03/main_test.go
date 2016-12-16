package main

import (
	"github.com/jyggen/advent-of-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, 0, solvePartOne("5  10  25"))
	assert.Equal(t, 862, solvePartOne(util.ReadFile("input")))
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(t, 6, solvePartTwo("101  301  501\n102  302  502\n103  303  503\n201  401  601\n202  402  602\n203  403  603"))
	assert.Equal(t, 1577, solvePartTwo(util.ReadFile("input")))
}
