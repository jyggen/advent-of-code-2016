package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/util"
	"regexp"
	"strconv"
	"strings"
)

func getValue(unknownValue interface{}, registers map[string]int) int {
	if intValue, ok := unknownValue.(int); ok {
		return intValue
	}

	return registers[unknownValue.(string)]
}

func parseInput(input string) [][]interface{} {
	instructions := strings.Split(input, "\n")
	parser := regexp.MustCompile("^([a-z]{3}) ([-a-z\\d]+)( ([-a-z\\d])+)?$")
	parsedInstructions := make([][]interface{}, len(instructions))

	for index, instruction := range instructions {
		byteInstruction := make([]interface{}, 3)
		match := parser.FindStringSubmatch(instruction)

		if len(match) == 0 {
			panic(fmt.Sprintf("Unable to parse \"%s\"", instruction))
		}

		byteInstruction[0] = match[1]

		switch match[1] {
		case "cpy":
			match[2] = strings.TrimSpace(match[2])
			intValue, err := strconv.Atoi(match[2])

			if err != nil {
				byteInstruction[1] = match[2]
			} else {
				byteInstruction[1] = intValue
			}

			byteInstruction[2] = match[4]
			break
		case "dec":
			byteInstruction[1] = match[2]
			break
		case "inc":
			byteInstruction[1] = match[2]
			break
		case "jnz":
			match[2] = strings.TrimSpace(match[2])
			intValue, err := strconv.Atoi(match[2])

			if err != nil {
				byteInstruction[1] = match[2]
			} else {
				byteInstruction[1] = intValue
			}

			match[3] = strings.TrimSpace(match[3])
			intValue2, _ := strconv.Atoi(match[3])

			byteInstruction[2] = intValue2
			break
		}

		parsedInstructions[index] = byteInstruction
	}

	return parsedInstructions
}

func solve(instructions [][]interface{}, registers map[string]int) int {
	numOfInstructions := len(instructions)

	for i := 0; i < numOfInstructions; i++ {
		instruction := instructions[i]

		switch instruction[0] {
		case "cpy":
			registers[instruction[2].(string)] = getValue(instruction[1], registers)
			break
		case "dec":
			registers[instruction[1].(string)]--
			break
		case "inc":
			registers[instruction[1].(string)]++
			break
		case "jnz":
			if getValue(instruction[1], registers) == 0 {
				continue
			}

			i = i + instruction[2].(int) - 1
			break
		}
	}

	return registers["a"]
}

func main() {
	instructions := parseInput(util.ReadFile("2016/12/input"))

	util.StartBenchmark()

	result := solve(instructions, map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
	})

	util.StopBenchmark()
	util.PrintAnswer(1, result)
	util.StartBenchmark()

	result = solve(instructions, map[string]int{
		"a": 0,
		"b": 0,
		"c": 1,
		"d": 0,
	})

	util.StopBenchmark()
	util.PrintAnswer(2, result)
}
