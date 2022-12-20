/*
	Day 6: Tuning Trouble
	Link: https://adventofcode.com/2022/day/6
*/

package day06

import (
	"adventofcode2022/utils/input"
	"strings"
)

func formatInput(reader input.InputReader) (string, error) {
	content := reader.GetInput(6)

	trimmedContent := strings.Trim(content, "\n")

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

func Part1(reader input.InputReader) (int, error) {
	input, err := formatInput(reader)

	if err != nil {
		return 0, err
	}

	return getIndexOfUniqueChars(input, 4), nil
}

func Part2(reader input.InputReader) (int, error) {
	input, err := formatInput(reader)

	if err != nil {
		return 0, err
	}

	return getIndexOfUniqueChars(input, 14), nil
}
