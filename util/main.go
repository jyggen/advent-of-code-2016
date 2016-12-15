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
	blue = color.New(color.FgBlue).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
	magentaBold = color.New(color.FgMagenta, color.Bold).SprintFunc()
	stop = time.Now()
	start = time.Now()
)

func PrintAnswer(answer interface{}) {
	fmt.Printf("[%s] %s\n", blue(roundN(stop.Sub(start), 2)), green(answer))
}

func PrintAnswers(answer interface{}, answer2 interface{}) {
	fmt.Printf("Part 1 returned %s and Part 2 returned %s after %s.\n", green(answer), green(answer2), blue(roundN(stop.Sub(start), 2)))
}

func StartBenchmark() {
	start = time.Now()
}

func StopBenchmark() {
	stop = time.Now()
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
