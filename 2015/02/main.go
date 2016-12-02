package main

import (
	"github.com/jyggen/advent-of-go/util"
	"sort"
	"strconv"
	"strings"
)

func solve(input string) (int, int) {
	var presents [][]int

	for _, present := range strings.Split(input, "\n") {
		var dimensions []int
		dimensionStrings := strings.Split(present, "x")

		for _, value := range dimensionStrings {
			intValue, _ := strconv.Atoi(value)

			dimensions = append(dimensions, intValue)
		}

		sort.Ints(dimensions) // lazy

		presents = append(presents, dimensions)
	}

	paperSize := 0
	ribbonSize := 0

	for _, present := range presents {
		sides := []int{
			present[0] * present[1],
			present[1] * present[2],
			present[2] * present[0],
		}

		paperSize += sides[0]*2 + sides[1]*2 + sides[2]*2
		paperSize += sides[0]
		ribbonSize += present[0]*2 + present[1]*2
		ribbonSize += present[0] * present[1] * present[2]
	}

	return paperSize, ribbonSize
}

func main() {
	part1, part2 := solve(util.ReadFile("2015/02/input"))

	util.PrintAnswers(strconv.Itoa(part1), strconv.Itoa(part2))
}
