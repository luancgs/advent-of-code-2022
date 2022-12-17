/*
	Day 5: Supply Stacks
	Link: https://adventofcode.com/2022/day/5
*/

package day05

import (
	S "adventofcode2022/utils/stack"
	"os"
	"strconv"
	"strings"
)

func getInput() (string, string, error) {
	inputPath := "./day5/input.txt"

	content, err := os.ReadFile(inputPath)

	if err != nil {
		return "", "", err
	}

	splitContent := strings.Split(string(content), "\n\n")

	return splitContent[0], splitContent[1], nil
}

func getInputStacks() ([]S.Stack, error) {
	inputStacks, _, err := getInput()

	if err != nil {
		return nil, err
	}

	inputStacksLines := strings.Split(inputStacks, "\n")

	for i, j := 0, len(inputStacksLines)-1; i < j; i, j = i+1, j-1 {
		inputStacksLines[i], inputStacksLines[j] = inputStacksLines[j], inputStacksLines[i]
	}

	stackCounterArray := strings.Split(inputStacksLines[0], "")

	var stackArray []S.Stack

	for _, stackNum := range stackCounterArray {
		if stackNum != " " {
			stackArray = append(stackArray, S.NewStack())
		}
	}

	for _, stacksLine := range inputStacksLines[1:] {
		for i, j := 1, 0; i < len(stacksLine); i, j = i+4, j+1 {
			crate := string(stacksLine[i])
			if crate != " " {
				stackArray[j].Push(crate)
			}
		}
	}

	return stackArray, nil
}

func getInputMovements() ([][]int, error) {
	_, inputMovements, err := getInput()

	if err != nil {
		return nil, err
	}

	inputMovementsLines := strings.Split(inputMovements, "\n")

	var movementsArray [][]int

	for _, movementsLine := range inputMovementsLines {

		splitMovements := strings.Split(movementsLine, " ")

		if len(splitMovements) != 6 {
			continue
		}

		numberCrates, err := strconv.Atoi(splitMovements[1])

		if err != nil {
			return nil, err
		}

		initalStack, err := strconv.Atoi(splitMovements[3])

		if err != nil {
			return nil, err
		}

		finalStack, err := strconv.Atoi(splitMovements[5])

		if err != nil {
			return nil, err
		}

		movementsArray = append(movementsArray, []int{numberCrates, initalStack - 1, finalStack - 1})
	}

	return movementsArray, nil
}

func Part1() (string, error) {

	stacks, err := getInputStacks()

	if err != nil {
		return "", err
	}

	movements, err := getInputMovements()

	if err != nil {
		return "", err
	}

	for _, movement := range movements {
		for i := 0; i < movement[0]; i++ {

			if stacks[movement[1]].Empty() {
				break
			}

			crate, err := stacks[movement[1]].Pop()

			if err != nil {
				return "", nil
			}

			stacks[movement[2]].Push(crate)
		}
	}

	var sb strings.Builder

	for _, stack := range stacks {

		topCrate, err := stack.Peek()

		if err != nil {
			return "", nil
		}

		sb.WriteString(topCrate.(string))
	}

	return sb.String(), nil
}

func Part2() (string, error) {
	stacks, err := getInputStacks()

	if err != nil {
		return "", err
	}

	movements, err := getInputMovements()

	if err != nil {
		return "", err
	}

	for _, movement := range movements {

		if stacks[movement[1]].Empty() {
			continue
		}

		crates, err := stacks[movement[1]].PopMany(movement[0])

		if err != nil {
			return "", nil
		}

		for i := len(crates) - 1; i >= 0; i-- {
			stacks[movement[2]].Push(crates[i])
		}
	}

	var sb strings.Builder

	for _, stack := range stacks {

		topCrate, err := stack.Peek()

		if err != nil {
			return "", nil
		}

		sb.WriteString(topCrate.(string))
	}

	return sb.String(), nil

}
