package main

import (
	"github.com/jyggen/advent-of-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, 1514, solvePartOne("aaaaa-bbb-z-y-x-123[abxyz]\na-b-c-d-e-f-g-h-987[abcde]\nnot-a-real-room-404[oarel]\ntotally-real-room-200[decoy]"))
	assert.Equal(t, 158835, solvePartOne(util.ReadFile("input")))
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(t, 993, solvePartTwo(util.ReadFile("input")))
}
