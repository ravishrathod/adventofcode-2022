package main

import (
	"math"
	"ravishrathod/adventofcode-2022/commons"
	"strconv"
	"strings"
)

func main() {
	lines, err := commons.ReadFile("input/day21.txt")
	if err != nil {
		panic(err)
	}
	solvedMonkeys := make(map[string]int)
	monkeys := make(map[string]Monkey)

	for _, line := range lines {
		parts := strings.Split(line, ":")
		name := parts[0]
		task := strings.TrimSpace(parts[1])
		value, err := strconv.Atoi(task)
		if err != nil {
			monkey := Monkey{
				Name:  name,
				Input: task,
			}
			monkeys[name] = monkey
			solvedMonkeys[name] = math.MinInt
		} else {
			solvedMonkeys[name] = value
		}
	}
	stack := &commons.Stack[string]{}
	solveForMonkey("root", stack, solvedMonkeys, monkeys)
	println(solvedMonkeys["root"])
}

func solveForMonkey(name string, stack *commons.Stack[string], solvedMonkeys map[string]int, monkeys map[string]Monkey) {
	monkey := monkeys[name]
	blocker1, blocker2 := monkey.DependsOn()
	if solvedMonkeys[blocker1] > math.MinInt && solvedMonkeys[blocker2] > math.MinInt {
		solution := monkey.Solve(solvedMonkeys)
		solvedMonkeys[name] = solution
	} else {
		stack.Push(name)
		if solvedMonkeys[blocker1] == math.MinInt {
			stack.Push(blocker1)
		}
		if solvedMonkeys[blocker2] == math.MinInt {
			stack.Push(blocker2)
		}
	}
	if !stack.IsEmpty() {
		nextMonkey, _ := stack.Pop()
		solveForMonkey(nextMonkey, stack, solvedMonkeys, monkeys)
	}
}

type Monkey struct {
	Name  string
	Input string
}

func (m *Monkey) DependsOn() (string, string) {
	parts := strings.Split(m.Input, " ")
	return parts[0], parts[2]
}
func (m *Monkey) Solve(solvedMonkeys map[string]int) int {
	parts := strings.Split(m.Input, " ")
	operand := parts[1]
	left := parts[0]
	right := parts[2]
	leftOperator := solvedMonkeys[left]
	rightOperator := solvedMonkeys[right]

	if operand == "+" {
		return leftOperator + rightOperator
	}
	if operand == "-" {
		return leftOperator - rightOperator
	}
	if operand == "*" {
		return leftOperator * rightOperator
	}
	if operand == "/" {
		return leftOperator / rightOperator
	}
	panic("invalid operand" + operand)
}
