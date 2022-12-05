package main

import (
	"ravishrathod/adventofcode-2022/commons"
	"strconv"
	"strings"
)

func main() {
	lines, err := commons.ReadFile("input/day5.txt")
	if err != nil {
		panic(err)
	}
	stacks, instructions := parseInput(lines)
	part1(stacks, instructions)
	stacks, _ = parseInput(lines)
	println("\n------")
	part2(stacks, instructions)
}

func part1(stacks []*commons.Stack, instructions []Instruction) {
	for _, instruction := range instructions {
		source := stacks[instruction.From-1]
		destination := stacks[instruction.To-1]
		for i := 0; i < instruction.Quantity; i++ {
			crate, err := source.Pop()
			if err != nil {
				panic(err)
			}
			destination.Push(crate)
		}
	}
	for _, stack := range stacks {
		topCrate, _ := stack.Peek()
		print(topCrate)
	}
}

func part2(stacks []*commons.Stack, instructions []Instruction) {
	for _, instruction := range instructions {
		source := stacks[instruction.From-1]
		destination := stacks[instruction.To-1]
		tempStack := &commons.Stack{}
		for i := 0; i < instruction.Quantity; i++ {
			crate, err := source.Pop()
			if err != nil {
				panic(err)
			}
			tempStack.Push(crate)
		}
		for !tempStack.IsEmpty() {
			crate, _ := tempStack.Pop()
			destination.Push(crate)
		}
	}
	for _, stack := range stacks {
		topCrate, _ := stack.Peek()
		print(topCrate)
	}
}

func parseInput(lines []string) ([]*commons.Stack, []Instruction) {
	var instructions []Instruction
	var stacks []*commons.Stack

	//create stacks
	for _, line := range lines {
		if strings.Contains(line, "[") {
			continue
		}
		stackIds := strings.Split(strings.TrimSpace(line), "   ")
		stacks = make([]*commons.Stack, len(stackIds))
		for idx, _ := range stackIds {
			stacks[idx] = &commons.Stack{}
		}
		break
	}

	stackLines := &commons.Stack{}
	for _, line := range lines {
		//parse instructions
		if strings.HasPrefix(line, "move") {
			updated := strings.Replace(line, "move ", "", 1)
			updated = strings.Replace(updated, "from ", "", 1)
			updated = strings.Replace(updated, "to ", "", 1)
			values := strings.Split(updated, " ")
			quantity, _ := strconv.Atoi(values[0])
			from, _ := strconv.Atoi(values[1])
			to, _ := strconv.Atoi(values[2])
			instructions = append(instructions, Instruction{From: from, To: to, Quantity: quantity})
		} else if strings.Contains(line, "[") {
			stackLines.Push(line)
		}
	}

	// process stack values
	var reversedStackLines []string
	for !stackLines.IsEmpty() {
		value, _ := stackLines.Pop()
		reversedStackLines = append(reversedStackLines, value)
	}
	for _, line := range reversedStackLines {
		updateStacks(line, stacks)
	}

	return stacks, instructions
}

func updateStacks(line string, stacks []*commons.Stack) {
	characters := []rune(line)
	for i := 0; i < len(characters); i++ {
		if characters[i] == '[' {
			crate := string(characters[i+1])
			stackIndex := findStackIndexForPosition(i)
			stackAtPosition := stacks[stackIndex]
			stackAtPosition.Push(crate)
		}
	}
}

func findStackIndexForPosition(position int) int {
	if position == 0 {
		return 0
	}
	return position / 4
}

type Instruction struct {
	From     int
	To       int
	Quantity int
}
