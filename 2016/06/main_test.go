package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePartOne(t *testing.T) {
	input := `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

	assert.Equal(t, "easter", solvePartOne(input))
}

func TestSolvePartTwo(t *testing.T) {
	input := `eedadn
drvtee
eandsr
raavrd
atevrs
tsrnev
sdttsa
rasrtv
nssdts
ntnada
svetve
tesnvt
vntsnd
vrdear
dvrsen
enarar`

	assert.Equal(t, "advent", solvePartTwo(input))
}
