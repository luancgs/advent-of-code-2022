/*
	Day 9: Rope Bridge
	Link: https://adventofcode.com/2022/day/9
*/

package day09

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type node struct {
	x, y int
}

func (n node) toString() string {
	return fmt.Sprintf("%d,%d", n.x, n.y)
}

type rope struct {
	nodes    []node
	tailPath map[string]bool
}

func (r rope) getNodesDistance(node int) float64 {
	return math.Sqrt(math.Pow(float64(r.nodes[node].x-r.nodes[node+1].x), 2) + math.Pow(float64(r.nodes[node].y-r.nodes[node+1].y), 2))
}

func (r rope) areNodesTooDistant(node int) bool {
	return r.getNodesDistance(node) >= 2
}

func (r *rope) moveNodes(node int) {

	if r.nodes[node].x > r.nodes[node-1].x {
		r.nodes[node].x--
	} else if r.nodes[node].x < r.nodes[node-1].x {
		r.nodes[node].x++
	}

	if r.nodes[node].y > r.nodes[node-1].y {
		r.nodes[node].y--
	} else if r.nodes[node].y < r.nodes[node-1].y {
		r.nodes[node].y++
	}

	if node == len(r.nodes)-1 {
		r.tailPath[r.nodes[node].toString()] = true
	}
}

func (r *rope) moveHead(direction int, times int) {
	if direction == 0 {
		for loops := 0; loops < times; loops++ {
			r.nodes[0].y++

			for nodeIndex := 0; nodeIndex < len(r.nodes)-1; nodeIndex++ {
				if r.areNodesTooDistant(nodeIndex) {
					r.moveNodes(nodeIndex + 1)
				}
			}
		}
	} else if direction == 1 {
		for loops := 0; loops < times; loops++ {
			r.nodes[0].x++

			for nodeIndex := 0; nodeIndex < len(r.nodes)-1; nodeIndex++ {
				if r.areNodesTooDistant(nodeIndex) {
					r.moveNodes(nodeIndex + 1)
				}
			}
		}
	} else if direction == 2 {
		for loops := 0; loops < times; loops++ {
			r.nodes[0].y--

			for nodeIndex := 0; nodeIndex < len(r.nodes)-1; nodeIndex++ {
				if r.areNodesTooDistant(nodeIndex) {
					r.moveNodes(nodeIndex + 1)
				}
			}
		}
	} else if direction == 3 {
		for loops := 0; loops < times; loops++ {
			r.nodes[0].x--

			for nodeIndex := 0; nodeIndex < len(r.nodes)-1; nodeIndex++ {
				if r.areNodesTooDistant(nodeIndex) {
					r.moveNodes(nodeIndex + 1)
				}
			}
		}
	}
}

func getDirectionNumber(direction string) int {
	switch direction {
	case "U":
		return 0
	case "R":
		return 1
	case "D":
		return 2
	case "L":
		return 3
	default:
		return -1
	}
}

func getInput() ([][]int, error) {
	inputPath := "./day9/input.txt"

	input, err := os.ReadFile(inputPath)

	if err != nil {
		return nil, err
	}

	var output [][]int

	headSequences := strings.Split(string(input), "\n")
	for _, sequenceString := range headSequences {

		if sequenceString == "" {
			continue
		}

		sequenceArray := strings.Split(sequenceString, " ")
		direction := getDirectionNumber(sequenceArray[0])
		number, err := strconv.Atoi(sequenceArray[1])

		if direction == -1 {
			return nil, fmt.Errorf("could not process direction input")
		}

		if err != nil {
			return nil, err
		}

		output = append(output, []int{direction, number})
	}

	return output, nil
}

func Part1() (int, error) {
	input, err := getInput()

	if err != nil {
		return 0, err
	}

	r := rope{[]node{{0, 0}, {0, 0}}, make(map[string]bool)}

	r.tailPath[r.nodes[len(r.nodes)-1].toString()] = true

	for _, headMovement := range input {
		r.moveHead(headMovement[0], headMovement[1])
	}

	return len(r.tailPath), nil
}

func Part2() (int, error) {
	input, err := getInput()

	if err != nil {
		return 0, err
	}

	r := rope{[]node{{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}}, make(map[string]bool)}

	r.tailPath[r.nodes[len(r.nodes)-1].toString()] = true

	for _, headMovement := range input {
		r.moveHead(headMovement[0], headMovement[1])
	}

	return len(r.tailPath), nil
}
