/*
	Day 2: Rock Paper Scissors
	Link: https://adventofcode.com/2022/day/2
*/

package day02

import (
	"os"
	"strings"
)

var possibleResult = map[string][]string{
	"rock":     {"paper", "scissors"},
	"paper":    {"scissors", "rock"},
	"scissors": {"rock", "paper"},
}

var resultPoints = map[string]int{
	"win":  6,
	"draw": 3,
	"lose": 0,
}

var inputPoints = map[string]int{
	"rock":     1,
	"paper":    2,
	"scissors": 3,
}

var opponentInputs = map[string]string{
	"A": "rock",
	"B": "paper",
	"C": "scissors",
}

var myInputsPart1 = map[string]string{
	"X": "rock",
	"Y": "paper",
	"Z": "scissors",
}

var myInputsPart2 = map[string]string{
	"X": "lose",
	"Y": "draw",
	"Z": "win",
}

type gameResult struct {
	name   string
	points int
}

type gameRound struct {
	opponentInput string
	myInput       string
}

func (round gameRound) getResult() gameResult {
	var result string

	if round.opponentInput == round.myInput {
		result = "draw"
	} else if (round.opponentInput == "rock" && round.myInput == "paper") || (round.opponentInput == "paper" && round.myInput == "scissors") || (round.opponentInput == "scissors" && round.myInput == "rock") {
		result = "win"
	} else {
		result = "lose"
	}

	return gameResult{result, resultPoints[result]}
}

func (round gameRound) getScore() int {
	scoreResult := round.getResult().points
	scoreInput := inputPoints[round.myInput]

	return scoreInput + scoreResult
}

type gameRoundPart1 struct {
	gameRound
}

type gameRoundPart2 struct {
	gameRound
}

func (round *gameRoundPart1) convertInputs() {
	round.opponentInput = opponentInputs[round.opponentInput]
	round.myInput = myInputsPart1[round.myInput]
}

func (round *gameRoundPart2) convertInputs() {
	round.opponentInput = opponentInputs[round.opponentInput]

	if myInputsPart2[round.myInput] == "win" {
		round.myInput = possibleResult[round.opponentInput][0]
	} else if myInputsPart2[round.myInput] == "lose" {
		round.myInput = possibleResult[round.opponentInput][1]
	} else {
		round.myInput = round.opponentInput
	}
}

func newRoundPart1(opponentInput string, myInput string) gameRoundPart1 {
	round := gameRoundPart1{gameRound{opponentInput, myInput}}
	round.convertInputs()
	return round
}

func newRoundPart2(opponentInput string, myInput string) gameRoundPart2 {
	round := gameRoundPart2{gameRound{opponentInput, myInput}}
	round.convertInputs()
	return round
}

func getInput() ([][]string, error) {
	inputPath := "./inputs/2022/day02.txt"

	content, err := os.ReadFile(inputPath)

	if err != nil {
		return nil, err
	}

	var output [][]string

	roundsArray := strings.Split(string(content), "\n")

	for _, element := range roundsArray {

		roundInputsArray := strings.Split(element, " ")

		if len(roundInputsArray) < 2 {
			continue
		}

		output = append(output, roundInputsArray)
	}

	return output, nil
}

func Part1() (int, error) {
	input, err := getInput()

	if err != nil {
		return 0, err
	}

	myScore := 0

	for _, round := range input {
		myScore += newRoundPart1(round[0], round[1]).getScore()
	}

	return myScore, nil
}

func Part2() (int, error) {
	input, err := getInput()

	if err != nil {
		return 0, err
	}

	myScore := 0

	for _, round := range input {
		myScore += newRoundPart2(round[0], round[1]).getScore()
	}

	return myScore, nil
}
