package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/util"
	"regexp"
	"strconv"
	"strings"
)

func solve(input string, registers map[string]int) int {
	instructions := strings.Split(input, "\n")
	numOfInstructions := len(instructions)
	parser := regexp.MustCompile("^([a-z]{3}) ([-a-z\\d]+)( ([-a-z\\d])+)?$")

	for index := 0; index < numOfInstructions; index++ {
		match := parser.FindStringSubmatch(instructions[index])

		if len(match) == 0 {
			panic(fmt.Sprintf("Unable to parse \"%s\"", instructions[index]))
		}

		switch(match[1]) {
		case "cpy":
			intValue, err := strconv.Atoi(match[2])

			if _, ok := registers[match[4]]; !ok {
				registers[match[4]] = 0
			}

			if (err != nil) {
				if _, ok := registers[match[2]]; !ok {
					registers[match[2]] = 0
				}

				registers[match[4]] = registers[match[2]]
			} else {
				registers[match[4]] = intValue
			}
			break
		case "inc":
			if _, ok := registers[match[2]]; !ok {
				registers[match[2]] = 0
			}

			registers[match[2]]++
			break
		case "dec":
			if _, ok := registers[match[2]]; !ok {
				registers[match[2]] = 0
			}

			registers[match[2]]--
			break
		case "jnz":
			match[2] = strings.TrimSpace(match[2])
			intValue, err := strconv.Atoi(match[2])

			if (err != nil) {
				if _, ok := registers[match[2]]; !ok {
					registers[match[2]] = 0
				}

				if (registers[match[2]] == 0) {
					continue
				}
			} else {
				if (intValue == 0) {
					continue
				}
			}

			match[3] = strings.TrimSpace(match[3])
			intValue2, err := strconv.Atoi(match[3])

			if (err != nil) {
				panic(err)
			}

			index = index + intValue2 - 1
			break
		}
	}

	return registers["a"]
}

func main() {
	input := util.ReadFile("2016/12/input")
	registers := make(map[string]int)
	part1 := solve(input, registers)
	registers = make(map[string]int)
	registers["c"] = 1
	part2 := solve(input, registers)

	util.PrintAnswers(strconv.Itoa(part1), strconv.Itoa(part2))
}
