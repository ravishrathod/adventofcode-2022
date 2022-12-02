package main

import (
	"ravishrathod/adventofcode-2022/commons"
	"strings"
)

const ROCK = "ROCK"
const PAPER = "PAPER"
const SCISSORS = "SCISSORS"

func main() {
	lines, err := commons.ReadFile("input/day2.txt")
	if err != nil {
		panic(err)
	}

	var opponentMoves = make(map[string]string)
	opponentMoves["A"] = ROCK
	opponentMoves["B"] = PAPER
	opponentMoves["C"] = SCISSORS

	var movePoints = make(map[string]int)
	movePoints[ROCK] = 1
	movePoints[PAPER] = 2
	movePoints[SCISSORS] = 3

	part1(lines, opponentMoves, movePoints)
	part2(lines, opponentMoves, movePoints)
}

func part1(lines []string, opponentMoves map[string]string, movePoints map[string]int) {
	var myMoves = make(map[string]string)
	myMoves["X"] = ROCK
	myMoves["Y"] = PAPER
	myMoves["Z"] = SCISSORS
	totalPoints := 0
	for _, line := range lines {
		moves := strings.Split(line, " ")
		points := calculateScore(opponentMoves[moves[0]], myMoves[moves[1]], movePoints)
		totalPoints += points
	}

	println(totalPoints)
}

func part2(lines []string, opponentMoves map[string]string, movePoints map[string]int) {
	totalPoints := 0
	for _, line := range lines {
		input := strings.Split(line, " ")
		opponentMove := opponentMoves[input[0]]
		desiredOutcome := input[1]
		myMove := calculateMove(opponentMove, desiredOutcome)
		points := calculateScore(opponentMove, myMove, movePoints)
		totalPoints += points
	}
	println(totalPoints)
}

func calculateMove(opponentMove string, desiredOutcome string) string {
	if desiredOutcome == "Y" {
		return opponentMove
	}
	if opponentMove == ROCK {
		if desiredOutcome == "Z" {
			return PAPER
		} else {
			return SCISSORS
		}
	}
	if opponentMove == PAPER {
		if desiredOutcome == "Z" {
			return SCISSORS
		} else {
			return ROCK
		}
	}
	if opponentMove == SCISSORS {
		if desiredOutcome == "Z" {
			return ROCK
		} else {
			return PAPER
		}
	}
	return ""
}

func calculateScore(opponentMove string, myMove string, movePoints map[string]int) int  {
	myMovePoint := movePoints[myMove]
	if opponentMove == myMove {
		return myMovePoint + 3
	}
	won := didIWin(opponentMove, myMove)
	if won {
		return myMovePoint + 6
	}
	return myMovePoint
}

func didIWin(opponentMove string, myMove string) bool {
	if opponentMove == ROCK {
		return myMove == PAPER
	}
	if opponentMove == PAPER {
		return myMove == SCISSORS
	}
	if opponentMove == SCISSORS {
		return myMove == ROCK
	}
	return true
}