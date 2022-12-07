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

func part1(stacks []*commons.Stack[string], instructions []Instruction) {
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

func part2(stacks []*commons.Stack[string], instructions []Instruction) {
	for _, instruction := range instructions {
		source := stacks[instruction.From-1]
		destination := stacks[instruction.To-1]
		tempStack := &commons.Stack[string]{}
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

func parseInput(lines []string) ([]*commons.Stack[string], []Instruction) {
	var instructions []Instruction
	var stacks []*commons.Stack[string]

	//create stacks
	for _, line := range lines {
		if strings.Contains(line, "[") {
			continue
		}
		stackIds := strings.Split(strings.TrimSpace(line), "   ")
		stacks = make([]*commons.Stack[string], len(stackIds))
		for idx := range stackIds {
			stacks[idx] = &commons.Stack[string]{}
		}
		break
	}

	stackLines := &commons.Stack[string]{}
	for _, line := range lines {
		//parse instructions
		if strings.HasPrefix(line, "move") {
			parts := strings.Fields(line)
			quantity, _ := strconv.Atoi(parts[1])
			from, _ := strconv.Atoi(parts[3])
			to, _ := strconv.Atoi(parts[5])
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

func updateStacks(line string, stacks []*commons.Stack[string]) {
	characters := []rune(line)
	for i, char := range characters {
		if char == '[' {
			crate := string(characters[i+1])
			stackIndex := i / 4
			stackAtPosition := stacks[stackIndex]
			stackAtPosition.Push(crate)
		}
	}
}

type Instruction struct {
	From     int
	To       int
	Quantity int
}
