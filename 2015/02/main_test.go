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
		Test{input: "2x3x4", expected1: 58, expected2: 34},
		Test{input: "1x1x10", expected1: 43, expected2: 14},
	}

	for _, test := range tests {
		paperSize, ribbonSize := solve(test.input)

		assert.Equal(test.expected1, paperSize)
		assert.Equal(test.expected2, ribbonSize)
	}
}
