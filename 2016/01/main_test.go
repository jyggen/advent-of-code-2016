package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert := assert.New(t)

	blocksAway, _ := solve(parseInput("R2, L3"))
	assert.Equal(5, blocksAway)

	blocksAway, _ = solve(parseInput("R2, R2, R2"))
	assert.Equal(2, blocksAway)

	blocksAway, _ = solve(parseInput("R5, L5, R5, R3"))
	assert.Equal(12, blocksAway)
}

func TestSolvePartTwo(t *testing.T) {
	assert := assert.New(t)

	_, firstRepeat := solve(parseInput("R8, R4, R4, R8"))
	assert.Equal(4, firstRepeat)
}
