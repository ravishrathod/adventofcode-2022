package main

import (
	"math"
	"ravishrathod/adventofcode-2022/commons"
)

func main() {
	lines, err := commons.ReadFile("input/day12.txt")
	if err != nil {
		panic(err)
	}
	grid, start, end := parseInput(lines)
	part1(grid, start, end)
	part2(grid, end)
}

func part1(grid [][]rune, start Point, end Point) {
	pathFinder := NewPathFinder(grid, start, end)
	println(pathFinder.ShortestPath())
}

func part2(grid [][]rune, end Point) {
	var lowestPoints []Point
	lowestMarker := []rune("a")[0]
	for y, row := range grid {
		for x, elevation := range row {
			if elevation == lowestMarker {
				point := Point{
					X: x,
					Y: y,
				}
				lowestPoints = append(lowestPoints, point)
			}
		}
	}
	shortestDistance := math.MaxInt
	for _, start := range lowestPoints {
		pathFinder := NewPathFinder(grid, start, end)
		distance := pathFinder.ShortestPath()
		if distance < shortestDistance {
			shortestDistance = distance
		}
	}
	println(shortestDistance)
}

func parseInput(lines []string) ([][]rune, Point, Point) {
	startMarker := []rune("S")[0]
	endMarker := []rune("E")[0]
	var start Point
	var end Point
	grid := make([][]rune, len(lines))
	for y, line := range lines {
		grid[y] = make([]rune, len(line))
		cells := []rune(line)
		for x, cell := range cells {
			value := cell
			if cell == startMarker {
				value = []rune("a")[0]
				start = Point{
					X: x,
					Y: y,
				}
			} else if cell == endMarker {
				value = []rune("z")[0]
				end = Point{
					X: x,
					Y: y,
				}
			}
			grid[y][x] = value
		}
	}
	return grid, start, end
}

type Point struct {
	Y int
	X int
}
