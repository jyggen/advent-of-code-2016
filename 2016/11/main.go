package main

import (
	"crypto/md5"
	"fmt"
	"github.com/jyggen/advent-of-go/util"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type State struct {
	best     int
	elevator int
	floors   [][]string
}

var (
	bestSolution      int
	statesEncountered map[string]*State
)

func copyState(state [][]string) [][]string {
	newState := make([][]string, len(state))

	for i := range newState {
		newState[i] = make([]string, len(state[i]))
		copy(newState[i], state[i])
	}

	return newState
}

func parseInput(input string) [][]string {
	instructions := strings.Split(input, "\n")
	floors := make([][]string, len(instructions))
	generator := regexp.MustCompile("^([a-z]{1})[a-z]+ generator$")
	microship := regexp.MustCompile("^([a-z]{1})[a-z]+-compatible microchip$")
	redundant := regexp.MustCompile("^The [a-z]+ floor contains ([a ]+)?")

	for index, instruction := range instructions {
		instruction = redundant.ReplaceAllString(instruction, "")
		instruction = instruction[0 : len(instruction)-1]

		items := strings.Split(instruction, ", a ")
		items = append(items[0:len(items)-1], strings.Split(items[len(items)-1], ", and a ")...)

		if items[0] == "nothing relevant" {
			floors[index] = make([]string, 0)
			continue
		}

		floors[index] = make([]string, len(items))

		for jndex, item := range items {
			match := generator.FindStringSubmatch(item)

			if len(match) != 0 {
				item = fmt.Sprintf("%sG", match[1])
			} else {
				match = microship.FindStringSubmatch(item)

				if len(match) == 0 {
					panic(fmt.Sprintf("Unable to parse item \"%s\"", item))
				}

				item = fmt.Sprintf("%sM", match[1])
			}

			floors[index][jndex] = strings.ToUpper(item)
		}

		sort.Strings(floors[index])
	}

	return floors
}

func isNewOrBetterState(floors [][]string, elevator int, steps int) bool {
	stateId := getStateId(floors, elevator)

	if statesEncountered == nil {
		statesEncountered = make(map[string]*State)
	}

	if _, ok := statesEncountered[stateId]; ok && statesEncountered[stateId].best < steps {
		return false
	}

	statesEncountered[stateId] = &State{best: steps, elevator: elevator, floors: floors}

	return true
}

func isSuccessful(floors [][]string, elevator int) bool {
	numberOfFloors := len(floors) - 1

	if elevator != numberOfFloors {
		return false
	}

	for index, floor := range floors {
		if index != numberOfFloors && len(floor) != 0 {
			return false
		}
	}

	return true
}

func isValidState(floors [][]string) bool {
	for _, items := range floors {
		for _, item := range items {
			isDangerous := false

			if item[1:2] == "G" {
				continue
			}

			for _, otherItem := range items {
				if item == otherItem {
					continue
				}

				if item[0:1] == otherItem[0:1] {
					isDangerous = false
					break
				}

				if otherItem[1:2] == "G" {
					isDangerous = true
				}
			}

			if isDangerous {
				return false
			}
		}
	}

	return true
}

func getStateId(floors [][]string, elevator int) string {
	state := []rune{}

	state = append(state, []rune("E:")...)
	state = append(state, []rune(strconv.Itoa(elevator))...)
	anonMap := make(map[string]string)
	anonInt := 0

	for index, items := range floors {
		state = append(state, []rune("|F:")...)
		state = append(state, []rune(strconv.Itoa(index))...)
		state = append(state, []rune(":I:")...)

		//sort.Strings(items)

		itemState := []string{}

		for _, item := range items {
			if _, ok := anonMap[item[0:1]]; !ok {
				anonMap[item[0:1]] = strconv.Itoa(anonInt)
				anonInt++
			}

			itemState = append(itemState, fmt.Sprintf("%s%s", anonMap[item[0:1]], item[1:2]))
		}

		sort.Strings(itemState)

		for _, item := range itemState {
			state = append(state, []rune(item)...)
		}
	}

	return string(state)
	return fmt.Sprintf("%x", md5.Sum([]byte(string(state))))
}

func printState(floors [][]string, elevator int, steps int) {
	stateId := getStateId(floors, elevator)
	items := []string{}

	for _, floor := range floors {
		items = append(items, floor...)
	}

	sort.Strings(items)

	fmt.Printf("%s @ %d steps\n", stateId, steps)

	for i := len(floors) - 1; i >= 0; i-- {
		fmt.Printf("F%d ", i+1)

		if elevator == i {
			fmt.Print("E  ")
		} else {
			fmt.Print(".  ")
		}

		for _, item := range items {
			onThisFloor := false

			for _, floorItem := range floors[i] {
				if floorItem == item {
					onThisFloor = true
					break
				}
			}

			if onThisFloor {
				fmt.Printf("%s ", item)
			} else {
				fmt.Print(".  ")
			}
		}

		fmt.Println()
	}

	fmt.Println()
}

func solve(floors [][]string, elevator int, steps int) int {

	if bestSolution != 0 && steps > bestSolution {
		return -1
	}

	if !isNewOrBetterState(floors, elevator, steps) {
		return -1
	}

	if !isValidState(floors) {
		return -1
	}

	if isSuccessful(floors, elevator) {
		return steps
	}

	var possibilities [][]string

	itemsOnFloor := len(floors[elevator])
	allPossibleSteps := []int{}
	possibleFloors := []int{}

	if elevator != 0 {
		possibleFloors = append(possibleFloors, elevator-1)
	}

	if elevator != len(floors)-1 {
		possibleFloors = append(possibleFloors, elevator+1)
	}

	for _, newFloor := range possibleFloors {
		possibilities = make([][]string, 0)

		if newFloor < elevator {
			shouldMoveDown := false

			for i := newFloor; i >= 0; i-- {
				if len(floors[i]) > 0 {
					shouldMoveDown = true
					break
				}
			}

			if !shouldMoveDown {
				continue
			}
		}

		for index, item := range floors[elevator] {

			possibility := append([]string{}, item)
			possibilities = append(possibilities, possibility)

			for i := index + 1; i < itemsOnFloor; i++ {
				possibility = append([]string{}, item, floors[elevator][i])
				possibilities = append(possibilities, possibility)
			}
		}

		for _, possibility := range possibilities {

			newState := copyState(floors)

			newState[newFloor] = append(newState[newFloor], possibility...)

			if len(newState[elevator]) == len(possibility) {
				newState[elevator] = []string{}
			} else {
				for _, itemToRemove := range possibility {
					for i := 0; i < len(newState[elevator]); i++ {
						if newState[elevator][i] == itemToRemove {
							newState[elevator] = append(newState[elevator][:i], newState[elevator][i+1:]...)
						}
					}
				}
			}

			returnedSteps := solve(newState, newFloor, steps+1)
			allPossibleSteps = append(allPossibleSteps, returnedSteps)
		}
	}

	filteredSteps := allPossibleSteps[:0]

	for _, steps := range allPossibleSteps {
		if steps != -1 {
			filteredSteps = append(filteredSteps, steps)
		}
	}

	if len(filteredSteps) == 0 {
		return -1
	}

	sort.Ints(filteredSteps)

	if bestSolution == 0 || filteredSteps[0] < bestSolution {
		bestSolution = filteredSteps[0]
	}

	return filteredSteps[0]
}

func main() {
	floors := parseInput(util.ReadFile("2016/11/input"))
	part1 := -1
	bestSolution = 1
	statesEncountered = make(map[string]*State)
	part1 = solve(floors, 0, 0)

	for bestSolution = 1; part1 == -1; bestSolution++ {
		steps := bestSolution - 1

		for _, state := range statesEncountered {
			if state.best == steps {
				part1 = solve(state.floors, state.elevator, steps)

				if part1 != -1 {
					break
				}
			}
		}
	}

	floors[0] = append(floors[0], "EG", "EM", "DG", "DM")

	part2 := -1
	bestSolution = 1
	statesEncountered = make(map[string]*State)
	part2 = solve(floors, 0, 0)

	for bestSolution = 1; part2 == -1; bestSolution++ {
		steps := bestSolution - 1

		for _, state := range statesEncountered {
			if state.best == steps {
				part2 = solve(state.floors, state.elevator, steps)

				if part2 != -1 {
					break
				}
			}
		}
	}

	util.PrintAnswers(strconv.Itoa(part1), strconv.Itoa(part2))
}
