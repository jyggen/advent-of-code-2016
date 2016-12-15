package main

import (
	"github.com/jyggen/advent-of-go/util"
	"strings"
)

func parseInput(input string) [][]rune {
	rows := strings.Split(input, "\n")
	instructions := make([][]rune, len(rows))

	for i, row := range rows {
		characters := strings.Split(row, "")
		instructions[i] = make([]rune, len(characters))

		for j, character := range characters {
			instructions[i][j] = []rune(character)[0]
		}
	}

	return instructions
}

func solve(instructions [][]rune, keypad [][]rune, y int, x int) string {
	code := make([]rune, len(instructions))
	maxX := len(keypad[0])
	maxY := len(keypad)
	runeDown := []rune("D")[0]
	runeLeft := []rune("L")[0]
	runeRight := []rune("R")[0]
	runeUp := []rune("U")[0]
	runeZero := []rune("0")[0]

	for index, row := range instructions {
		for _, instruction := range row {
			switch instruction {
			case runeUp:
				if y-1 >= 0 && keypad[y-1][x] != runeZero {
					y--
				}
				break
			case runeRight:
				if x+1 < maxX && keypad[y][x+1] != runeZero {
					x++
				}
				break
			case runeDown:
				if y+1 < maxY && keypad[y+1][x] != runeZero {
					y++
				}
				break
			case runeLeft:
				if x-1 >= 0 && keypad[y][x-1] != runeZero {
					x--
				}
				break
			}
		}

		code[index] = keypad[y][x]
	}

	return string(code)
}

func main() {
	instructions := parseInput(util.ReadFile("2016/02/input"))

	util.StartBenchmark()

	result := solve(instructions, [][]rune{
		0: {0: []rune("1")[0], 1: []rune("2")[0], 2: []rune("3")[0]},
		1: {0: []rune("4")[0], 1: []rune("5")[0], 2: []rune("6")[0]},
		2: {0: []rune("7")[0], 1: []rune("8")[0], 2: []rune("9")[0]},
	}, 1, 1)

	util.StopBenchmark()
	util.PrintAnswer(result)
	util.StartBenchmark()

	result = solve(instructions, [][]rune{
		0: {0: []rune("0")[0], 1: []rune("0")[0], 2: []rune("1")[0], 3: []rune("0")[0], 4: []rune("0")[0]},
		1: {0: []rune("0")[0], 1: []rune("2")[0], 2: []rune("3")[0], 3: []rune("4")[0], 4: []rune("0")[0]},
		2: {0: []rune("5")[0], 1: []rune("6")[0], 2: []rune("7")[0], 3: []rune("8")[0], 4: []rune("9")[0]},
		3: {0: []rune("0")[0], 1: []rune("A")[0], 2: []rune("B")[0], 3: []rune("C")[0], 4: []rune("0")[0]},
		4: {0: []rune("0")[0], 1: []rune("0")[0], 2: []rune("D")[0], 3: []rune("0")[0], 4: []rune("0")[0]},
	}, 2, 0)

	util.StopBenchmark()
	util.PrintAnswer(result)
}
