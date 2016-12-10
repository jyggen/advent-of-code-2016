package main

import (
	"fmt"
	"github.com/jyggen/advent-of-go/util"
	"regexp"
	"strconv"
	"strings"
)

type Bot struct {
	id     int
	values [2]int
}

type Output struct {
	id     int
	values []int
}

var bots map[int]*Bot
var outputs map[int]*Output

func botFactory(id int) *Bot {
	if bots == nil {
		bots = make(map[int]*Bot)
	}

	if _, ok := bots[id]; !ok {
		bots[id] = &Bot{id: id}
		fmt.Printf("Bot %d created.\n", id)
	}

	return bots[id]
}

func outputFactory(id int) *Output {
	if outputs == nil {
		outputs = make(map[int]*Output)
	}

	if _, ok := outputs[id]; !ok {
		outputs[id] = &Output{id: id}
		fmt.Printf("Output %d created.\n", id)
	}

	return outputs[id]
}

func (givingBot *Bot) GiveToBot(receivingBot *Bot) {
	if givingBot.values[1] == 0 || (givingBot.values[0] != 0 && givingBot.values[0] < givingBot.values[1]) {
		receivingBot.Receive(givingBot.values[0])
		givingBot.values[0] = 0
	} else {
		receivingBot.Receive(givingBot.values[1])
		givingBot.values[1] = 0
	}
}

func (bot *Bot) GiveToOutput(output *Output) {
	if bot.values[1] == 0 || (bot.values[0] != 0 && bot.values[0] < bot.values[1]) {
		output.Receive(bot.values[0])
		bot.values[0] = 0
	} else {
		output.Receive(bot.values[1])
		bot.values[1] = 0
	}
}

func (bot *Bot) Receive(value int) {
	if bot.values[0] == 0 {
		bot.values[0] = value
	} else if bot.values[1] == 0 {
		bot.values[1] = value
	} else {
		panic(fmt.Sprintf("Bot %d received too many chips!", bot.id))
	}

	fmt.Printf("Bot %d received %d (holding %v).\n", bot.id, value, bot.values)
}

func (output *Output) Receive(value int) {
	output.values = append(output.values, value)

	fmt.Printf("Output %d received %d (holding %v).\n", output.id, value, output.values)
}

func solve(input string, compare1 int, compare2 int) (int, int) {
	botRegex := regexp.MustCompile("^bot ([\\d]+) gives low to (bot|output) ([\\d]+) and high to (bot|output) ([\\d]+)$")
	valueRegex := regexp.MustCompile("^value ([\\d]+) goes to bot ([\\d]+)$")
	instructions := strings.Split(input, "\n")
	partOne := 0

	for i := 0; i < len(instructions); i++ {
		fmt.Printf("Step %d of %d: %s\n", i+1, len(instructions), instructions[i])

		match := valueRegex.FindStringSubmatch(instructions[i])

		if len(match) != 0 {
			value, _ := strconv.Atoi(match[1])
			receivingBotId, _ := strconv.Atoi(match[2])
			receivingBot := botFactory(receivingBotId)

			receivingBot.Receive(value)

			if (receivingBot.values[0] == compare1 || receivingBot.values[0] == compare2) && (receivingBot.values[1] == compare1 || receivingBot.values[1] == compare2) {
				partOne = receivingBot.id
			}
		} else {
			match := botRegex.FindStringSubmatch(instructions[i])

			if len(match) == 0 {
				panic(fmt.Sprintf("unable to parse instruction \"%s\"", instructions[i]))
			}

			givingBotId, _ := strconv.Atoi(match[1])
			givingBot := botFactory(givingBotId)

			if givingBot.values[0] == 0 || givingBot.values[1] == 0 {
				continue
			}

			receiver1Id, _ := strconv.Atoi(match[3])
			receiver2Id, _ := strconv.Atoi(match[5])

			if match[2] == "bot" {
				receivingBot := botFactory(receiver1Id)
				givingBot.GiveToBot(receivingBot)

				if (receivingBot.values[0] == compare1 || receivingBot.values[0] == compare2) && (receivingBot.values[1] == compare1 || receivingBot.values[1] == compare2) {
					partOne = receivingBot.id
				}
			} else {
				givingBot.GiveToOutput(outputFactory(receiver1Id))
			}

			if match[4] == "bot" {
				receivingBot := botFactory(receiver2Id)
				givingBot.GiveToBot(receivingBot)

				if (receivingBot.values[0] == compare1 || receivingBot.values[0] == compare2) && (receivingBot.values[1] == compare1 || receivingBot.values[1] == compare2) {
					partOne = receivingBot.id
				}
			} else {
				givingBot.GiveToOutput(outputFactory(receiver2Id))
			}
		}

		instructions = append(instructions[:i], instructions[i+1:]...)
		i = -1
	}

	output0 := outputFactory(0)
	output1 := outputFactory(1)
	output2 := outputFactory(2)

	return partOne, output0.values[0] * output1.values[0] * output2.values[0]
}

func main() {
	input := util.ReadFile("2016/10/input")
	part1, part2 := solve(input, 61, 17)

	util.PrintAnswers(strconv.Itoa(part1), strconv.Itoa(part2))
}
