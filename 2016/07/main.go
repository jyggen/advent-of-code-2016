package main

import (
	"github.com/jyggen/advent-of-go/util"
	"strconv"
	"strings"
)

func solvePartOne(input string) int {
	valid := 0

	for _, address := range strings.Split(input, "\n") {
		characters := []rune(address)
		addressLength := len(characters)
		openBracket := []rune("[")[0]
		closeBracket := []rune("]")[0]
		isSquareBracket := false
		isValid := false
		validInSquareBracket := false

		for index, character := range characters {
			if character == openBracket {
				isSquareBracket = true
				continue
			}

			if character == closeBracket {
				isSquareBracket = false
				continue
			}

			if index+3+1 > addressLength {
				continue
			}

			first := characters[index : index+2]
			second := characters[index+2 : index+4]

			if first[0] == first[1] {
				continue
			}

			if first[0] == second[1] && first[1] == second[0] {
				if isSquareBracket == true {
					validInSquareBracket = true
				} else {
					isValid = true
				}
			}
		}

		if validInSquareBracket == true {
			continue
		}

		if validInSquareBracket == false && isValid == true {
			valid++
		}
	}

	return valid
}

func solvePartTwo(input string) int {
	valid := 0

	for _, address := range strings.Split(input, "\n") {
		characters := []rune(address)
		addressLength := len(characters)
		openBracket := []rune("[")[0]
		closeBracket := []rune("]")[0]
		isSquareBracket := false

		var ABA [][]rune
		var BAB [][]rune

		for index, character := range characters {
			if character == openBracket {
				isSquareBracket = true
				continue
			}

			if character == closeBracket {
				isSquareBracket = false
				continue
			}

			if index+2+1 > addressLength {
				continue
			}

			sequence := characters[index : index+3]

			if sequence[0] == sequence[2] && sequence[0] != sequence[1] {
				if isSquareBracket == true {
					BAB = append(BAB, sequence)
				} else {
					ABA = append(ABA, sequence)
				}
			}
		}

	Loop:
		for _, abaSequence := range ABA {
			for _, babSequence := range BAB {
				if babSequence[0] == abaSequence[1] && babSequence[1] == abaSequence[0] {
					valid++
					break Loop
				}
			}
		}
	}

	return valid
}

func main() {
	input := util.ReadFile("2016/07/input")
	part1 := solvePartOne(input)
	part2 := solvePartTwo(input)

	util.PrintAnswers(strconv.Itoa(part1), strconv.Itoa(part2))
}
