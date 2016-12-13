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

	assert.Equal(t, "1985", solve(parseInput(input), [][]rune{
		0: {0: []rune("1")[0], 1: []rune("2")[0], 2: []rune("3")[0]},
		1: {0: []rune("4")[0], 1: []rune("5")[0], 2: []rune("6")[0]},
		2: {0: []rune("7")[0], 1: []rune("8")[0], 2: []rune("9")[0]},
	}, 1, 1))
}

func TestSolvePartTwo(t *testing.T) {
	input := `ULL
RRDDD
LURDL
UUUUD`

	assert.Equal(t, "5DB3", solve(parseInput(input), [][]rune{
		0: {0: []rune("0")[0], 1: []rune("0")[0], 2: []rune("1")[0], 3: []rune("0")[0], 4: []rune("0")[0]},
		1: {0: []rune("0")[0], 1: []rune("2")[0], 2: []rune("3")[0], 3: []rune("4")[0], 4: []rune("0")[0]},
		2: {0: []rune("5")[0], 1: []rune("6")[0], 2: []rune("7")[0], 3: []rune("8")[0], 4: []rune("9")[0]},
		3: {0: []rune("0")[0], 1: []rune("A")[0], 2: []rune("B")[0], 3: []rune("C")[0], 4: []rune("0")[0]},
		4: {0: []rune("0")[0], 1: []rune("0")[0], 2: []rune("D")[0], 3: []rune("0")[0], 4: []rune("0")[0]},
	}, 2, 0))
}
