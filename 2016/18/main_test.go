package main

import (
	"github.com/jyggen/advent-of-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	assert.Equal(t, 6, solve("..^^.", 3))
	assert.Equal(t, 38, solve(".^^.^.^^^^", 10))
	assert.Equal(t, 1956, solve(util.ReadFile("input"), 40))
	assert.Equal(t, 19995121, solve(util.ReadFile("input"), 400000))
}
