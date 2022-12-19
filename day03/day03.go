/*
	Day 3: Rucksack Reorganization
	Link: https://adventofcode.com/2022/day/3
*/

package day03

import (
	"os"
	"strings"
)

const allItems = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func getInput() ([]string, [][]string, error) {
	inputPath := "./inputs/2022/day03.txt"

	content, err := os.ReadFile(inputPath)

	if err != nil {
		return nil, nil, err
	}

	rucksacksArray := strings.Split(string(content), "\n")
	var splitedRucksacksArray [][]string
	var filteredRucksacksArray []string

	for _, rucksack := range rucksacksArray {
		if rucksack == "" {
			continue
		}

		filteredRucksacksArray = append(filteredRucksacksArray, rucksack)

		division := len(rucksack) / 2
		splitedRucksacksArray = append(splitedRucksacksArray, []string{rucksack[:division], rucksack[division:]})
	}

	return filteredRucksacksArray, splitedRucksacksArray, nil
}

func Part1() (int, error) {

	_, input, err := getInput()

	if err != nil {
		return 0, err
	}

	prioritySum := 0

	for _, rucksack := range input {
		firstCompartment := strings.Split(rucksack[0], "")
		secondCompartment := rucksack[1]

		for _, item := range firstCompartment {
			if strings.Contains(secondCompartment, item) {
				prioritySum += (strings.Index(allItems, item) + 1)
				break
			}
		}
	}

	return prioritySum, nil
}

func Part2() (int, error) {
	input, _, err := getInput()

	if err != nil {
		return 0, err
	}

	prioritySum := 0
	count := 0

	for count < len(input) {
		var group []string

		for i := 0; i < 3; i++ {
			group = append(group, input[count])
			count++
		}

		firstRucksack := strings.Split(group[0], "")
		secondRucksack := group[1]
		thirdRucksack := group[2]

		for _, item := range firstRucksack {
			if strings.Contains(secondRucksack, item) && strings.Contains(thirdRucksack, item) {
				prioritySum += (strings.Index(allItems, item) + 1)
				break
			}
		}

	}

	return prioritySum, nil
}
