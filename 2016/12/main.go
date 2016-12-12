package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/util"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(input string) [][]interface{} {
	instructions := strings.Split(input, "\n")
	parser := regexp.MustCompile("^([a-z]{3}) ([-a-z\\d]+)( ([-a-z\\d])+)?$")
	parsedInstructions := make([][]interface{}, len(instructions))

	for index, instruction := range instructions {
		byteInstruction := make([]interface{}, 0)
		match := parser.FindStringSubmatch(instruction)

		if len(match) == 0 {
			panic(fmt.Sprintf("Unable to parse \"%s\"", instruction))
		}

		byteInstruction = append(byteInstruction, match[1])

		switch(match[1]) {
		case "cpy":
			match[2] = strings.TrimSpace(match[2])
			intValue, err := strconv.Atoi(match[2])

			if (err != nil) {
				byteInstruction = append(byteInstruction, match[2])
			} else {
				byteInstruction = append(byteInstruction, intValue)
			}

			byteInstruction = append(byteInstruction, match[4])
			break
		case "dec":
			byteInstruction = append(byteInstruction, match[2])
			break
		case "inc":
			byteInstruction = append(byteInstruction, match[2])
			break
		case "jnz":
			match[2] = strings.TrimSpace(match[2])
			intValue, err := strconv.Atoi(match[2])

			if (err != nil) {
				byteInstruction = append(byteInstruction, match[2])
			} else {
				byteInstruction = append(byteInstruction, intValue)
			}

			match[3] = strings.TrimSpace(match[3])
			intValue2, _ := strconv.Atoi(match[3])

			byteInstruction = append(byteInstruction, intValue2)
			break
		}

		parsedInstructions[index] = byteInstruction
	}

	return parsedInstructions
}

func solve(instructions [][]interface{}, registers map[string]int) int {
	numOfInstructions := len(instructions)

	for index := 0; index < numOfInstructions; index++ {
		instruction := instructions[index]

		switch(instruction[0]) {
		case "cpy":
			if _, ok := registers[instruction[2].(string)]; !ok {
				registers[instruction[2].(string)] = 0
			}

			if _, ok := instruction[1].(int); ok {
				registers[instruction[2].(string)] = instruction[1].(int)
			} else {
				if _, ok := registers[instruction[1].(string)]; !ok {
					registers[instruction[1].(string)] = 0
				}

				registers[instruction[2].(string)] = registers[instruction[1].(string)]
			}
			break
		case "dec":
			if _, ok := registers[instruction[1].(string)]; !ok {
				registers[instruction[1].(string)] = 0
			}

			registers[instruction[1].(string)]--
			break
		case "inc":
			if _, ok := registers[instruction[1].(string)]; !ok {
				registers[instruction[1].(string)] = 0
			}

			registers[instruction[1].(string)]++
			break
		case "jnz":
			if _, ok := instruction[1].(int); ok {
				if (instruction[1].(int) == 0) {
					continue
				}
			} else {
				if _, ok := registers[instruction[1].(string)]; !ok {
					registers[instruction[1].(string)] = 0
				}

				if (registers[instruction[1].(string)] == 0) {
					continue
				}
			}

			index = index + instruction[2].(int) - 1
			break
		}
	}

	return registers["a"]
}

func main() {
	instructions := parseInput(util.ReadFile("2016/12/input"))

	registers := make(map[string]int)
	part1 := solve(instructions, registers)

	registers = make(map[string]int)
	registers["c"] = 1
	part2 := solve(instructions, registers)

	util.PrintAnswers(strconv.Itoa(part1), strconv.Itoa(part2))
}
