/*
	Day 4: Camp Cleanup
	Link: https://adventofcode.com/2022/day/4
*/

package day04

import (
	"os"
	"strconv"
	"strings"
)

func getInput() ([][][]int, error) {
	inputPath := "./day4/input.txt"

	content, err := os.ReadFile(inputPath)

	if err != nil {
		return nil, err
	}

	pairAssignmentRangeArrayString := strings.Split(string(content), "\n")
	var pairAssignmentRangeArrayInt [][][]int

	for _, pairAssignmentRangeString := range pairAssignmentRangeArrayString {

		if pairAssignmentRangeString == "" {
			continue
		}

		singleAssignmentRangeArrayString := strings.Split(pairAssignmentRangeString, ",")
		var singleAssignmentRangeArrayInt [][]int

		for _, singleAssignmentRangeString := range singleAssignmentRangeArrayString {
			assignmentIdArrayString := strings.Split(singleAssignmentRangeString, "-")
			var assignmentIdArrayInt []int

			for _, assignmentIdString := range assignmentIdArrayString {

				num, err := strconv.Atoi(assignmentIdString)

				if err != nil {
					return nil, err
				}

				assignmentIdArrayInt = append(assignmentIdArrayInt, num)
			}

			singleAssignmentRangeArrayInt = append(singleAssignmentRangeArrayInt, assignmentIdArrayInt)
		}

		pairAssignmentRangeArrayInt = append(pairAssignmentRangeArrayInt, singleAssignmentRangeArrayInt)
	}

	return pairAssignmentRangeArrayInt, nil
}

func Part1() (int, error) {

	input, err := getInput()

	if err != nil {
		return 0, err
	}

	countContains := 0

	for _, pairAssignment := range input {

		firstContainsSecond := (pairAssignment[0][0] <= pairAssignment[1][0]) && (pairAssignment[0][1] >= pairAssignment[1][1])
		secondContainsFirst := (pairAssignment[1][0] <= pairAssignment[0][0]) && (pairAssignment[1][1] >= pairAssignment[0][1])

		if firstContainsSecond || secondContainsFirst {
			countContains++
		}
	}

	return countContains, nil
}

func Part2() (int, error) {
	input, err := getInput()

	if err != nil {
		return 0, err
	}

	countOverlaps := 0

	for _, pairAssignment := range input {

		firstOverlapsSecond := (pairAssignment[0][1] >= pairAssignment[1][0]) && (pairAssignment[0][0] <= pairAssignment[1][0])
		secondOverlapsFirst := (pairAssignment[1][1] >= pairAssignment[0][0]) && (pairAssignment[1][0] <= pairAssignment[0][0])

		if firstOverlapsSecond || secondOverlapsFirst {
			countOverlaps++
		}
	}

	return countOverlaps, nil
}
