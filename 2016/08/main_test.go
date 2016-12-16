package main

import (
	"github.com/jyggen/advent-of-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	answer, _ := solve("rect 3x2\nrotate column x=1 by 1\nrotate row y=0 by 4\nrotate column x=1 by 1", 7, 3)

	assert.Equal(t, 6, answer)
}
