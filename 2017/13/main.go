package main

import (
	"github.com/jyggen/advent-of-go/util"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	rows := strings.Split(input, "\n")
	length, _ := strconv.Atoi(strings.Split(rows[len(rows)-1], ": ")[0])
	firewall := make([]int, length+1)

	for _, row := range rows {
		parts := strings.Split(row, ": ")
		layer, _ := strconv.Atoi(parts[0])
		depth, _ := strconv.Atoi(parts[1])
		firewall[layer] = depth*2 - 2
	}

	return firewall
}

func sendPackage(firewall []int, allowDetection bool, delay int) (bool, int) {
	length := len(firewall) - 1
	current := -1
	severity := 0

	for current < length {
		current++

		if firewall[current] == 0 {
			continue
		}

		position := current + delay
		position %= firewall[current]

		if position != 0 {
			continue
		}

		severity += current * (firewall[current] + 2) / 2

		if allowDetection == false {
			return false, severity
		}
	}

	return true, severity
}

func solvePartOne(firewall []int) int {
	_, severity := sendPackage(firewall, true, 0)

	return severity
}

func solvePartTwo(firewall []int) int {
	sleep := 0
	success := false

	for success != true {
		success, _ = sendPackage(firewall, false, sleep)

		sleep++
	}

	return sleep - 1
}

func main() {
	firewall := parseInput(util.ReadFile("2017/13/input"))

	util.StartBenchmark()

	result := solvePartOne(firewall)

	util.StopBenchmark()
	util.PrintAnswer(result)
	util.StartBenchmark()

	result = solvePartTwo(firewall)

	util.StopBenchmark()
	util.PrintAnswer(result)
}
