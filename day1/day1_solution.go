package main

import (
	"ravishrathod/adventofcode-2022/commons"
	"sort"
	"strconv"
)

func main() {
	lines, err := commons.ReadFile("input/day1.txt")
	if err != nil {
		panic(err)
	}

	var caloriesByElves []int
	var caloriesForElf = 0
	for _, calories := range lines {
		if calories == "" {
			caloriesByElves = append(caloriesByElves, caloriesForElf)
			caloriesForElf = 0
		} else {
			numericValue, _ := strconv.Atoi(calories)
			caloriesForElf += numericValue
		}
	}
	part1(caloriesByElves)
	part2(caloriesByElves)
}

func part1(caloriesByElves []int) {
	max := -1
	for _, value := range caloriesByElves {
		if value > max {
			max = value
		}
	}
	println(max)
}

func part2(caloriesByElves []int) {
	sort.Ints(caloriesByElves)
	totalCaloriesForTopThree := 0
	for i := 1; i <= 3; i++ {
		totalCaloriesForTopThree += caloriesByElves[len(caloriesByElves)-i]
	}
	println(totalCaloriesForTopThree)
}
