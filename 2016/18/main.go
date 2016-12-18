package main

import (
	"github.com/jyggen/advent-of-go/util"
)

func solve(input string, rows int) int {
	layout := make([][]rune, rows)
	numOfTiles := len(input)
	safeRune := []rune(".")[0]
	trapRune := []rune("^")[0]

	for i := 0; i < rows; i++ {
		if i == 0 {
			layout[i] = append(layout[i], []rune(input)...)
			continue
		}

		layout[i] = make([]rune, numOfTiles)

		for index, _ := range layout[i] {
			var center rune
			var left rune
			var right rune

			if index == 0 {
				left = safeRune
			} else {
				left = layout[i-1][index-1]
			}

			center = layout[i-1][index]

			if index == numOfTiles-1 {
				right = safeRune
			} else {
				right = layout[i-1][index+1]
			}

			if left == trapRune && center == trapRune && right == safeRune {
				layout[i][index] = trapRune
			} else if left == safeRune && center == trapRune && right == trapRune {
				layout[i][index] = trapRune
			} else if left == trapRune && center == safeRune && right == safeRune {
				layout[i][index] = trapRune
			} else if left == safeRune && center == safeRune && right == trapRune {
				layout[i][index] = trapRune
			} else {
				layout[i][index] = safeRune
			}
		}
	}

	safe := 0

	for _, row := range layout {
		for _, tile := range row {
			if tile == safeRune {
				safe++
			}
		}
	}

	return safe
}

func main() {
	input := util.ReadFile("2016/18/input")

	util.StartBenchmark()

	result := solve(input, 40)

	util.StopBenchmark()
	util.PrintAnswer(result)
	util.StartBenchmark()

	result = solve(input, 400000)

	util.StopBenchmark()
	util.PrintAnswer(result)
}
