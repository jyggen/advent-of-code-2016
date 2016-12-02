package main

import (
	"bytes"
	"github.com/jyggen/advent-of-go/util"
	"strings"
)

func solve(input string, keypad map[int]map[int]string, x int, y int) string {
	var code bytes.Buffer

	for _, row := range strings.Split(input, "\n") {
		for _, instruction := range strings.Split(row, "") {
			switch instruction {
			case "U":
				if _, ok := keypad[y-1][x]; ok {
					y--
				}
				break
			case "R":
				if _, ok := keypad[y][x+1]; ok {
					x++
				}
				break
			case "D":
				if _, ok := keypad[y+1][x]; ok {
					y++
				}
				break
			case "L":
				if _, ok := keypad[y][x-1]; ok {
					x--
				}
				break
			}
		}

		code.WriteString(keypad[y][x])
	}

	return code.String()
}

func solvePartOne(input string) string {
	keypad := make(map[int]map[int]string)
	keypad = map[int]map[int]string{
		0: {0: "1", 1: "2", 2: "3"},
		1: {0: "4", 1: "5", 2: "6"},
		2: {0: "7", 1: "8", 2: "9"},
	}

	return solve(input, keypad, 1, 1)
}

func solvePartTwo(input string) string {
	keypad := make(map[int]map[int]string)
	keypad = map[int]map[int]string{
		0: {2: "1"},
		1: {1: "2", 2: "3", 3: "4"},
		2: {0: "5", 1: "6", 2: "7", 3: "8", 4: "9"},
		3: {1: "A", 2: "B", 3: "C"},
		4: {2: "D"},
	}

	return solve(input, keypad, 0, 2)
}

func main() {
	input := util.ReadFile("2016/02/input")
	part1 := solvePartOne(input)
	part2 := solvePartTwo(input)

	util.PrintAnswers(part1, part2)
}
