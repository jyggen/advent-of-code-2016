package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	input := `value 5 goes to bot 2
bot 2 gives low to bot 1 and high to bot 0
value 3 goes to bot 1
bot 1 gives low to output 1 and high to bot 0
bot 0 gives low to output 2 and high to output 0
value 2 goes to bot 2`

	part1, _ := solve(input, 5, 2)

	assert.Equal(t, 2, part1)
}
