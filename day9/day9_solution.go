package main

import (
	"ravishrathod/adventofcode-2022/commons"
	"strconv"
	"strings"
)

func main() {
	lines, err := commons.ReadFile("input/day9.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	head := &position{}
	tail := &position{}

	tailMoves := make(map[string]bool)
	tailMoves[tail.toString()] = true

	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		steps, _ := strconv.Atoi(parts[1])
		for i := 0; i < steps; i++ {
			move(direction, head)
			moved := moveTail(head, tail)
			if moved {
				tailMoves[tail.toString()] = true
			}
		}
	}
	println(len(tailMoves))
}

func part2(lines []string) {
	knots := make([]*position, 10)
	for i := 0; i < 10; i++ {
		knots[i] = &position{}
	}
	tail := knots[9]
	head := knots[0]
	tailMoves := make(map[string]bool)
	tailMoves[tail.toString()] = true
	for _, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		steps, _ := strconv.Atoi(parts[1])
		for i := 0; i < steps; i++ {
			move(direction, head)
			for knotIdx := 1; knotIdx < 10; knotIdx++ {
				moved := moveTail(knots[knotIdx-1], knots[knotIdx])
				if !moved {
					break
				}
				if knotIdx == 9 {
					tailMoves[tail.toString()] = true
				}
			}
		}
	}
	println(len(tailMoves))
}

func moveTail(head *position, tail *position) bool {
	xDelta := head.X - tail.X
	yDelta := head.Y - tail.Y
	xMod := mod(xDelta)
	yMod := mod(yDelta)
	if xMod <= 1 && yMod <= 1 {
		return false
	}
	if xMod > 1 {
		tail.X = tail.X + calculateStep(xDelta)

		if yMod > 0 {
			tail.Y = tail.Y + calculateStep(yDelta)
		}
	} else if yMod > 1 {
		tail.Y = tail.Y + calculateStep(yDelta)

		if xMod > 0 {
			tail.X = tail.X + calculateStep(xDelta)
		}
	}
	return true
}

func calculateStep(delta int) int {
	step := 1
	if delta < 0 {
		step = -1
	}
	return step
}
func move(direction string, position *position) {
	if direction == "R" {
		position.X = position.X + 1
	} else if direction == "L" {
		position.X = position.X - 1
	} else if direction == "U" {
		position.Y = position.Y + 1
	} else if direction == "D" {
		position.Y = position.Y - 1
	}
}

type position struct {
	X int
	Y int
}

func (p *position) toString() string {
	return strconv.Itoa(p.X) + "," + strconv.Itoa(p.Y)
}

func mod(input int) int {
	if input >= 0 {
		return input
	}
	return input * -1
}
