package main

import (
	"fmt"
	"ravishrathod/adventofcode-2022/commons"
	"strconv"
	"strings"
)

func main() {

	lines, err := commons.ReadFile("input/day22.txt")
	if err != nil {
		panic(err)
	}
	grid, paths := parseInput(lines)
	row := grid[0]
	startX := 0
	for x := 0; x < len(row); x++ {
		if row[x] == "." {
			startX = x
			break
		}
	}
	position := &Position{
		X:      startX,
		Y:      0,
		Facing: right,
	}
	for idx, path := range paths {
		println("Processing path ", idx)
		for i := 0; i < path.Steps; i++ {
			x, y, _ := nextStepCoords(grid, position)
			if grid[y][x] == "." {
				position.X = x
				position.Y = y
			} else if grid[y][x] == "#" {
				//position.Turn(path.Direction)
				break
			}
		}
		position.Turn(path.Direction)
	}
	fmt.Printf("\n%v\n", position)
	println(calculatePassword(position))
}

func calculatePassword(position *Position) int {
	direction := -1
	switch position.Facing {
	case right:
		direction = 0
		break
	case down:
		direction = 1
		break
	case left:
		direction = 2
		break
	case up:
		direction = 3
		break
	}
	return 1000*(position.Y+1) + 4*(position.X+1) + direction
}

func getFirstOpenTileForRow(row []string) int {
	for i := 0; i < len(row); i++ {
		if row[i] != " " {
			if row[i] == "." {
				return i
			} else {
				return -1
			}
		}
	}
	return -1
}

func getLastOpenTileForRow(row []string) int {
	for i := len(row) - 1; i >= 0; i-- {
		if row[i] != " " {
			if row[i] == "." {
				return i
			} else {
				return -1
			}
		}
	}
	return -1
}

func getFirstOpenTileForColumn(x int, grid [][]string) int {
	for y, row := range grid {
		if row[x] != " " {
			if row[x] == "." {
				return y
			} else {
				return -1
			}
		}
	}
	return -1
}

func getLastOpenTileForColumn(x int, grid [][]string) int {
	for y := len(grid) - 1; y >= 0; y-- {
		row := grid[y]
		if row[x] != " " {
			if row[x] == "." {
				return y
			} else {
				return -1
			}
		}
	}
	return -1
}

func nextStepCoords(grid [][]string, position *Position) (int, int, bool) {
	x, y := position.NextStepCoords()
	if position.Facing == right || position.Facing == left {
		if x > len(grid[y])-1 || (x >= 0 && grid[y][x] == " " && position.Facing == right) {
			wrapAroundX := getFirstOpenTileForRow(grid[y])
			if wrapAroundX != -1 {
				return wrapAroundX, y, true
			} else {
				return x - 1, y, false
			}
		} else if x < 0 || (grid[y][x] == " " && position.Facing == left) {
			wrapAroundX := getLastOpenTileForRow(grid[y])
			if wrapAroundX != -1 {
				return wrapAroundX, y, true
			} else {
				return x + 1, y, false
			}
		}
	} else if position.Facing == up || position.Facing == down {
		if y > len(grid)-1 || (y >= 0 && grid[y][x] == " " && position.Facing == down) {
			wrapAroundY := getFirstOpenTileForColumn(x, grid)
			if wrapAroundY != -1 {
				return x, wrapAroundY, true
			} else {
				return x, y - 1, false
			}
		} else if y < 0 || (grid[y][x] == " " && position.Facing == up) {
			wrapAroundY := getLastOpenTileForColumn(x, grid)
			if wrapAroundY != -1 {
				return x, wrapAroundY, true
			} else {
				return x, y + 1, false
			}
		}
	}
	return x, y, true
}
func parseInput(lines []string) ([][]string, []Path) {
	totalLines := len(lines)
	grid := make([][]string, totalLines-2)
	maxWidth := 0
	for i := 0; i < totalLines-2; i++ {
		line := lines[i]
		chars := []rune(line)
		width := len(chars)
		if width > maxWidth {
			maxWidth = width
		}
		grid[i] = make([]string, width)
		row := grid[i]
		for x := 0; x < width; x++ {
			row[x] = string(chars[x])
		}
	}

	for idx, row := range grid {
		if len(row) < maxWidth {
			difference := maxWidth - len(row)
			for i := 0; i < difference; i++ {
				row = append(row, " ")
			}
			grid[idx] = row
		}
	}

	commandLine := lines[totalLines-1]
	var turns []string
	steps := strings.FieldsFunc(commandLine, func(r rune) bool {
		char := string(r)
		if char == "R" || char == "L" {
			turns = append(turns, char)
			return true
		}
		return false
	})
	var paths []Path
	for i := 0; i < len(steps); i++ {
		step, _ := strconv.Atoi(steps[i])
		direction := ""
		if i < len(turns) {
			direction = turns[i]
		}
		path := Path{
			Steps:     step,
			Direction: direction,
		}
		paths = append(paths, path)
	}
	return grid, paths
}

type Facing string

const down Facing = "down"
const up Facing = "up"
const right Facing = "right"
const left Facing = "left"

type Path struct {
	Steps     int
	Direction string
}

type Position struct {
	X      int
	Y      int
	Facing Facing
}

func (p *Position) Turn(direction string) {
	if p.Facing == right {
		if direction == "R" {
			p.Facing = down
		} else if direction == "L" {
			p.Facing = up
		}
	} else if p.Facing == left {
		if direction == "R" {
			p.Facing = up
		} else if direction == "L" {
			p.Facing = down
		}
	} else if p.Facing == down {
		if direction == "R" {
			p.Facing = left
		} else if direction == "L" {
			p.Facing = right
		}
	} else if p.Facing == up {
		if direction == "R" {
			p.Facing = right
		} else if direction == "L" {
			p.Facing = left
		}
	}
}

func (p *Position) NextStepCoords() (int, int) {
	if p.Facing == right {
		return p.X + 1, p.Y
	}
	if p.Facing == left {
		return p.X - 1, p.Y
	}
	if p.Facing == up {
		return p.X, p.Y - 1
	}
	return p.X, p.Y + 1
}
