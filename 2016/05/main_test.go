package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, "18f47a30", solvePartOne("abc"))
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(t, "05ace8e3", solvePartTwo("abc"))
}
