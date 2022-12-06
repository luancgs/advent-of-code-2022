/*
	Day 6: Tuning Trouble
	Link: https://adventofcode.com/2022/day/6
*/

package day6

import (
	"os"
	"strings"
)

func getInput() (string, error) {
	inputPath := "./day6/input.txt"

	content, err := os.ReadFile(inputPath)

	if err != nil {
		return "", err
	}

	trimmedContent := strings.Trim(string(content), "\n")

	return trimmedContent, nil
}

func getIndexOfUniqueChars(input string, chars int) int {
	var index int
	for i, j := 0, chars; j < len(input); i, j = i+1, j+1 {
		possiblePackage := input[i:j]
		hasUniqueChars := true

		for _, char := range possiblePackage {
			copyPackage := strings.ReplaceAll(possiblePackage, string(char), "")

			if len(copyPackage) < chars-1 {
				hasUniqueChars = false
				break
			}
		}

		if hasUniqueChars {
			index = j
			break
		}
	}

	return index
}

func Part1() (int, error) {
	input, err := getInput()

	if err != nil {
		return 0, err
	}

	return getIndexOfUniqueChars(input, 4), nil
}

func Part2() (int, error) {
	input, err := getInput()

	if err != nil {
		return 0, err
	}

	return getIndexOfUniqueChars(input, 14), nil
}
