package main

import (
	"github.com/jyggen/advent-of-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, 2, solvePartOne("abba[mnop]qrst\nabcd[bddb]xyyx\naaaa[qwer]tyui\nioxxoj[asdfgh]zxcvbn"))
	assert.Equal(t, 105, solvePartOne(util.ReadFile("input")))
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(t, 3, solvePartTwo("aba[bab]xyz\nxyx[xyx]xyx\naaa[kek]eke\nzazbz[bzb]cdb"))
	assert.Equal(t, 258, solvePartTwo(util.ReadFile("input")))
}
