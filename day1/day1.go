/*
	Day 1: Calorie Counting
	Link: https://adventofcode.com/2022/day/1
*/

package day1

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInput() ([][]int, error) {
	inputPath := "./day1/input.txt"

	content, err := os.ReadFile(inputPath)

	if err != nil {
		return nil, err
	}

	totalElvesCaloriesArrayString := strings.Split(string(content), "\n\n")
	var totalElvesCaloriesArrayInt [][]int

	for _, element := range totalElvesCaloriesArrayString {

		singleElfCaloriesArrayString := strings.Split(element, "\n")
		var singleElfCaloriesArrayInt []int

		for _, element := range singleElfCaloriesArrayString {

			if element == "" {
				continue
			}

			num, err := strconv.Atoi(element)

			if err != nil {
				return nil, err
			}

			singleElfCaloriesArrayInt = append(singleElfCaloriesArrayInt, num)
		}

		totalElvesCaloriesArrayInt = append(totalElvesCaloriesArrayInt, singleElfCaloriesArrayInt)
	}

	return totalElvesCaloriesArrayInt, nil
}

func Part1() (int, error) {

	input, err := getInput()

	if err != nil {
		return 0, err
	}

	greatestSum := 0

	for _, singleElfCaloriesArray := range input {

		singleElfCaloriesSum := 0

		for _, calories := range singleElfCaloriesArray {
			singleElfCaloriesSum += calories
		}

		if singleElfCaloriesSum > greatestSum {
			greatestSum = singleElfCaloriesSum
		}

	}

	return greatestSum, nil
}

func Part2() (int, error) {

	input, err := getInput()

	if err != nil {
		return 0, err
	}

	var sumCaloriesArray []int

	for _, singleElfCaloriesArray := range input {

		singleElfCaloriesSum := 0

		for _, calories := range singleElfCaloriesArray {
			singleElfCaloriesSum += calories
		}

		sumCaloriesArray = append(sumCaloriesArray, singleElfCaloriesSum)
	}

	sumSlice := sumCaloriesArray[:]
	sort.Sort(sort.Reverse(sort.IntSlice(sumSlice)))

	topThreeCalories := sumCaloriesArray[:3]

	topThreeSum := 0
	for _, element := range topThreeCalories {
		topThreeSum += element
	}

	return topThreeSum, nil
}
