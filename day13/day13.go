/*
	Day 13: Distress Signal
	Link: https://adventofcode.com/2022/day/13
*/

package day13

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type packet struct {
	value []any
}

type packetPair struct {
	left  packet
	right packet
}

func getInput() ([]packetPair, error) {
	inputPath := "./inputs/2022/day13.txt"

	content, err := os.ReadFile(inputPath)

	if err != nil {
		return nil, err
	}

	allPairsString := strings.Split(string(content), "\n\n")
	var output []packetPair

	for _, pairString := range allPairsString {
		pairPacketString := strings.Split(pairString, "\n")
		leftPacketString := pairPacketString[0]
		rightPacketString := pairPacketString[1]

		var leftPacket []any
		var rightPacket []any

		err := json.Unmarshal([]byte(leftPacketString), &leftPacket)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal([]byte(rightPacketString), &rightPacket)

		if err != nil {
			return nil, err
		}

		output = append(output, packetPair{packet{leftPacket}, packet{rightPacket}})
	}

	return output, nil
}

func bothLists(listLeft []any, listRight []any) int {
	var shortestList []any
	var short string

	if len(listLeft) < len(listRight) {
		shortestList = listLeft
		short = "left"
	} else if len(listLeft) > len(listRight) {
		shortestList = listRight
		short = "right"
	} else {
		shortestList = listLeft
		short = "none"
	}

	for i := 0; i < len(shortestList); i++ {
		out := compareValues(listLeft[i], listRight[i])

		if out != 0 {
			return out
		}
	}

	if short != "none" {
		if short == "left" {
			return 1
		} else {
			return -1
		}
	} else {
		return 0
	}
}

func bothInts(intLeft int, intRight int) int {
	if intLeft < intRight {
		return 1
	} else if intLeft > intRight {
		return -1
	} else {
		return 0
	}
}

func compareValues(valueLeft any, valueRight any) int {
	if reflect.TypeOf(valueLeft).String() == "[]interface {}" && reflect.TypeOf(valueRight).String() == "[]interface {}" {
		return bothLists(valueLeft.([]any), valueRight.([]any))
	} else if reflect.TypeOf(valueLeft).String() == "float64" && reflect.TypeOf(valueRight).String() == "float64" {
		return bothInts(int(valueLeft.(float64)), int(valueRight.(float64)))
	} else {
		var number any
		if reflect.TypeOf(valueLeft).String() == "float64" {
			number = []any{valueLeft.(float64)}
			return compareValues(number, valueRight)
		} else {
			number = []any{valueRight.(float64)}
			return compareValues(valueLeft, number)
		}
	}
}

func Part1() (int, error) {
	input, err := getInput()

	if err != nil {
		return 0, err
	}

	indexSum := 0

	for index, pair := range input {
		out := compareValues(pair.left.value, pair.right.value)

		if out == 1 {
			indexSum += index + 1
		}
	}

	return indexSum, nil
}

func Part2() (int, error) {
	input, err := getInput()

	if err != nil {
		return 0, err
	}

	var orderedPackets []packet

	for _, pair := range input {
		orderedPackets = append(orderedPackets, pair.left, pair.right)
	}

	dividerLeft := packet{[]any{[]any{2.0}}}
	dividerRight := packet{[]any{[]any{6.0}}}

	orderedPackets = append(orderedPackets, dividerLeft, dividerRight)

	var ordered bool = false

	for !ordered {
		ordered = true

		for i := 0; i < len(orderedPackets)-1; i++ {
			out := compareValues(orderedPackets[i].value, orderedPackets[i+1].value)

			if out == -1 {
				temp := orderedPackets[i]
				orderedPackets[i] = orderedPackets[i+1]
				orderedPackets[i+1] = temp
				ordered = false
			}
		}
	}

	var output []int

	for index, ordPacket := range orderedPackets {
		if fmt.Sprint(ordPacket) == fmt.Sprint(dividerLeft) || fmt.Sprint(ordPacket) == fmt.Sprint(dividerRight) {
			output = append(output, index+1)
		}
	}

	return output[0] * output[1], nil
}
