/*
	Day 11: Monkey in the Middle
	Link: https://adventofcode.com/2022/day/11
*/

package day11

import (
	"adventofcode2022/utils/input"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type item struct {
	worryLevel int
}

type monkey struct {
	startingItems []item
	operation     []string
	test          int
	ifTrue        int
	ifFalse       int
	inspected     int
}

func (m monkey) operationWorryLevel(item item) (int, error) {
	num, err := strconv.Atoi(m.operation[1])

	if err != nil {
		return 0, err
	}

	if num == -1 {
		num = item.worryLevel
	}

	switch m.operation[0] {
	case "+":
		return item.worryLevel + num, nil
	case "-":
		return item.worryLevel - num, nil
	case "*":
		return item.worryLevel * num, nil
	case "/":
		return item.worryLevel / num, nil
	default:
		return 0, fmt.Errorf("error calculating operation in worry level")
	}
}

func throwItems(monkey int, monkeysPointer *[]monkey, reliefFactor bool) error {
	monkeys := *monkeysPointer

	for _, monkeyItem := range monkeys[monkey].startingItems {
		monkeys[monkey].inspected++

		worryLevel, err := monkeys[monkey].operationWorryLevel(monkeyItem)

		if err != nil {
			return err
		}

		if reliefFactor {
			worryLevel = int(math.Floor(float64(worryLevel) / 3))
		}

		monk := monkeys[monkey]

		if !reliefFactor {
			var module int = 1
			for _, mk := range monkeys {
				module *= mk.test
			}

			worryLevel %= module
		}

		if worryLevel%monkeys[monkey].test == 0 {
			monkeys[monk.ifTrue].startingItems = append(monkeys[monk.ifTrue].startingItems, item{worryLevel})
		} else {
			monkeys[monk.ifFalse].startingItems = append(monkeys[monk.ifFalse].startingItems, item{worryLevel})
		}
	}

	monkeys[monkey].startingItems = []item{}

	return nil
}

func formatInput(reader input.InputReader) ([]monkey, error) {
	content := reader.GetInput(11)

	inputMonkeys := strings.Split(content, "\n\n")

	var monkeys []monkey

	for _, inputMonkey := range inputMonkeys {
		inputMonkeyLines := strings.Split(inputMonkey, "\n")

		startingItemsString := strings.Split(strings.Trim(strings.ReplaceAll(inputMonkeyLines[1], "Starting items: ", ""), " "), ", ")
		var startingItems []item

		for _, itemString := range startingItemsString {
			num, err := strconv.Atoi(itemString)

			if err != nil {
				return nil, err
			}

			startingItems = append(startingItems, item{num})
		}

		operation := strings.Split(strings.Trim(strings.ReplaceAll(inputMonkeyLines[2], "Operation: new = old ", ""), " "), " ")

		if operation[1] == "old" {
			operation[1] = "-1"
		}

		testString := strings.Trim(strings.ReplaceAll(inputMonkeyLines[3], "Test: divisible by ", ""), " ")

		test, err := strconv.Atoi(testString)

		if err != nil {
			return nil, err
		}

		ifTrueString := strings.Trim(strings.ReplaceAll(inputMonkeyLines[4], "If true: throw to monkey ", ""), " ")

		ifTrue, err := strconv.Atoi(ifTrueString)

		if err != nil {
			return nil, err
		}

		ifFalseString := strings.Trim(strings.ReplaceAll(inputMonkeyLines[5], "If false: throw to monkey ", ""), " ")

		ifFalse, err := strconv.Atoi(ifFalseString)

		if err != nil {
			return nil, err
		}

		monkeys = append(monkeys, monkey{startingItems, operation, test, ifTrue, ifFalse, 0})
	}

	return monkeys, nil

}

func Part1(reader input.InputReader) (int, error) {
	monkeys, err := formatInput(reader)

	if err != nil {
		return 0, err
	}

	for round := 1; round <= 20; round++ {
		for monkey := 0; monkey < len(monkeys); monkey++ {
			throwItems(monkey, &monkeys, true)
		}
	}

	var monkeyInspects []int

	for _, monkey := range monkeys {
		monkeyInspects = append(monkeyInspects, monkey.inspected)
	}

	sort.Ints(monkeyInspects)
	monkeyInspects = monkeyInspects[len(monkeyInspects)-2:]

	return monkeyInspects[0] * monkeyInspects[1], nil
}

func Part2(reader input.InputReader) (int, error) {
	monkeys, err := formatInput(reader)

	if err != nil {
		return 0, err
	}

	var module int = 1
	for _, mk := range monkeys {
		module *= mk.test
	}

	for round := 1; round <= 10000; round++ {
		for monkey := 0; monkey < len(monkeys); monkey++ {
			throwItems(monkey, &monkeys, false)
		}
	}

	var monkeyInspects []int

	for _, monkey := range monkeys {
		monkeyInspects = append(monkeyInspects, monkey.inspected)
	}

	sort.Ints(monkeyInspects)
	monkeyInspects = monkeyInspects[len(monkeyInspects)-2:]

	return monkeyInspects[0] * monkeyInspects[1], nil
}
