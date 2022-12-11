package main

import (
	"ravishrathod/adventofcode-2022/commons"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := commons.ReadFile("input/day11_sample.txt")
	if err != nil {
		panic(err)
	}
	monkeyMap := make(map[string]*Monkey)
	var monkeys []*Monkey

	for i := 0; i < len(lines); {
		chunk := lines[i : i+6]
		i += 7
		monkey := parseMonkey(chunk)
		monkeyMap[monkey.Id] = monkey
		monkeys = append(monkeys, monkey)
	}
	part1(monkeyMap, monkeys)
	part2(monkeyMap, monkeys)
}

func part1(monkeyMap map[string]*Monkey, monkeys []*Monkey) {
	monkeyBusiness := processRounds(20, 3, monkeyMap, monkeys)
	println(monkeyBusiness)
}

func part2(monkeyMap map[string]*Monkey, monkeys []*Monkey) {
	monkeyBusiness := processRounds(10000, 1, monkeyMap, monkeys)
	println(monkeyBusiness)
}
func processRounds(rounds int, reliefFactor uint64, monkeyMap map[string]*Monkey, monkeys []*Monkey) int {
	itemsInspectedByMonkey := make(map[string]int)
	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			for _, itemWorryLevel := range monkey.ItemsWorryLevel {
				newWorryLevel := monkey.GetNewWorryLevel(itemWorryLevel)
				newWorryLevel = newWorryLevel / reliefFactor
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
func parseMonkey(lines []string) *Monkey {
	part := strings.Split(lines[0], " ")[1]
	monkeyId := strings.Replace(part, ":", "", 1)
	//worryLevels := commons.LinetoIntArray(strings.Replace(lines[1], "Starting items: ", "", 1))
	var worryLevels []uint64
	levelsString := strings.Replace(lines[1], "Starting items: ", "", 1)
	for _, str := range strings.Split(levelsString, ",") {
		str = strings.TrimSpace(str)
		val, _ := strconv.ParseUint(str, 0, 64)
		worryLevels = append(worryLevels, val)
	}

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
		Condition: uint64(testCondition),
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
	ItemsWorryLevel []uint64
	Operation       *Operation
	Test            *ThrowTest
}

func (m *Monkey) GetNewWorryLevel(old uint64) uint64 {
	return m.Operation.apply(old)
}

func (m *Monkey) GetMonkeyToThrowAt(worryLevel uint64) string {
	return m.Test.Evaluate(worryLevel)
}

func (m *Monkey) AddItem(worryLevel uint64) {
	m.ItemsWorryLevel = append(m.ItemsWorryLevel, worryLevel)
}

func (m *Monkey) ThrowItem() {
	if len(m.ItemsWorryLevel) == 0 {
		panic("No items")
	}
	m.ItemsWorryLevel = m.ItemsWorryLevel[1:]
}

type ThrowTest struct {
	Condition uint64
	OnPass    string
	OnFail    string
}

func (t *ThrowTest) Evaluate(worryLevel uint64) string {
	if worryLevel%t.Condition == 0 {
		return t.OnPass
	}
	return t.OnFail
}

type Operation struct {
	Operator string
	Operand  string
}

func (o *Operation) apply(old uint64) uint64 {
	currentOperand := uint64(0)
	if o.Operand == "old" {
		currentOperand = old
	} else {
		strconv.ParseUint(o.Operand, 0, 64)
		currentOperand, _ = strconv.ParseUint(o.Operand, 0, 64)
	}
	if o.Operator == "*" {
		return old * currentOperand
	} else if o.Operator == "+" {
		return old + currentOperand
	}
	panic("invalid operator")
}
