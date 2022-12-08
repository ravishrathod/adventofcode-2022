package main

import (
	"ravishrathod/adventofcode-2022/commons"
	"strings"
)

type Move string

const (
	ROCK     Move = "ROCK"
	PAPER    Move = "PAPER"
	SCISSORS Move = "SCISSORS"
)

func main() {
	lines, err := commons.ReadFile("input/day2.txt")
	if err != nil {
		panic(err)
	}

	var opponentMoves = make(map[string]Move)
	opponentMoves["A"] = ROCK
	opponentMoves["B"] = PAPER
	opponentMoves["C"] = SCISSORS

	var movePoints = make(map[Move]int)
	movePoints[ROCK] = 1
	movePoints[PAPER] = 2
	movePoints[SCISSORS] = 3

	part1(lines, opponentMoves, movePoints)
	part2(lines, opponentMoves, movePoints)
}

func part1(lines []string, opponentMoves map[string]Move, movePoints map[Move]int) {
	var myMoves = make(map[string]Move)
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

func part2(lines []string, opponentMoves map[string]Move, movePoints map[Move]int) {
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

func calculateMove(opponentMove Move, desiredOutcome string) Move {
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

func calculateScore(opponentMove Move, myMove Move, movePoints map[Move]int) int {
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

func didIWin(opponentMove Move, myMove Move) bool {
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
