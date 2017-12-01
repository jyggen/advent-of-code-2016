package main

import (
	"github.com/jyggen/advent-of-go/util"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	instructions := strings.Split(input, "")
	instructionInts := make([]int, len(instructions))

	for index, instruction := range instructions {
		instructionInts[index], _ = strconv.Atoi(instruction)
	}

	return instructionInts
}

func solvePartOne(instructions []int) int {
	total := 0

	for index, instruction := range instructions {
		previous := index - 1

		if (previous < 0) {
			previous = len(instructions) - 1
		}

		if (instructions[previous] != instruction) {
			continue
		}

		total += instruction
	}

	return total
}

func solvePartTwo(instructions []int) int {
	total := 0
	half := len(instructions) / 2

	for index, instruction := range instructions {
		previous := index - half

		if (previous < 0) {
			previous = len(instructions) - util.AbsInt(previous)
		}

		if (instructions[previous] != instruction) {
			continue
		}

		total += instruction
	}

	return total
}

func main() {
	instructions := parseInput(util.ReadFile("2017/01/input"))

	util.StartBenchmark()

	result := solvePartOne(instructions)

	util.StopBenchmark()
	util.PrintAnswer(result)
	util.StartBenchmark()

	result = solvePartTwo(instructions)

	util.StopBenchmark()
	util.PrintAnswer(result)
}
