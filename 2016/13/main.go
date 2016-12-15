package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/util"
	"sort"
	"strconv"
)

type State struct {
	best int
	y    int
	x    int
}

var (
	states []State
)

func printState(maze [][]bool) {
	max := len(maze)

	for y := 0; y < max; y++ {
		for x := 0; x < max; x++ {
			if maze[y][x] {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func validPaths(maze [][]bool) int {
	max := len(maze)
	valid := 0

	for y := 0; y < max; y++ {
		for x := 0; x < max; x++ {
			if maze[y][x] {
				valid++
			}
		}
	}

	return valid
}

func solve(input int, goalX int, goalY int) (int, int) {
	bitOne := []rune("1")[0]
	max := 50
	maze := make([][]bool, max)

	for y := 0; y < max; y++ {
		maze[y] = make([]bool, max)

		for x := 0; x < max; x++ {
			result := x*x + 3*x + 2*x*y + y + y*y + input
			binary := strconv.FormatInt(int64(result), 2)
			oneBits := 0

			for _, bit := range binary {
				if bit == bitOne {
					oneBits++
				}
			}

			if oneBits%2 == 0 {
				maze[y][x] = true
			} else {
				maze[y][x] = false
			}
		}
	}

	states = []State{}
	best := walkPart1(maze, 0, 1, 1, goalY, goalX)

	states = []State{}
	walkPart2(maze, 0, 1, 1)

	return best, len(states)
}

func copyMaze(maze [][]bool) [][]bool {
	newMaze := make([][]bool, len(maze))

	for i := range newMaze {
		newMaze[i] = make([]bool, len(maze[i]))
		copy(newMaze[i], maze[i])
	}

	return newMaze
}

func walkPart1(maze [][]bool, steps int, y int, x int, goalY int, goalX int) int {
	if y == goalY && x == goalX {
		return steps
	}

	found := false

	for _, state := range states {
		if state.y == y && state.x == x {
			if state.best < steps {
				return -1
			}

			found = true
			break
		}
	}

	if !found {
		states = append(states, State{best: steps, y: y, x: x})
	}

	max := len(maze)

	maze[y][x] = false
	allBest := []int{}

	// North
	if y-1 >= 0 && maze[y-1][x] {
		allBest = append(allBest, walkPart1(copyMaze(maze), steps+1, y-1, x, goalY, goalX))
	}

	// East
	if x+1 < max && maze[y][x+1] {
		allBest = append(allBest, walkPart1(copyMaze(maze), steps+1, y, x+1, goalY, goalX))
	}

	// South
	if y+1 < max && maze[y+1][x] {
		allBest = append(allBest, walkPart1(copyMaze(maze), steps+1, y+1, x, goalY, goalX))
	}

	// West
	if x-1 >= 0 && maze[y][x-1] {
		allBest = append(allBest, walkPart1(copyMaze(maze), steps+1, y, x-1, goalY, goalX))
	}

	filteredBest := allBest[:0]

	for _, best := range allBest {
		if best != -1 {
			filteredBest = append(filteredBest, best)
		}
	}

	if len(filteredBest) == 0 {
		return -1
	}

	sort.Ints(filteredBest)

	return filteredBest[0]
}

func walkPart2(maze [][]bool, steps int, y int, x int) {
	if steps > 50 {
		return
	}

	found := false

	for _, state := range states {
		if state.y == y && state.x == x {
			if state.best < steps {
				return
			}

			found = true
			break
		}
	}

	if !found {
		states = append(states, State{best: steps, y: y, x: x})
	}

	max := len(maze)

	maze[y][x] = false

	// North
	if y-1 >= 0 && maze[y-1][x] {
		walkPart2(copyMaze(maze), steps+1, y-1, x)
	}

	// East
	if x+1 < max && maze[y][x+1] {
		walkPart2(copyMaze(maze), steps+1, y, x+1)
	}

	// South
	if y+1 < max && maze[y+1][x] {
		walkPart2(copyMaze(maze), steps+1, y+1, x)
	}

	// West
	if x-1 >= 0 && maze[y][x-1] {
		walkPart2(copyMaze(maze), steps+1, y, x-1)
	}
}

func main() {
	input, err := strconv.Atoi(util.ReadFile("2016/13/input"))

	if err != nil {
		panic(err)
	}

	util.StartBenchmark()

	numberOfSteps, numberOfLocations := solve(input, 31, 39)

	util.StopBenchmark()
	util.PrintAnswers(numberOfSteps, numberOfLocations)
}
