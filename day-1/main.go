package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2022/day/1

func main() {
	output1, err := part1()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 1 - Part 1 | Output: ", output1)

	output2, err := part2()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Day 1 - Part 2 | Output: ", output2)

}

func part1() (int, error) {
	inputName := "./day-1/input.txt"

	content, err := os.ReadFile(inputName)

	if err != nil {
		return 0, err
	}

	totalCaloriesList := strings.Split(string(content), "\n\n")
	greatestNum := 0

	for _, element := range totalCaloriesList {
		singleCaloriesList := strings.Split(element, "\n")
		singleCaloriesSum := 0
		for _, element := range singleCaloriesList {

			if element == "" {
				continue
			}

			num, err := strconv.Atoi(element)

			if err != nil {
				return 0, err
			}

			singleCaloriesSum += num
		}

		if singleCaloriesSum > greatestNum {
			greatestNum = singleCaloriesSum
		}
	}

	return greatestNum, nil
}

func part2() (int, error) {
	inputName := "./day-1/input.txt"

	content, err := os.ReadFile(inputName)

	if err != nil {
		return 0, err
	}

	totalCaloriesList := strings.Split(string(content), "\n\n")
	var sumCaloriesList []int

	for _, element := range totalCaloriesList {
		singleCaloriesList := strings.Split(element, "\n")
		singleCaloriesSum := 0
		for _, element := range singleCaloriesList {

			if element == "" {
				continue
			}

			num, err := strconv.Atoi(element)

			if err != nil {
				return 0, err
			}

			singleCaloriesSum += num
		}

		sumCaloriesList = append(sumCaloriesList, singleCaloriesSum)
	}

	sumSlice := sumCaloriesList[:]
	sort.Sort(sort.Reverse(sort.IntSlice(sumSlice)))

	topThreeCalories := sumCaloriesList[:3]
	sum := 0
	for _, element := range topThreeCalories {
		sum += element
	}

	return sum, nil
}
