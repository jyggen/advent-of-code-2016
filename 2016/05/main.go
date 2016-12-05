package main

import (
	"crypto/md5"
	"fmt"
	"github.com/jyggen/advent-of-go/util"
	"strconv"
	"strings"
)

func solvePartOne(input string) string {
	index := 0
	password := ""

	for true {
		preHash := []byte(fmt.Sprintf("%s%d", input, index))
		checksum := fmt.Sprintf("%x", md5.Sum(preHash))

		if checksum[0:5] == "00000" {
			password += checksum[5:6]

			if len(password) == 8 {
				break
			}
		}

		index++
	}

	return password
}

func solvePartTwo(input string) string {
	password := make([]rune, 8)
	index := 0
	emptyRune := password[0]
	trim := string(emptyRune)

	for true {
		preHash := []byte(fmt.Sprintf("%s%d", input, index))
		checksum := fmt.Sprintf("%x", md5.Sum(preHash))
		position, err := strconv.Atoi(checksum[5:6])

		if err == nil && checksum[0:5] == "00000" && position <= 7 && password[position] == emptyRune {
			password[position] = []rune(checksum[6:7])[0]

			if len(strings.Replace(string(password), trim, "", -1)) == 8 {
				break
			}
		}

		index++
	}

	return string(password)
}

func main() {
	input := util.ReadFile("2016/05/input")
	part1 := solvePartOne(input)
	part2 := solvePartTwo(input)

	util.PrintAnswers(part1, part2)
}
