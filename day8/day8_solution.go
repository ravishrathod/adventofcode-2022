package main

import (
	"ravishrathod/adventofcode-2022/commons"
)

func main() {

	lines, err := commons.ReadFile("input/day8.txt")
	if err != nil {
		panic(err)
	}
	grid := parseInput(lines)
	part1(grid)
	part2(grid)
}

func part1(grid [][]int) {
	visibleTrees := ((len(grid) - 2) + len(grid[0])) * 2
	gridHeight := len(grid)
	for y := 1; y < gridHeight-1; y++ {
		gridWidth := len(grid[y])
		for x := 1; x < gridWidth-1; x++ {
			if isVisible(y, x, grid) {
				visibleTrees++
			}
		}
	}
	println(visibleTrees)
}

func part2(grid [][]int) {
	gridHeight := len(grid)
	bestScore := -1
	for y := 1; y < gridHeight-1; y++ {
		gridWidth := len(grid[y])
		for x := 1; x < gridWidth-1; x++ {
			score := calculateVisibilityScore(y, x, grid)
			if score > bestScore {
				bestScore = score
			}
		}
	}
	println(bestScore)
}

func calculateVisibilityScore(y int, x int, grid [][]int) int {
	left := 0
	right := 0
	top := 0
	bottom := 0
	height := grid[y][x]
	for i := x - 1; i >= 0; i-- {
		left++
		if grid[y][i] >= height {
			break
		}
	}
	for i := x + 1; i < len(grid[y]); i++ {
		right++
		if grid[y][i] >= height {
			break
		}
	}
	for i := y - 1; i >= 0; i-- {
		top++
		if grid[i][x] >= height {
			break
		}
	}
	for i := y + 1; i < len(grid); i++ {
		bottom++
		if grid[i][x] >= height {
			break
		}
	}
	return left * right * top * bottom
}

func isVisible(y int, x int, grid [][]int) bool {
	height := grid[y][x]
	visible := true
	for i := x - 1; i >= 0; i-- {
		if grid[y][i] >= height {
			visible = false
			break
		}
	}

	if visible {
		return visible
	}

	visible = true
	for i := x + 1; i < len(grid[y]); i++ {
		if grid[y][i] >= height {
			visible = false
			break
		}
	}

	if visible {
		return visible
	}

	visible = true
	for i := y - 1; i >= 0; i-- {
		if grid[i][x] >= height {
			visible = false
			break
		}
	}

	if visible {
		return visible
	}

	visible = true
	for i := y + 1; i < len(grid); i++ {
		if grid[i][x] >= height {
			visible = false
			break
		}
	}

	return visible
}

func parseInput(lines []string) [][]int {
	grid := make([][]int, len(lines))
	for i, line := range lines {
		row := commons.LineToIntArrayNoSeparator(line)
		grid[i] = row
	}
	return grid
}
