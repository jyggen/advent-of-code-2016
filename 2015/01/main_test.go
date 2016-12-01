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
		Test{input: "(())", expected: 0},
		Test{input: "()()", expected: 0},
		Test{input: "(((", expected: 3},
		Test{input: "(()(()(", expected: 3},
		Test{input: "))(((((", expected: 3},
		Test{input: "())", expected: -1},
		Test{input: "))(", expected: -1},
		Test{input: ")))", expected: -3},
		Test{input: ")())())", expected: -3},
	}

	for _, test := range tests {
		partOne := solvePartOne(test.input)

		assert.Equal(test.expected, partOne)
	}
}

func TestSolvePartTwo(t *testing.T) {
	assert := assert.New(t)
	tests := []Test{
		Test{input: ")", expected: 1},
		Test{input: "()())", expected: 5},
	}

	for _, test := range tests {
		partTwo := solvePartTwo(test.input)

		assert.Equal(test.expected, partTwo)
	}
}
