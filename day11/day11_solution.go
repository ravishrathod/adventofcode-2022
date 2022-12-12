package main

import (
	"ravishrathod/adventofcode-2022/commons"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := commons.ReadFile("input/day11.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	monkeyMap, monkeys := parseMonkeys(lines)
	monkeyBusiness := processRounds(20, func(input int) int { return input / 3 }, monkeyMap, monkeys)
	println(monkeyBusiness)
}

func part2(lines []string) {
	monkeyMap, monkeys := parseMonkeys(lines)
	reliefFactor := 1
	for _, monkey := range monkeys {
		reliefFactor *= monkey.Test.Condition
	}
	monkeyBusiness := processRounds(10000, func(input int) int { return input % reliefFactor }, monkeyMap, monkeys)
	println(monkeyBusiness)
}
func processRounds(rounds int, reliefFactor func(input int) int, monkeyMap map[string]*Monkey, monkeys []*Monkey) int {
	itemsInspectedByMonkey := make(map[string]int)
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			for _, itemWorryLevel := range monkey.ItemsWorryLevel {
				newWorryLevel := monkey.GetNewWorryLevel(itemWorryLevel)
				newWorryLevel = reliefFactor(newWorryLevel)
				targetMonkeyId := monkey.GetMonkeyToThrowAt(newWorryLevel)
				monkeyMap[targetMonkeyId].AddItem(newWorryLevel)
				monkey.ThrowItem()
				itemsInspectedByMonkey[monkey.Id] += 1
			}
		}
	}
	println()
	var throws []int
	for _, throw := range itemsInspectedByMonkey {
		throws = append(throws, throw)
	}
	sort.Ints(throws)
	monkeyBusiness := throws[len(throws)-1] * throws[len(throws)-2]
	return monkeyBusiness
}

func parseMonkeys(lines []string) (map[string]*Monkey, []*Monkey) {
	monkeyMap := make(map[string]*Monkey)
	var monkeys []*Monkey

	for i := 0; i < len(lines); {
		chunk := lines[i : i+6]
		i += 7
		monkey := parseMonkey(chunk)
		monkeyMap[monkey.Id] = monkey
		monkeys = append(monkeys, monkey)
	}
	return monkeyMap, monkeys
}
func parseMonkey(lines []string) *Monkey {
	part := strings.Split(lines[0], " ")[1]
	monkeyId := strings.Replace(part, ":", "", 1)
	worryLevels := commons.LinetoIntArray(strings.Replace(lines[1], "Starting items: ", "", 1))
	operationString := strings.Replace(lines[2], "  Operation: new = old ", "", 1)
	operationString = strings.TrimSpace(operationString)
	parts := strings.Split(operationString, " ")
	operation := &Operation{
		Operator: parts[0],
		Operand:  parts[1],
	}
	testConditionString := strings.Replace(lines[3], "  Test: divisible by ", "", 1)
	testCondition, _ := strconv.Atoi(testConditionString)
	onPass := strings.Replace(lines[4], "    If true: throw to monkey ", "", 1)
	onFail := strings.Replace(lines[5], "    If false: throw to monkey ", "", 1)
	throwTest := &ThrowTest{
		Condition: testCondition,
		OnPass:    onPass,
		OnFail:    onFail,
	}
	return &Monkey{
		Id:              monkeyId,
		ItemsWorryLevel: worryLevels,
		Operation:       operation,
		Test:            throwTest,
	}
}

type Monkey struct {
	Id              string
	ItemsWorryLevel []int
	Operation       *Operation
	Test            *ThrowTest
}

func (m *Monkey) GetNewWorryLevel(old int) int {
	return m.Operation.apply(old)
}

func (m *Monkey) GetMonkeyToThrowAt(worryLevel int) string {
	return m.Test.Evaluate(worryLevel)
}

func (m *Monkey) AddItem(worryLevel int) {
	m.ItemsWorryLevel = append(m.ItemsWorryLevel, worryLevel)
}

func (m *Monkey) ThrowItem() {
	if len(m.ItemsWorryLevel) == 0 {
		panic("No items")
	}
	m.ItemsWorryLevel = m.ItemsWorryLevel[1:]
}

type ThrowTest struct {
	Condition int
	OnPass    string
	OnFail    string
}

func (t *ThrowTest) Evaluate(worryLevel int) string {
	if worryLevel%t.Condition == 0 {
		return t.OnPass
	}
	return t.OnFail
}

type Operation struct {
	Operator string
	Operand  string
}

func (o *Operation) apply(old int) int {
	currentOperand := 0
	if o.Operand == "old" {
		currentOperand = old
	} else {
		currentOperand, _ = strconv.Atoi(o.Operand)
	}
	if o.Operator == "*" {
		return old * currentOperand
	} else if o.Operator == "+" {
		return old + currentOperand
	}
	panic("invalid operator")
}
