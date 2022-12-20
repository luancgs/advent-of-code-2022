/*
	Day 12: Hill Climbing Algorithm
	Link: https://adventofcode.com/2022/day/12
*/

package day12

import (
	"adventofcode2022/utils/input"
	"fmt"
	"math"
	"sort"
	"strings"
)

type square struct {
	line             int
	column           int
	height           rune
	neighbors        map[string]*square
	previousSquare   *square
	shortestDistance int
}

func (s *square) setNeighbors(heightMap [][]square) {

	if s.line != 0 && s.canClimb(heightMap[s.line-1][s.column]) {
		s.neighbors["up"] = &heightMap[s.line-1][s.column]
	}
	if s.line != len(heightMap)-1 && s.canClimb(heightMap[s.line+1][s.column]) {
		s.neighbors["down"] = &heightMap[s.line+1][s.column]
	}
	if s.column != 0 && s.canClimb(heightMap[s.line][s.column-1]) {
		s.neighbors["left"] = &heightMap[s.line][s.column-1]
	}
	if s.column != len(heightMap[s.line])-1 && s.canClimb(heightMap[s.line][s.column+1]) {
		s.neighbors["right"] = &heightMap[s.line][s.column+1]
	}
}

func (s square) canClimb(nextSquare square) bool {
	return nextSquare.height <= s.height+1
}

func formatInput(reader input.InputReader) ([][]square, []int, []int, error) {
	content := reader.GetInput(12)

	contentLines := strings.Split(content, "\n")

	var heightMap [][]rune

	for _, line := range contentLines {

		if line == "" {
			continue
		}

		heightMap = append(heightMap, []rune(line))
	}

	var output [][]square
	var outputStart []int
	var outputEnd []int

	for line, heightLine := range heightMap {
		var outputLine []square
		for column, height := range heightLine {
			if height == 'S' {
				outputStart = []int{line, column}
				height = 'a'
			} else if height == 'E' {
				outputEnd = []int{line, column}
				height = 'z'
			}

			outputLine = append(outputLine, square{line, column, height, make(map[string]*square), nil, math.MaxInt})
		}

		output = append(output, outputLine)
	}

	return output, outputStart, outputEnd, nil
}

func Part1(reader input.InputReader) (int, error) {
	heightMap, startIndex, endIndex, err := formatInput(reader)

	if err != nil {
		return 0, err
	}

	for _, inputLine := range heightMap {
		for _, sqr := range inputLine {
			sqr.setNeighbors(heightMap)
		}
	}

	heightMap[startIndex[0]][startIndex[1]].shortestDistance = 0

	var unvisitedSquares []*square
	var visitedSquares []*square

	for i, inputLine := range heightMap {
		for j := range inputLine {
			unvisitedSquares = append(unvisitedSquares, &heightMap[i][j])
		}
	}

	squaresDistance := 1

	for len(unvisitedSquares) >= 1 {

		// Sorts unvisitedSquares from closest to furthest
		sort.Slice(unvisitedSquares, func(a, b int) bool {
			return (*unvisitedSquares[a]).shortestDistance < (*unvisitedSquares[b]).shortestDistance
		})

		//Pop current square
		currentSquare := unvisitedSquares[0]
		unvisitedSquares = unvisitedSquares[1:]

		for _, neighbor := range (*currentSquare).neighbors {

			// Check if neighbor is unvisited
			var isVisited = false
			for _, sqr := range visitedSquares {
				if sqr == neighbor {
					isVisited = true
					break
				}
			}

			if isVisited {
				continue
			}

			distanceFromStart := (*currentSquare).shortestDistance + squaresDistance

			if distanceFromStart < (*neighbor).shortestDistance {
				(*neighbor).shortestDistance = distanceFromStart
				(*neighbor).previousSquare = currentSquare
			}

			visitedSquares = append(visitedSquares, currentSquare)
		}

	}

	endSquare := &heightMap[endIndex[0]][endIndex[1]]

	return (*endSquare).shortestDistance, nil
}

func Part2(reader input.InputReader) (int, error) {
	heightMap, startIndex, endIndex, err := formatInput(reader)

	if err != nil {
		return 0, err
	}

	startSquare := &heightMap[startIndex[0]][startIndex[1]]

	for i, line := range heightMap {
		for j, height := range line {
			if height.height == 'a' {
				if i != startIndex[0] && j != startIndex[1] {
					coordinateString := fmt.Sprintf("%d,%d", i, j)
					(*startSquare).neighbors[coordinateString] = &heightMap[i][j]
				}
				heightMap[i][j].shortestDistance = 0
			}
		}
	}

	for _, inputLine := range heightMap {
		for _, sqr := range inputLine {
			sqr.setNeighbors(heightMap)
		}
	}

	var unvisitedSquares []*square
	var visitedSquares []*square

	for i, inputLine := range heightMap {
		for j := range inputLine {
			unvisitedSquares = append(unvisitedSquares, &heightMap[i][j])
		}
	}

	squaresDistance := 1

	for len(unvisitedSquares) >= 1 {

		// Sorts unvisitedSquares from closest to furthest
		sort.Slice(unvisitedSquares, func(a, b int) bool {
			return (*unvisitedSquares[a]).shortestDistance < (*unvisitedSquares[b]).shortestDistance
		})

		//Pop current square
		currentSquare := unvisitedSquares[0]
		unvisitedSquares = unvisitedSquares[1:]

		for _, neighbor := range (*currentSquare).neighbors {

			// Check if neighbor is unvisited
			var isVisited = false
			for _, sqr := range visitedSquares {
				if sqr == neighbor {
					isVisited = true
					break
				}
			}

			if isVisited {
				continue
			}

			distanceFromStart := (*currentSquare).shortestDistance + squaresDistance

			if distanceFromStart < (*neighbor).shortestDistance {
				(*neighbor).shortestDistance = distanceFromStart
				(*neighbor).previousSquare = currentSquare
			}

			visitedSquares = append(visitedSquares, currentSquare)
		}

	}

	endSquare := &heightMap[endIndex[0]][endIndex[1]]

	return (*endSquare).shortestDistance, nil
}
