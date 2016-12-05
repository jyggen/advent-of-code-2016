package util

import (
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"path/filepath"
	"sort"
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

func RankByLetterCount(letterFrequencies map[string]int) PairList {
	pl := make(PairList, len(letterFrequencies))
	i := 0
	for k, v := range letterFrequencies {
		pl[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(pl))

	return pl

}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool {
	if p[i].Value == p[j].Value {
		return p[i].Key > p[j].Key
	}

	return p[i].Value < p[j].Value
}

func (p PairList) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

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
