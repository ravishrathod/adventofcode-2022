package main

import (
	"ravishrathod/adventofcode-2022/commons"
	"strconv"
	"strings"
)

func main() {
	lines, err := commons.ReadFile("input/day10.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	x := 1
	cycle := 1
	totalStrength := 0
	for _, line := range lines {
		cycle++
		if isCycleOfInterest(cycle) {
			totalStrength += cycle * x
		}
		if strings.HasPrefix(line, "addx") {
			parts := strings.Split(line, " ")
			numberToAdd, _ := strconv.Atoi(parts[1])
			x += numberToAdd
			cycle++
			if isCycleOfInterest(cycle) {
				totalStrength += cycle * x
			}
		}
	}
	println(totalStrength)
}

func part2(lines []string) {
	crt := make([][]string, 6)
	for i := range crt {
		crt[i] = make([]string, 40)
	}
	cycle := 0
	x := 1
	for _, line := range lines {
		if line == "noop" {
			cycle++
			drawPixel(cycle, x, crt)
		} else {
			//addx
			cycle++
			drawPixel(cycle, x, crt)
			cycle++
			drawPixel(cycle, x, crt)
			parts := strings.Split(line, " ")
			numberToAdd, _ := strconv.Atoi(parts[1])
			x += numberToAdd
		}
	}
	for _, row := range crt {
		for _, pixel := range row {
			print(" " + pixel)
		}
		println("")
	}
}

func drawPixel(cycle int, x int, crt [][]string) {
	rowNumber := (cycle - 1) / 40
	row := crt[rowNumber]
	pixelPos := (cycle - rowNumber*40) - 1
	if pixelPos >= (x-1) && pixelPos < (x+2) {
		row[pixelPos] = "#"
	} else {
		row[pixelPos] = "."
	}
}

func isCycleOfInterest(cycle int) bool {
	if cycle == 20 {
		return true
	}
	if cycle < 20 {
		return false
	}
	return ((cycle - 20) % 40) == 0
}
