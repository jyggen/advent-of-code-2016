package main

import (
	"crypto/md5"
	"fmt"
	"github.com/jyggen/advent-of-go/util"
)

type Match struct {
	character rune
	index int
}

var (
	lookup = map[string]string{}
)

func getChecksum(message string) string {
	if _, ok := lookup[message]; ok {
		return lookup[message]
	}

	lookup[message] = fmt.Sprintf("%x", md5.Sum([]byte(message)))

	return lookup[message]
}

func solve(input string, stretching int) int {
	index := 0
	passwordsFound := 0
	matches := []Match{}

	for true {
		checksum := getChecksum(fmt.Sprintf("%s%d", input, index))

		for i := 0; i < stretching; i++ {
			checksum = getChecksum(checksum)
		}

		repeat := 1
		count := true

		var last rune;

		for _, character := range checksum {
			if (last == character) {
				repeat++

				if (repeat == 5) {
					for i := 0; i < len(matches); {
						if (matches[i].index < index - 1000) {
							matches = append(matches[:i], matches[i+1:]...)
							continue
						}

						if (matches[i].character == character && matches[i].index != index) {
							passwordsFound++

							if (passwordsFound == 64) {
								return matches[i].index
							}

							matches = append(matches[:i], matches[i+1:]...)
							continue
						}

						i++
					}

					break
				} else if (repeat == 3 && count) {
					matches = append(matches, Match{character: character, index: index})
					count = false
				}

				continue
			}

			last = character
			repeat = 1
		}

		index++
	}

	return -1
}

func main() {
	input := util.ReadFile("2016/14/input")

	util.StartBenchmark()

	result := solve(input, 0)

	util.StopBenchmark()
	util.PrintAnswer(result)
	util.StartBenchmark()

	result = solve(input, 2016)

	util.StopBenchmark()
	util.PrintAnswer(result)
}
