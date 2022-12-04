package main

import (
	"ravishrathod/adventofcode-2022/commons"
	"strconv"
	"strings"
)

func main() {
	lines, err := commons.ReadFile("input/day4.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	count := 0
	for _, line := range lines {
		assignments := strings.Split(line, ",")
		if oneContainsOther(assignments[0], assignments[1]) {
			count++
		}
	}
	println(count)
}

func part2(lines []string) {
	count := 0
	for _, line := range lines {
		assignments := strings.Split(line, ",")
		if areOverlapping(assignments[0], assignments[1]) {
			count++
		}
	}
	println(count)
}

func oneContainsOther(firstAssignment string, secondAssignment string) bool {
	firstBounds := parseAssignmentBounds(firstAssignment)
	secondBounds := parseAssignmentBounds(secondAssignment)
	return (firstBounds[0] >= secondBounds[0] && firstBounds[1] <= secondBounds[1]) ||
		(secondBounds[0] >= firstBounds[0] && secondBounds[1] <= firstBounds[1])
}

func areOverlapping(firstAssignment string, secondAssignment string) bool {
	firstBounds := parseAssignmentBounds(firstAssignment)
	secondBounds := parseAssignmentBounds(secondAssignment)
	return (firstBounds[0] >= secondBounds[0] && firstBounds[0] <= secondBounds[1]) ||
		(secondBounds[0] >= firstBounds[0] && secondBounds[0] <= firstBounds[1])
}

func parseAssignmentBounds(assignment string) []int {
	parts := strings.Split(assignment, "-")
	lowerBound, _ := strconv.Atoi(parts[0])
	upperBound, _ := strconv.Atoi(parts[1])
	return []int{lowerBound, upperBound}
}
