package main

import (
	"github.com/jyggen/advent-of-go/util"
	"sort"
	"strings"
)

func solvePartOne(input string) string {
	rows := make(map[int]map[int]string)
	password := ""

	for column, letters := range strings.Split(input, "\n") {
		for row, letter := range strings.Split(letters, "") {
			if _, ok := rows[row]; ok == false {
				rows[row] = map[int]string{}
			}

			rows[row][column] = letter
		}
	}

	numOfRows := len(rows)

	for i := 0; i < numOfRows; i++ {
		letterCount := make(map[string]int)

		for _, letter := range rows[i] {
			if _, ok := letterCount[letter]; ok == false {
				letterCount[letter] = 0
			}

			letterCount[letter]++
		}

		password += util.RankByLetterCount(letterCount)[0].Key
	}

	return password
}

func solvePartTwo(input string) string {
	rows := make(map[int]map[int]string)
	password := ""

	for column, letters := range strings.Split(input, "\n") {
		for row, letter := range strings.Split(letters, "") {
			if _, ok := rows[row]; ok == false {
				rows[row] = map[int]string{}
			}

			rows[row][column] = letter
		}
	}

	numOfRows := len(rows)

	for i := 0; i < numOfRows; i++ {
		letterCount := make(map[string]int)

		for _, letter := range rows[i] {
			if _, ok := letterCount[letter]; ok == false {
				letterCount[letter] = 0
			}

			letterCount[letter]++
		}

		mostFrequent := util.RankByLetterCount(letterCount)

		sort.Sort(mostFrequent)

		password += mostFrequent[0].Key
	}

	return password
}

func main() {
	input := util.ReadFile("2016/06/input")
	part1 := solvePartOne(input)
	part2 := solvePartTwo(input)

	util.PrintAnswers(part1, part2)
}
