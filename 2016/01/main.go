package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/util"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	direction rune
	distance  int
}

func getLocationKey(position []int) string {
	return fmt.Sprintf("%sx%s", position[0], position[1])
}

func parseInput(input string) []Instruction {
	instructions := strings.Split(input, ",")
	instructionStructs := make([]Instruction, len(instructions))
	instructionRegexp := regexp.MustCompile("^(L|R)([\\d]+)$")

	for index, instruction := range instructions {
		instruction = strings.TrimSpace(instruction)
		match := instructionRegexp.FindStringSubmatch(instruction)
		direction := []rune(match[1])[0]
		distance, _ := strconv.Atoi(match[2])

		instructionStructs[index] = Instruction{
			direction: direction,
			distance:  distance,
		}
	}

	return instructionStructs
}

func solvePartOne(instructions []Instruction) int {
	face := 0
	position := []int{0, 0}
	runeLeft := []rune("L")[0]

	for _, instruction := range instructions {
		if instruction.direction == runeLeft {
			face += 1
		} else {
			face += -1
		}

		face %= 4

		if face < 0 {
			face += 4
		}

		coord := face % 2
		dir := 0

		if face/2 == 0 {
			dir = 1
		} else {
			dir = -1
		}

		position[coord] += dir * instruction.distance
	}

	return util.AbsInt(position[0]) + util.AbsInt(position[1])
}

func solvePartTwo(instructions []Instruction) int {
	face := 0
	position := []int{0, 0}
	runeLeft := []rune("L")[0]
	locations := map[string]bool{
		getLocationKey(position): true,
	}

Loop:
	for _, instruction := range instructions {
		if instruction.direction == runeLeft {
			face += 1
		} else {
			face += -1
		}

		face %= 4

		if face < 0 {
			face += 4
		}

		coord := face % 2
		dir := 0

		if face/2 == 0 {
			dir = 1
		} else {
			dir = -1
		}

		for i := 0; i < instruction.distance; i++ {
			position[coord] += dir

			locationKey := getLocationKey(position)

			if _, ok := locations[locationKey]; ok {
				break Loop
			}

			locations[locationKey] = true
		}
	}

	return util.AbsInt(position[0]) + util.AbsInt(position[1])
}

func main() {
	instructions := parseInput(util.ReadFile("2016/01/input"))

	util.StartBenchmark()

	result := solvePartOne(instructions)

	util.StopBenchmark()
	util.PrintAnswer(result)
	util.StartBenchmark()

	result = solvePartTwo(instructions)

	util.StopBenchmark()
	util.PrintAnswer(result)
}
