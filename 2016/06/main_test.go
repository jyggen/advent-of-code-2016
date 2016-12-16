package main

import (
	"github.com/jyggen/advent-of-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	assert.Equal(t, "easter", solvePartOne("eedadn\ndrvtee\neandsr\nraavrd\natevrs\ntsrnev\nsdttsa\nrasrtv\nnssdts\nntnada\nsvetve\ntesnvt\nvntsnd\nvrdear\ndvrsen\nenarar"))
	assert.Equal(t, "ikerpcty", solvePartOne(util.ReadFile("input")))
}

func TestSolvePartTwo(t *testing.T) {
	assert.Equal(t, "advent", solvePartTwo("eedadn\ndrvtee\neandsr\nraavrd\natevrs\ntsrnev\nsdttsa\nrasrtv\nnssdts\nntnada\nsvetve\ntesnvt\nvntsnd\nvrdear\ndvrsen\nenarar"))
	assert.Equal(t, "uwpfaqrq", solvePartTwo(util.ReadFile("input")))
}
