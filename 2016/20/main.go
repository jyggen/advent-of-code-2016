package main

import (
	"github.com/jyggen/advent-of-go/util"
	"strconv"
	"strings"
)

func normalizeRanges(ipRanges [][]int) [][]int {
	newRanges := make([][]int, 0)

	for _, ipRange := range ipRanges {
		normalized := false

		for index, newRange := range newRanges {
			// Range is already fully covered by the existing range.
			if ipRange[0] >= newRange[0] && ipRange[1] <= newRange[1] {
				normalized = true
			}

			// Range covers the existing range.
			if ipRange[0] <= newRange[0] && ipRange[1] >= newRange[1] {
				newRanges[index][0] = ipRange[0]
				newRanges[index][1] = ipRange[1]
				normalized = true
			}

			// Range is somewhat covered by the existing range.
			if ipRange[0] < newRange[0] && ipRange[1] >= newRange[0] {
				newRanges[index][0] = ipRange[0]
				normalized = true
			}

			// Range is somewhat covered by the existing range.
			if ipRange[0] <= newRange[1] && ipRange[1] > newRange[1] {
				newRanges[index][1] = ipRange[1]
				normalized = true
			}
		}

		if (normalized) {
			continue
		}

		newRanges = append(newRanges, []int{ipRange[0], ipRange[1]})
	}

	return newRanges
}

func parseInput(input string) [][]int {
	splitInput := strings.Split(input, "\n")
	ranges := make([][]int, len(splitInput))

	for index, ipRange := range splitInput {
		ranges[index] = make([]int, 2)
		parts := strings.Split(ipRange, "-")
		ranges[index][0], _ = strconv.Atoi(parts[0])
		ranges[index][1], _ = strconv.Atoi(parts[1])
	}

	for {
		preLength := len(ranges)
		ranges = normalizeRanges(ranges)
		postLength := len(ranges)

		if (preLength == postLength) {
			break
		}
	}

	return ranges
}

func solvePartOne(ipRanges [][]int) int {
	solution := 0

	for solution <= 4294967295 {
		okay := true

		for _, ipRange := range ipRanges {
			if (solution >= ipRange[0] && solution <= ipRange[1]) {
				okay = false
				solution = ipRange[1] + 1
				break
			}
		}

		if okay {
			break
		}
	}

	return solution

}

func solvePartTwo(ipRanges [][]int) int {
	solution := 0
	allowed := 0

	for solution <= 4294967295 {
		okay := true

		for _, ipRange := range ipRanges {
			if (solution >= ipRange[0] && solution <= ipRange[1]) {
				okay = false
				solution = ipRange[1] + 1
				break
			}
		}

		if okay {
			allowed++
			solution++
		}
	}

	return allowed
}

func main() {
	input := parseInput(util.ReadFile("2016/20/input"))

	util.StartBenchmark()

	result := solvePartOne(input)

	util.StopBenchmark()
	util.PrintAnswer(result)
	util.StartBenchmark()

	result = solvePartTwo(input)

	util.StopBenchmark()
	util.PrintAnswer(result)
}
