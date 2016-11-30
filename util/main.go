package util

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

var (
	green = color.New(color.FgGreen).SprintFunc()
	start = time.Now()
)

func PrintAnswers(part1 string, part2 string) {
	fmt.Printf("Part 1: %s\n", green(part1))
	fmt.Printf("Part 2: %s\n", green(part2))
	fmt.Print("\n")
	fmt.Printf("Executed in %s\n", green(time.Since(start)))
}

func ReadFile(path string) string {
	path, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(data))
}
