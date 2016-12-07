package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	input := `abba[mnop]qrst
abcd[bddb]xyyx
aaaa[qwer]tyui
ioxxoj[asdfgh]zxcvbn`

	assert.Equal(t, 2, solvePartOne(input))
}

func TestSolvePartTwo(t *testing.T) {
	input := `aba[bab]xyz
xyx[xyx]xyx
aaa[kek]eke
zazbz[bzb]cdb`

	assert.Equal(t, 3, solvePartTwo(input))
}
