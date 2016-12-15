package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve(t *testing.T) {
	assert.Equal(t, 22728, solve("abc", 0))
	assert.Equal(t, 22551, solve("abc", 2016))
}
