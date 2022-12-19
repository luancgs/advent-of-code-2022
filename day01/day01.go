/*
	Day 1: Calorie Counting
	Link: https://adventofcode.com/2022/day/1
*/

package day01

import (
	"adventofcode2022/utils/input"
	"sort"
	"strconv"
	"strings"
)

func formatInput() ([][]int, error) {
	content := input.GetInput(1)

	elvesCaloriesSlice := strings.Split(content, "\n\n")
	var elvesCaloriesSliceInt [][]int

	for _, elfCalories := range elvesCaloriesSlice {
		elfCaloriesSlice := strings.Split(elfCalories, "\n")
		var elfCaloriesInt []int

		for _, calories := range elfCaloriesSlice {
			if calories == "" {
				continue
			}

			num, err := strconv.Atoi(calories)

			if err != nil {
				return nil, err
			}

			elfCaloriesInt = append(elfCaloriesInt, num)
		}

		elvesCaloriesSliceInt = append(elvesCaloriesSliceInt, elfCaloriesInt)
	}

	return elvesCaloriesSliceInt, nil
}

func Part1() (int, error) {
	input, err := formatInput()

	if err != nil {
		return 0, err
	}

	greatestSum := 0

	for _, elfCalories := range input {
		elfCaloriesSum := 0

		for _, calories := range elfCalories {
			elfCaloriesSum += calories
		}

		if elfCaloriesSum > greatestSum {
			greatestSum = elfCaloriesSum
		}
	}

	return greatestSum, nil
}

func Part2() (int, error) {

	input, err := formatInput()

	if err != nil {
		return 0, err
	}

	var sumCaloriesSlice []int

	for _, elfCalories := range input {
		elfCaloriesSum := 0

		for _, calories := range elfCalories {
			elfCaloriesSum += calories
		}

		sumCaloriesSlice = append(sumCaloriesSlice, elfCaloriesSum)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sumCaloriesSlice)))

	topThreeCalories := sumCaloriesSlice[:3]
	topThreeSum := 0

	for _, element := range topThreeCalories {
		topThreeSum += element
	}

	return topThreeSum, nil
}
