package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	input := `aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]`

	assert.Equal(t, 1514, solvePartOne(input))
}

func TestSolvePartTwo(t *testing.T) {
	input := `101  301  501
102  302  502
103  303  503
201  401  601
202  402  602
203  403  603`

	assert.Equal(t, 6, solvePartTwo(input))
}
