package main

import (
	"github.com/jyggen/advent-of-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, "18f47a30", solvePartOne("abc"))
	assert.Equal(t, "f97c354d", solvePartOne(util.ReadFile("input")))
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(t, "05ace8e3", solvePartTwo("abc"))
	assert.Equal(t, "863dde27", solvePartTwo(util.ReadFile("input")))
}
