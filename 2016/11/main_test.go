package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	input := `The first floor contains a hydrogen-compatible microchip, and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.`

	assert.Equal(t, 11, solve(parseInput(input), 0, 0))
}
