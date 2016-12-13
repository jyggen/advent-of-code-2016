package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	input := `cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`

	assert.Equal(t, 42, solve(parseInput(input), map[string]int{
		"a": 0,
	}))
}
