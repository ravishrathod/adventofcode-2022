package main

import (
	"math"
	"ravishrathod/adventofcode-2022/commons"
	"strconv"
	"strings"
)

func main() {
	lines, err := commons.ReadFile("input/day14.txt")
	if err != nil {
		panic(err)
	}
	grid, maxDepth, minX, maxX := parseGrid(lines)
	part1(grid, maxDepth, minX, maxX)
	grid, maxDepth, minX, maxX = parseGrid(lines)
	part2(grid, maxDepth, minX, maxX)
}

func part1(grid map[Point]string, maxDepth int, minX int, maxX int) {
	calculateFlow(maxDepth, grid)
	printGrid(minX, maxX, maxDepth, grid)
}

func part2(grid map[Point]string, maxDepth int, minX int, maxX int) {
	floorY := maxDepth + 2

	for x := minX - 1000; x <= maxX+1000; x++ {
		point := Point{
			X: x,
			Y: floorY,
		}
		grid[point] = "rock"
	}
	calculateFlow(floorY, grid)
}

func calculateFlow(maxDepth int, grid map[Point]string) {
	sandParticlesFlown := 0
	particleSettled := true
	var particlePositions []Point
	for particleSettled {
		particlePos := Point{
			X: 500,
			Y: 0,
		}
		if !isEmptyAt(particlePos, grid) {
			break
		}
		particleSettled = false
		for particlePos.Y <= maxDepth {
			if !isEmptyAt(particlePos.down(), grid) {
				if !isEmptyAt(particlePos.downLeft(), grid) {
					if !isEmptyAt(particlePos.downRight(), grid) {
						grid[particlePos] = "sand"
						particlePositions = append(particlePositions, particlePos)
						sandParticlesFlown++
						particleSettled = true
						break
					} else {
						particlePos = particlePos.downRight()
					}
				} else {
					particlePos = particlePos.downLeft()
				}
			} else {
				particlePos = particlePos.down()
			}
		}
	}
	println(sandParticlesFlown)
}

func parseGrid(lines []string) (map[Point]string, int, int, int) {
	grid := make(map[Point]string)
	maxDepth := 0
	minX := math.MaxInt
	maxX := 0
	for _, line := range lines {
		coordsList := strings.Split(line, " -> ")
		for i := 0; i < len(coordsList); i++ {
			point := toPoint(coordsList[i])
			if point.Y > maxDepth {
				maxDepth = point.Y
			}
			if point.X < minX {
				minX = point.X
			}
			if point.X > maxX {
				maxX = point.X
			}
			if i == 0 {
				grid[point] = "rock"
			} else {
				lastPoint := toPoint(coordsList[i-1])
				if lastPoint.X == point.X {
					start, end := startEndAscending(lastPoint.Y, point.Y)
					for y := start; y <= end; y++ {
						point := Point{
							X: point.X,
							Y: y,
						}
						grid[point] = "rock"
					}
				} else {
					start, end := startEndAscending(lastPoint.X, point.X)
					for x := start; x <= end; x++ {
						point := Point{
							X: x,
							Y: point.Y,
						}
						grid[point] = "rock"
					}
				}
			}
		}
	}
	return grid, maxDepth, minX, maxX
}

func startEndAscending(a int, b int) (int, int) {
	start := a
	end := b
	if start > end {
		// swap
		start = start + end
		end = start - end
		start = start - end
	}
	return start, end
}

func printGrid(minX int, maxX int, maxDepth int, grid map[Point]string) {
	for y := 0; y <= maxDepth; y++ {
		for x := minX; x <= maxX; x++ {
			point := Point{
				X: x,
				Y: y,
			}
			symbol := "."
			if y == 0 && x == 500 {
				symbol = "x"
			}
			if object := grid[point]; object != "" {
				if object == "rock" {
					symbol = "#"
				} else {
					symbol = "o"
				}
			}
			print(symbol)
		}
		println()
	}
}

func isEmptyAt(point Point, grid map[Point]string) bool {
	object := grid[point]
	return object == ""
}
func toPoint(input string) Point {
	xy := strings.Split(input, ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	return Point{
		X: x,
		Y: y,
	}
}

type Point struct {
	X int
	Y int
}

func (p *Point) down() Point {
	return Point{
		X: p.X,
		Y: p.Y + 1,
	}
}

func (p *Point) downLeft() Point {
	return Point{
		X: p.X - 1,
		Y: p.Y + 1,
	}
}

func (p *Point) downRight() Point {
	return Point{
		X: p.X + 1,
		Y: p.Y + 1,
	}
}
