package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	input     string
	expected1 int
	expected2 int
}

func TestSolve(t *testing.T) {
	assert := assert.New(t)
	tests := []Test{
		Test{input: ">", expected1: 2, expected2: 3},
		Test{input: "^>v<", expected1: 4, expected2: 3},
		Test{input: "^v^v^v^v^v", expected1: 2, expected2: 11},
	}

	for _, test := range tests {
		oneSanta, TwoSantas := solve(test.input)

		assert.Equal(test.expected1, oneSanta)
		assert.Equal(test.expected2, TwoSantas)
	}
}
