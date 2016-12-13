package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	shortest, _ := solve(10, 7, 4)

	assert.Equal(t, 11, shortest)
}
