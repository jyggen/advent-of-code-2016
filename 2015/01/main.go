package main

import (
	"github.com/jyggen/advent-of-go/util"
	"strconv"
	"strings"
)

func solvePartOne(instructions string) int {
	up := len(strings.Replace(instructions, "(", "", -1))
	down := len(strings.Replace(instructions, ")", "", -1))

	return -up + down
}

func solvePartTwo(instructions string) int {
	var floor int

	for index, instruction := range instructions {
		if string(instruction) == "(" {
			floor++
		} else {
			floor--
		}

		if floor < 0 {
			return index + 1
		}
	}

	return -1
}

func main() {
	instructions := util.ReadFile("2015/01/input")
	partOne := solvePartOne(instructions)
	partTwo := solvePartTwo(instructions)

	util.PrintAnswers(strconv.Itoa(partOne), strconv.Itoa(partTwo))
}
