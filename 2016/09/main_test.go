package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, 6, solvePartOne("ADVENT"))
	assert.Equal(t, 7, solvePartOne("A(1x5)BC"))
	assert.Equal(t, 9, solvePartOne("(3x3)XYZ"))
	assert.Equal(t, 11, solvePartOne("A(2x2)BCD(2x2)EFG"))
	assert.Equal(t, 6, solvePartOne("(6x1)(1x3)A"))
	assert.Equal(t, 18, solvePartOne("X(8x2)(3x3)ABCY"))
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(t, 9, solvePartTwo([]rune("(3x3)XYZ")))
	assert.Equal(t, 20, solvePartTwo([]rune("X(8x2)(3x3)ABCY")))
	assert.Equal(t, 241920, solvePartTwo([]rune("(27x12)(20x12)(13x14)(7x10)(1x12)A")))
	assert.Equal(t, 445, solvePartTwo([]rune("(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN")))
}
