package main

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"github.com/jyggen/advent-of-go/util"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	start = time.Now()
)

func solve(input string, width int, height int) (int, string) {
	rectRegex := regexp.MustCompile("^rect ([\\d]+)x([\\d]+)$")
	rotateRegex := regexp.MustCompile("^rotate (column|row) (y|x)=([\\d]+) by ([\\d]+)$")

	display := make([][]bool, width)

	for i := range display {
		display[i] = make([]bool, height)
	}

	for _, instruction := range strings.Split(input, "\n") {
		match := rectRegex.FindStringSubmatch(instruction)

		if len(match) != 0 {
			maxX, _ := strconv.Atoi(match[1])
			maxY, _ := strconv.Atoi(match[2])

			for x := 0; x < maxX; x++ {
				for y := 0; y < maxY; y++ {
					display[x][y] = true
				}
			}
		} else {

			match = rotateRegex.FindStringSubmatch(instruction)

			if len(match) == 0 {
				panic(fmt.Sprintf("Unable to parse instruction \"%s\"", instruction))
			}

			row, _ := strconv.Atoi(match[3])
			shift, _ := strconv.Atoi(match[4])

			if match[2] == "x" {
				stateTracker := make([]bool, height)

				for y := 0; y < height; y++ {
					newY := y - (shift % height)

					if newY < 0 {
						newY += height
					}

					stateTracker[y] = display[row][newY]
				}

				for index, value := range stateTracker {
					display[row][index] = value
				}
			} else if match[2] == "y" {
				stateTracker := make([]bool, width)

				for x := 0; x < width; x++ {

					newX := x - (shift % width)

					if newX < 0 {
						newX += width
					}

					stateTracker[x] = display[newX][row]
				}

				for index, value := range stateTracker {
					display[index][row] = value
				}
			} else {
				panic(fmt.Sprintf("Unable to parse instruction \"%s\"", instruction))
			}
		}
	}

	var ascii bytes.Buffer
	enabled := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if display[x][y] {
				enabled++
				ascii.WriteString("#")
			} else {
				ascii.WriteString(" ")
			}
		}

		ascii.WriteString("\n")
	}

	return enabled, ascii.String()
}

func main() {
	input := util.ReadFile("2016/08/input")
	part1, part2 := solve(input, 50, 6)

	fmt.Printf("Part 1: %s\n", green(part1))
	fmt.Printf("Part 2:\n%s\n", green(part2))
	fmt.Printf("Executed in %s\n", green(time.Since(start)))
}
