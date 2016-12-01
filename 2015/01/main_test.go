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
		floor, _ := solve(test.input)

		assert.Equal(test.expected, floor)
	}
}

func TestSolvePartTwo(t *testing.T) {
	assert := assert.New(t)
	tests := []Test{
		Test{input: ")", expected: 1},
		Test{input: "()())", expected: 5},
	}

	for _, test := range tests {
		_, reachedBasementAt := solve(test.input)

		assert.Equal(test.expected, reachedBasementAt)
	}
}
