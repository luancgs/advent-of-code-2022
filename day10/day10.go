/*
	Day 10: Cathode-Ray Tube
	Link: https://adventofcode.com/2022/day/10
*/

package day10

import (
	"os"
	"strconv"
	"strings"
)

func getInput() ([][]int, error) {
	inputPath := "./day10/input.txt"

	input, err := os.ReadFile(inputPath)

	if err != nil {
		return nil, err
	}

	commands := strings.Split(string(input), "\n")

	var output [][]int

	for _, command := range commands {
		if command == "" {
			continue
		}

		splitCommand := strings.Split(command, " ")

		if len(splitCommand) == 2 {
			value, err := strconv.Atoi(splitCommand[1])

			if err != nil {
				return nil, err
			}

			output = append(output, []int{1, value})
		} else {
			output = append(output, []int{0, 0})
		}
	}

	return output, nil
}

func cyclePart1(cpuCycle int, cpuRegister int, signalSum *int) {
	if cpuCycle == 20 || cpuCycle == 60 || cpuCycle == 100 || cpuCycle == 140 || cpuCycle == 180 || cpuCycle == 220 {
		*signalSum += (cpuCycle * cpuRegister)
	}
}

func cyclePart2(cpuCycle int, cpuRegister int, screen *[][]string) {
	screenLine := (cpuCycle - 1) / 40
	screenColumn := (cpuCycle - 1) % 40

	cycleScreen := *screen

	if screenColumn == cpuRegister-1 || screenColumn == cpuRegister || screenColumn == cpuRegister+1 {
		cycleScreen[screenLine][screenColumn] = "#"
	}
}

func Part1() (int, error) {
	input, err := getInput()

	if err != nil {
		return 0, err
	}

	cpuRegister := 1
	cpuCycle := 1
	signalSum := 0

	for _, command := range input {
		if command[0] == 0 {
			cyclePart1(cpuCycle, cpuRegister, &signalSum)
			cpuCycle++
		} else {
			for i := 0; i < 2; i++ {
				cyclePart1(cpuCycle, cpuRegister, &signalSum)
				cpuCycle++
			}

			cpuRegister += command[1]
		}
	}

	return signalSum, nil
}

func Part2() (string, error) {
	input, err := getInput()

	if err != nil {
		return "", err
	}

	cpuRegister := 1
	cpuCycle := 1
	var screen [][]string

	for i := 0; i < 6; i++ {
		var initalScreenLine []string
		for i := 0; i < 40; i++ {
			initalScreenLine = append(initalScreenLine, ".")
		}
		screen = append(screen, initalScreenLine)
	}

	for _, command := range input {
		if command[0] == 0 {
			cyclePart2(cpuCycle, cpuRegister, &screen)
			cpuCycle++
		} else {
			for i := 0; i < 2; i++ {
				cyclePart2(cpuCycle, cpuRegister, &screen)
				cpuCycle++
			}

			cpuRegister += command[1]
		}
	}

	var outputLines []string

	for _, screenLine := range screen {
		outputLines = append(outputLines, strings.Join(screenLine, " "))
	}

	output := "\n" + strings.Join(outputLines, "\n")

	return output, nil
}
