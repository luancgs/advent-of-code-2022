/*
	Day 8: Treetop Tree House
	Link: https://adventofcode.com/2022/day/8
*/

package day08

import (
	"os"
	"strconv"
	"strings"
)

func getInput() ([][]int, error) {
	inputPath := "./inputs/2022/day08.txt"

	input, err := os.ReadFile(inputPath)

	if err != nil {
		return nil, err
	}

	treeLines := strings.Split(string(input), "\n")
	var treeGrid [][]int

	for _, treeLine := range treeLines {

		if len(treeLine) == 0 {
			continue
		}

		treeArray := strings.Split(treeLine, "")
		var treeLineHeight []int

		for _, tree := range treeArray {
			treeHeight, err := strconv.Atoi(tree)

			if err != nil {
				return nil, err
			}

			treeLineHeight = append(treeLineHeight, treeHeight)
		}

		treeGrid = append(treeGrid, treeLineHeight)
	}

	return treeGrid, nil
}

func treeVisibleFromOutsideGrid(tree int, treeGrid [][]int, line int, column int) bool {

	treeIsTallestUp := true
	treeIsTallestDown := true
	treeIsTallestLeft := true
	treeIsTallestRight := true

	// Up
	for i := 0; i < line; i++ {
		if treeGrid[i][column] >= tree {
			treeIsTallestUp = false
			break
		}
	}

	if treeIsTallestUp {
		return true
	}

	// Down
	for i := line + 1; i < len(treeGrid); i++ {
		if treeGrid[i][column] >= tree {
			treeIsTallestDown = false
			break
		}
	}

	if treeIsTallestDown {
		return true
	}

	// Left
	for j := 0; j < column; j++ {
		if treeGrid[line][j] >= tree {
			treeIsTallestLeft = false
			break
		}
	}

	if treeIsTallestLeft {
		return true
	}

	// Right
	for j := column + 1; j < len(treeGrid[line]); j++ {
		if treeGrid[line][j] >= tree {
			treeIsTallestRight = false
			break
		}
	}

	return treeIsTallestRight
}

func treeScenicScore(tree int, treeGrid [][]int, line int, column int) int {

	viewingDistanceUp := 0
	viewingDistanceDown := 0
	viewingDistanceLeft := 0
	viewingDistanceRight := 0

	viewingDistance := 0

	// Up
	for i := line - 1; i >= 0; i-- {
		viewingDistance++

		if treeGrid[i][column] >= tree {
			break
		}
	}

	viewingDistanceUp = viewingDistance
	viewingDistance = 0

	// Down
	for i := line + 1; i < len(treeGrid); i++ {
		viewingDistance++

		if treeGrid[i][column] >= tree {
			break
		}
	}

	viewingDistanceDown = viewingDistance
	viewingDistance = 0

	// Left
	for j := column - 1; j >= 0; j-- {
		viewingDistance++

		if treeGrid[line][j] >= tree {
			break
		}
	}

	viewingDistanceLeft = viewingDistance
	viewingDistance = 0

	// Right
	for j := column + 1; j < len(treeGrid[line]); j++ {
		viewingDistance++

		if treeGrid[line][j] >= tree {
			break
		}
	}

	viewingDistanceRight = viewingDistance

	return viewingDistanceUp * viewingDistanceDown * viewingDistanceLeft * viewingDistanceRight
}

func Part1() (int, error) {
	treeGrid, err := getInput()

	if err != nil {
		return 0, err
	}

	countVisible := 0

	for i, treeLine := range treeGrid {
		for j, tree := range treeLine {
			if i == 0 || i == len(treeGrid)-1 || j == 0 || j == len(treeLine)-1 {
				countVisible++
			} else if treeVisibleFromOutsideGrid(tree, treeGrid, i, j) {
				countVisible++
			}
		}
	}

	return countVisible, nil
}

func Part2() (int, error) {
	treeGrid, err := getInput()

	if err != nil {
		return 0, err
	}

	highestScenicScore := 0

	for i, treeLine := range treeGrid {
		for j, tree := range treeLine {
			scenicScore := treeScenicScore(tree, treeGrid, i, j)

			if scenicScore > highestScenicScore {
				highestScenicScore = scenicScore
			}
		}
	}

	return highestScenicScore, nil
}
