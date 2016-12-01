package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input    string
	expected int
}

func TestSolvePartOne(t *testing.T) {
	assert := assert.New(t)
	tests := []Test{
		Test{input: "R2, L3", expected: 5},
		Test{input: "R2, R2, R2", expected: 2},
		Test{input: "R5, L5, R5, R3", expected: 12},
	}

	for _, test := range tests {
		blocksAway, _ := solve(test.input)

		assert.Equal(test.expected, blocksAway)
	}
}

func TestSolvePartTwo(t *testing.T) {
	assert := assert.New(t)
	tests := []Test{
		Test{input: "R8, R4, R4, R8", expected: 4},
	}

	for _, test := range tests {
		_, firstVisitedTwice := solve(test.input)

		assert.Equal(test.expected, firstVisitedTwice)
	}
}
