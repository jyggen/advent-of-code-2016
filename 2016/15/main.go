package main

import (
	"github.com/jyggen/advent-of-go/util"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	positions int
	current   int
}

func parseInput(input string) []Instruction {
	instructions := strings.Split(input, "\n")
	instructionStructs := make([]Instruction, len(instructions))
	parser := regexp.MustCompile("^Disc #([\\d]+) has ([\\d]+) positions; at time=0, it is at position ([\\d]+).$")

	for index, instruction := range instructions {
		match := parser.FindStringSubmatch(instruction)
		positions, _ := strconv.Atoi(match[2])
		current, _ := strconv.Atoi(match[3])
		instructionStructs[index] = Instruction{
			positions: positions,
			current:   current,
		}
	}

	return instructionStructs
}

func solve(instructions []Instruction) int {
	time := 0

	for {
		through := true

		for index, instruction := range instructions {
			currentTime := time + index + 1
			currentPosition := currentTime + instruction.current
			currentPosition %= instruction.positions

			if currentPosition != 0 {
				through = false
				break
			}
		}

		if through {
			break
		}

		time++
	}

	return time
}

func main() {
	instructions := parseInput(util.ReadFile("2016/15/input"))

	util.StartBenchmark()

	result := solve(instructions)

	util.StopBenchmark()
	util.PrintAnswer(result)
	util.StartBenchmark()

	instructions = append(instructions, Instruction{
		positions: 11,
		current:   0,
	})

	result = solve(instructions)

	util.StopBenchmark()
	util.PrintAnswer(result)
}
