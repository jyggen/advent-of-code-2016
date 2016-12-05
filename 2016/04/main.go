package main

import (
	"github.com/jyggen/advent-of-go/util"
	"regexp"
	"strconv"
	"strings"
)

func solvePartOne(input string) int {
	sectorSum := 0
	rooms := strings.Split(input, "\n")
	regex := regexp.MustCompile("^(([a-z-]+))([\\d]+)\\[([a-z]+)\\]$")

	for _, room := range rooms {
		match := regex.FindStringSubmatch(room)
		letters := strings.Split(strings.Replace(match[2], "-", "", -1), "")
		sectorId, _ := strconv.Atoi(match[3])
		checksum := strings.Split(match[4], "")
		letterCount := make(map[string]int)

		for _, letter := range letters {
			if _, ok := letterCount[letter]; ok == false {
				letterCount[letter] = 0
			}

			letterCount[letter]++
		}

		mostFrequent := util.RankByLetterCount(letterCount)[0:5]
		isValid := true

		for index, letter := range checksum {
			if letter != mostFrequent[index].Key {
				isValid = false
				break
			}
		}

		if isValid {
			sectorSum += sectorId
		}
	}

	return sectorSum
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func solvePartTwo(input string) int {
	rooms := strings.Split(input, "\n")
	regex := regexp.MustCompile("^(([a-z-]+))([\\d]+)\\[([a-z]+)\\]$")

	for _, room := range rooms {
		match := regex.FindStringSubmatch(room)
		fullString := strings.Replace(match[2], "-", "", -1)
		letters := strings.Split(fullString, "")
		sectorId, _ := strconv.Atoi(match[3])
		checksum := strings.Split(match[4], "")
		letterCount := make(map[string]int)

		for _, letter := range letters {
			if _, ok := letterCount[letter]; ok == false {
				letterCount[letter] = 0
			}

			letterCount[letter]++
		}

		mostFrequent := util.RankByLetterCount(letterCount)[0:5]
		isValid := true

		for index, letter := range checksum {
			if letter != mostFrequent[index].Key {
				isValid = false
				break
			}
		}

		if isValid {
			var rot int = sectorId % 26
			b := []byte(fullString)
			for i, c := range b {
				c |= 0x20
				if 'a' <= c && c <= 'z' {
					b[i] = alphabet[(int(('z'-'a'+1)+(c-'a'))+rot)%26]
				}
			}

			message := string(b)

			if message == "northpoleobjectstorage" {
				return sectorId
			}
		}
	}

	return -1
}

func main() {
	input := util.ReadFile("2016/04/input")
	part1 := solvePartOne(input)
	part2 := solvePartTwo(input)

	util.PrintAnswers(strconv.Itoa(part1), strconv.Itoa(part2))
}
