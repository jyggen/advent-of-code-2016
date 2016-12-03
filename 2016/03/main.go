package main

import (
	"github.com/jyggen/advent-of-go/util"
	"regexp"
	"strconv"
	"strings"
)

func solvePartOne(input string) int {
	possible := 0
	triangles := strings.Split(input, "\n")
	regex := regexp.MustCompile("[ ]+")

	for _, triangle := range triangles {
		sides := regex.Split(strings.TrimSpace(triangle), -1)
		side1, _ := strconv.Atoi(sides[0])
		side2, _ := strconv.Atoi(sides[1])
		side3, _ := strconv.Atoi(sides[2])

		if side1+side2 > side3 && side2+side3 > side1 && side3+side1 > side2 {
			possible++
		}
	}

	return possible
}

func solvePartTwo(input string) int {
	possible := 0
	triangles := strings.Split(input, "\n")
	sideNumber := 0
	index := 0
	regex := regexp.MustCompile("[ ]+")

	var sides [][]int

	for _, triangle := range triangles {
		numbers := regex.Split(strings.TrimSpace(triangle), -1)
		side1, _ := strconv.Atoi(numbers[0])
		side2, _ := strconv.Atoi(numbers[1])
		side3, _ := strconv.Atoi(numbers[2])

		if sideNumber == 0 {
			sideSlice1 := make([]int, 0)
			sideSlice2 := make([]int, 0)
			sideSlice3 := make([]int, 0)
			sides = append(sides, append(sideSlice1, side1), append(sideSlice2, side2), append(sideSlice3, side3))
		} else {
			sides[index] = append(sides[index], side1)
			sides[index+1] = append(sides[index+1], side2)
			sides[index+2] = append(sides[index+2], side3)
		}

		sideNumber++

		if sideNumber == 3 {
			index = index + 3
			sideNumber = 0
		}
	}

	for _, side := range sides {
		if side[0]+side[1] > side[2] && side[1]+side[2] > side[0] && side[2]+side[0] > side[1] {
			possible++
		}
	}

	return possible
}

func main() {
	input := util.ReadFile("2016/03/input")
	part1 := solvePartOne(input)
	part2 := solvePartTwo(input)

	util.PrintAnswers(strconv.Itoa(part1), strconv.Itoa(part2))
}
