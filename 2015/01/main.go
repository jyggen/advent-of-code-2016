package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/util"
	"strconv"
	"strings"
)

type Santa struct {
	currentPosition    int
	hasReachedBasement bool
	numberOfMoves      int
	reachedBasementAt  int
}

func (santa *Santa) Move(instruction string) {
	if instruction == "(" {
		santa.currentPosition++
	} else if instruction == ")" {
		santa.currentPosition--
	} else {
		panic(fmt.Sprintf("unknown instruction \"%s\"", instruction))
	}

	santa.numberOfMoves++

	if santa.currentPosition == -1 && santa.hasReachedBasement == false {
		santa.hasReachedBasement = true
		santa.reachedBasementAt = santa.numberOfMoves
	}
}

func solve(input string) (int, int) {
	instructions := strings.Split(input, "")
	santa := Santa{
		currentPosition:    0,
		hasReachedBasement: false,
		numberOfMoves:      0,
		reachedBasementAt:  0,
	}

	for _, instruction := range instructions {
		santa.Move(instruction)
	}

	return santa.currentPosition, santa.reachedBasementAt
}

func main() {
	part1, part2 := solve(util.ReadFile("2015/01/input"))

	util.PrintAnswers(strconv.Itoa(part1), strconv.Itoa(part2))
}
