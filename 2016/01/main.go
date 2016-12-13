package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/util"
	"math"
	"strconv"
	"strings"
)

type Santa struct {
	facing            int
	firstVisitedTwice string
	history           map[string]int
	x                 int
	y                 int
}

func (santa *Santa) GetBlocksAway() int {
	x := math.Abs(float64(santa.x))
	y := math.Abs(float64(santa.y))

	return int(x) + int(y)
}

func (santa *Santa) GetFirstVisitedTwice() int {
	var blocksAway int

	if santa.firstVisitedTwice == "" {
		return -1
	}

	for _, coordString := range strings.Split(santa.firstVisitedTwice, "x") {
		coord, err := strconv.Atoi(coordString)

		if err != nil {
			panic(err)
		}

		blocksAway += int(math.Abs(float64(coord)))
	}

	return blocksAway
}

func (santa *Santa) Move(instruction string) {
	direction := instruction[:1]
	steps, err := strconv.Atoi(instruction[1:])

	if err != nil {
		panic(err)
	}

	if direction == "R" {
		santa.facing++

		if santa.facing == 4 {
			santa.facing = 0
		}
	} else if direction == "L" {
		santa.facing--

		if santa.facing == -1 {
			santa.facing = 3
		}
	} else {
		panic(fmt.Sprintf("unknown direction \"%s\"", direction))
	}

	for i := 1; i <= steps; i++ {
		switch santa.facing {
		case 0:
			santa.y--
		case 1:
			santa.x++
		case 2:
			santa.y++
		case 3:
			santa.x--
		default:
			panic(fmt.Sprintf("santa is facing a weird way (\"%d\")", santa.facing))
		}

		coords := fmt.Sprintf("%dx%d", santa.x, santa.y)

		if _, ok := santa.history[coords]; ok == false {
			santa.history[coords] = 0
		}

		santa.history[coords]++

		if santa.firstVisitedTwice == "" && santa.history[coords] == 2 {
			santa.firstVisitedTwice = coords
		}
	}
}

func parseInput(input string) ([]string) {
	return strings.Split(input, ", ")
}

func solve(instructions []string) (int, int) {
	santa := Santa{
		facing:            0,
		firstVisitedTwice: "",
		history:           make(map[string]int),
		x:                 0,
		y:                 0,
	}

	for _, instruction := range instructions {
		santa.Move(instruction)
	}

	return santa.GetBlocksAway(), santa.GetFirstVisitedTwice()
}

func main() {
	instructions := parseInput(util.ReadFile("2016/01/input"))

	util.StartBenchmark()

	blocksAway, firstRepeat := solve(instructions)

	util.StopBenchmark()
	util.PrintAnswers(blocksAway,firstRepeat)
}
