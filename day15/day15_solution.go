package main

import (
	"fmt"
	"ravishrathod/adventofcode-2022/commons"
	"strconv"
	"strings"
	"time"
)

func main() {
	lines, err := commons.ReadFile("input/day15.txt")
	if err != nil {
		panic(err)
	}
	sensorToBeaconMap, beaconLocations := parseInput(lines)
	part1(sensorToBeaconMap, beaconLocations)
	//part2(sensorToBeaconMap)
}

func part1(sensorToBeaconMap map[Point]Point, beaconLocations map[Point]bool) {
	start := time.Now()
	pointsMonitored := make(map[Point]bool)
	rowToMonitor := 2000000
	for sensor, beacon := range sensorToBeaconMap {
		distance := calculateDistance(sensor, beacon)
		target := Point{
			X: sensor.X,
			Y: rowToMonitor,
		}
		if isFartherThan(sensor, target, distance) {
			continue
		}
		pointsMonitored[target] = true
		offset := 1
		for true {
			leftPoint := target.left(offset)
			if isFartherThan(sensor, leftPoint, distance) {
				break
			}
			pointsMonitored[leftPoint] = true
			pointsMonitored[target.right(offset)] = true
			offset++
		}
	}
	monitoredEmptyLoc := 0
	for location, _ := range pointsMonitored {
		if !beaconLocations[location] {
			monitoredEmptyLoc++
		}
	}
	println(monitoredEmptyLoc)
	fmt.Printf("Time taken: %v\n", time.Since(start))
}

func part2(sensorToBeaconMap map[Point]Point) {
	//a := 4000000
	upperLimit := 4000000
	//distressLoc := Point{}
	freqChan := make(chan int)
	for x := 0; x <= upperLimit; x++ {
		for y := 0; y <= upperLimit; y++ {
			//frequency := a*x + y
			//if !isMonitored(x, y, sensorToBeaconMap) {
			//	distressLoc.X = x
			//	distressLoc.Y = y
			//	fmt.Printf("\nDistress beacon at: %v", distressLoc)
			//	fmt.Printf("\nFrequency: %v", frequency)
			//	return
			//}
			go calculateDistressFrequency(x, y, sensorToBeaconMap, freqChan)
		}
	}
	frequency := <-freqChan
	println(frequency)
}
func calculateDistressFrequency(x int, y int, sensorToBeaconMap map[Point]Point, freqChan chan int) {
	//fmt.Printf("\ncalculating for (%v,%v)", x, y)
	a := 4000000
	frequency := a*x + y
	if !isMonitored(x, y, sensorToBeaconMap) {
		fmt.Printf("\nDistress beacon at: %v, %v", x, y)
		fmt.Printf("\nFrequency: %v\n", frequency)
		freqChan <- frequency
	}
}
func isMonitored(x int, y int, sensorToBeaconMap map[Point]Point) bool {
	target := Point{
		X: x,
		Y: y,
	}
	for sensor, beacon := range sensorToBeaconMap {
		distance := calculateDistance(sensor, beacon)
		if !isFartherThan(sensor, target, distance) {
			return true
		}
	}
	return false
}

func isFartherThan(source Point, target Point, distance int) bool {
	return calculateDistance(source, target) > distance
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

func parseInput(lines []string) (map[Point]Point, map[Point]bool) {
	sensorToBeaconMap := make(map[Point]Point)
	beaconLocations := make(map[Point]bool)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		sensorPart := parts[0]
		sensorPart = strings.Replace(sensorPart, "Sensor at ", "", -1)
		sensorLocation := parsePoint(sensorPart)

		beaconPart := parts[1]
		beaconPart = strings.Replace(beaconPart, " closest beacon is at ", "", -1)
		beaconLocation := parsePoint(beaconPart)

		sensorToBeaconMap[*sensorLocation] = *beaconLocation
		beaconLocations[*beaconLocation] = true
	}
	return sensorToBeaconMap, beaconLocations
}

// x*4000000 + y = 56000011
