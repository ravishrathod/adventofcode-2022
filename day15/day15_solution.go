package main

import (
	"fmt"
	"math"
	"ravishrathod/adventofcode-2022/commons"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	lines, err := commons.ReadFile("input/day15.txt")
	if err != nil {
		panic(err)
	}
	sensorToBeaconMap := parseInput(lines)
	part1(sensorToBeaconMap)
	part2(sensorToBeaconMap)
}

func part1(sensorToBeaconMap map[Point]Point) {
	start := time.Now()
	rowToMonitor := 2000000
	coverages := calculateCoverageAtRow(sensorToBeaconMap, rowToMonitor)

	beaconsAtX := make(map[int]bool)
	for _, beacon := range sensorToBeaconMap {
		if beacon.Y == rowToMonitor {
			beaconsAtX[beacon.X] = true
		}
	}
	totalPointsCovered := 0
	beaconsCovered := make(map[int]bool)
	for i := 0; i < len(coverages); i++ {
		if i == 0 {
			current := coverages[i]
			last := Coverage{
				From: current.From - 1,
				To:   current.From - 1,
			}
			pointsCovered := findTrueCoverage(last, current, beaconsAtX, beaconsCovered)
			totalPointsCovered += pointsCovered
		} else {
			last := coverages[i-1]
			current := coverages[i]
			pointsCovered := findTrueCoverage(last, current, beaconsAtX, beaconsCovered)
			totalPointsCovered += pointsCovered
		}
	}
	println(totalPointsCovered - len(beaconsCovered))
	fmt.Printf("Time taken: %v\n", time.Since(start))
}

func part2(sensorToBeaconMap map[Point]Point) {
	start := time.Now()
	upperLimit := 4000000
	for y := 0; y <= upperLimit; y++ {
		coverages := calculateCoverageAtRow(sensorToBeaconMap, y)
		x := findUnScannedPoint(coverages, 0, upperLimit)
		if x >= 0 {
			frequency := x*4000000 + y
			println(frequency)
			fmt.Printf("Time taken: %v\n", time.Since(start))
			return
		}
	}
	fmt.Printf("Time taken: %v\n", time.Since(start))
	println("failed to find empty points")
}

func calculateCoverageAtRow(sensorToBeaconMap map[Point]Point, rowToMonitor int) []Coverage {
	var coverages []Coverage
	for sensor, beacon := range sensorToBeaconMap {
		monitoringDistance := calculateDistance(sensor, beacon)
		nearestPointOnLine := Point{
			X: sensor.X,
			Y: rowToMonitor,
		}
		shortestDistanceToLine := calculateDistance(sensor, nearestPointOnLine)
		if shortestDistanceToLine > monitoringDistance {
			continue
		}
		sweep := monitoringDistance - shortestDistanceToLine
		coverage := Coverage{
			From: nearestPointOnLine.X - sweep,
			To:   nearestPointOnLine.X + sweep,
		}
		coverages = append(coverages, coverage)
	}
	sort.Slice(coverages, func(i, j int) bool {
		if coverages[i].From == coverages[j].From {
			return coverages[i].To < coverages[j].To
		}
		return coverages[i].From < coverages[j].From
	})
	return coverages
}

func findTrueCoverage(last Coverage, current Coverage, beaconsAtX map[int]bool, beaconsCovered map[int]bool) int {
	coverage := int(math.Abs(float64(current.To-current.From))) + 1
	overlap := 0
	if current.From > last.To {
		overlap = 0
	} else {
		overlapFrom := current.From
		overlapTo := 0
		if current.To <= last.To {
			overlapTo = current.To
		} else {
			overlapTo = last.To
		}
		overlap = int(math.Abs(float64(overlapTo-overlapFrom))) + 1
	}
	for x, _ := range beaconsAtX {
		if x >= current.From && x <= current.To {
			beaconsCovered[x] = true
		}
	}
	adjustedCoverage := coverage - overlap
	if adjustedCoverage < 0 {
		adjustedCoverage = 0
	}
	return adjustedCoverage
}

func findUnScannedPoint(coverages []Coverage, lowerLimit int, upperLimit int) int {
	maxX := 0
	for i := 1; i < len(coverages); i++ {
		last := coverages[i-1]
		current := coverages[i]
		if current.From > last.To+1 {
			possiblePoint := current.From - 1
			if possiblePoint <= upperLimit && possiblePoint >= lowerLimit && possiblePoint > maxX {
				return possiblePoint
			}
		}
		currentMaxX := int(math.Max(float64(last.To), float64(current.To)))
		if currentMaxX > maxX {
			maxX = currentMaxX
		}
	}
	if maxX <= upperLimit && maxX >= lowerLimit {
		return maxX
	}
	return -1
}

func calculateDistance(from Point, to Point) int {
	deltaX := from.X - to.X
	if deltaX < 0 {
		deltaX *= -1
	}
	deltaY := from.Y - to.Y
	if deltaY < 0 {
		deltaY *= -1
	}
	return deltaX + deltaY
}
func parsePoint(input string) *Point {
	parts := strings.Split(input, ",")
	xStr := strings.Split(parts[0], "=")[1]
	x, _ := strconv.Atoi(xStr)
	yStr := strings.Split(strings.TrimSpace(parts[1]), "=")[1]
	y, _ := strconv.Atoi(yStr)

	return &Point{
		X: x,
		Y: y,
	}
}

type Point struct {
	X int
	Y int
}

func (p *Point) left(offset int) Point {
	return Point{
		X: p.X - offset,
		Y: p.Y,
	}
}

func (p *Point) right(offset int) Point {
	return Point{
		X: p.X + offset,
		Y: p.Y,
	}
}

type Coverage struct {
	From int
	To   int
}

func parseInput(lines []string) map[Point]Point {
	sensorToBeaconMap := make(map[Point]Point)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		sensorPart := parts[0]
		sensorPart = strings.Replace(sensorPart, "Sensor at ", "", -1)
		sensorLocation := parsePoint(sensorPart)

		beaconPart := parts[1]
		beaconPart = strings.Replace(beaconPart, " closest beacon is at ", "", -1)
		beaconLocation := parsePoint(beaconPart)

		sensorToBeaconMap[*sensorLocation] = *beaconLocation
	}
	return sensorToBeaconMap
}
