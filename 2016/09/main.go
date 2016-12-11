package main

import (
	"bytes"
	"github.com/jyggen/advent-of-go/util"
	"strconv"
	"strings"
)

func solvePartOne(input string) int {
	var currentMarker bytes.Buffer
	var decompressed bytes.Buffer
	var repeat int
	var subsequent int

	isDecompressing := false
	isParsingMarker := false

	for _, character := range strings.Split(input, "") {
		if isParsingMarker {
			if character == ")" {
				isDecompressing = true
				isParsingMarker = false
				marker := strings.Split(currentMarker.String(), "x")
				repeat, _ = strconv.Atoi(marker[1])
				subsequent, _ = strconv.Atoi(marker[0])
			} else {
				currentMarker.WriteString(character)
			}
		} else if isDecompressing {
			decompressed.WriteString(strings.Repeat(character, repeat))

			subsequent--

			if subsequent == 0 {
				isDecompressing = false
			}
		} else if character == "(" {
			currentMarker.Reset()
			isParsingMarker = true
		} else {
			decompressed.WriteString(character)
		}
	}

	return len(decompressed.String())
}

func solvePartTwo(input []rune) int {
	compressedLength := len(input)
	decompressedLength := 0
	markerStart := []rune("(")[0]
	markerDivider := []rune("x")[0]
	markerEnd := []rune(")")[0]

	for i := 0; i < compressedLength; {
		if input[i] == markerStart {
			hasHitDivider := false
			subsequentRunes := []rune{}
			repeatRunes := []rune{}

			var j int

			for j = i + 1; i < compressedLength; j++ {
				if input[j] == markerEnd {
					break
				} else if input[j] == markerDivider {
					hasHitDivider = true
				} else if hasHitDivider {
					repeatRunes = append(repeatRunes, input[j])
				} else {
					subsequentRunes = append(subsequentRunes, input[j])
				}
			}

			subsequent, _ := strconv.Atoi(string(subsequentRunes))
			repeat, _ := strconv.Atoi(string(repeatRunes))
			sliceStart := j + 1
			sliceEnd := sliceStart + subsequent

			decompressedLength += repeat * solvePartTwo(input[sliceStart:sliceEnd])
			i = sliceEnd
		} else {
			decompressedLength++
			i++
		}
	}

	return decompressedLength
}

func main() {
	input := util.ReadFile("2016/09/input")
	part1 := solvePartOne(input)
	part2 := solvePartTwo([]rune(input))

	util.PrintAnswers(strconv.Itoa(part1), strconv.Itoa(part2))
}
