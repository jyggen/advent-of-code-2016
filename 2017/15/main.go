package main

import (
	"github.com/jyggen/advent-of-go/util"
	"regexp"
	"strconv"
)

func parseInput(input string) []int {
	regex := regexp.MustCompile("([\\d]+)")
	matches := regex.FindAllStringSubmatch(input, -1)
	generators := make([]int, 2)

	generators[0], _ = strconv.Atoi(matches[0][0])
	generators[1], _ = strconv.Atoi(matches[1][0])

	return generators
}

func solvePartOne(generators []int) int {
	score := 0

	for i := 0; i < 40000000; i++ {
		generators[0] = generators[0] * 16807 % 2147483647
		generators[1] = generators[1] * 48271 % 2147483647

		if generators[0]&0xffff == generators[1]&0xffff {
			score++
		}
	}

	return score
}

func solvePartTwo(generators []int) int {
	score := 0

	for i := 0; i < 5000000; i++ {
		for {
			generators[0] = generators[0] * 16807 % 2147483647

			if generators[0]%4 == 0 {
				break
			}
		}

		for {
			generators[1] = generators[1] * 48271 % 2147483647

			if generators[1]%8 == 0 {
				break
			}
		}

		if generators[0]&0xffff == generators[1]&0xffff {
			score++
		}
	}

	return score
}

func main() {
	generators := parseInput(util.ReadFile("2017/15/input"))

	util.StartBenchmark()

	result := solvePartOne(generators)

	util.StopBenchmark()
	util.PrintAnswer(result)
	util.StartBenchmark()

	result = solvePartTwo(generators)

	util.StopBenchmark()
	util.PrintAnswer(result)
}
