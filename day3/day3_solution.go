package main

import "ravishrathod/adventofcode-2022/commons"

func main() {
	lines, err := commons.ReadFile("input/day3.txt")
	if err != nil {
		panic(err)
	}
	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	totalPriority := 0
	for _, line := range lines {
		items := []rune(line)
		leftCompItems, rightCompItems := splitItemsByCompartment(items)
		intersectingValues := findIntersectingValues(leftCompItems, rightCompItems)
		totalPriority += calculatePriority(intersectingValues)
	}
	println(totalPriority)
}

func part2(lines []string) {
	totalPriority := 0
	for i := 0; i <= len(lines)-3; i = i + 3 {
		badge := findBadgeItem(lines[i], lines[i+1], lines[i+2])
		totalPriority += calculatePriority([]rune{badge})
	}
	println(totalPriority)

}

func findBadgeItem(line1 string, line2 string, line3 string) rune {
	firstBag := itemListToMap([]rune(line1))
	secondBag := itemListToMap([]rune(line2))
	thirdBag := itemListToMap([]rune(line3))
	return findCommonItem(firstBag, secondBag, thirdBag)
}

func findCommonItem(bag map[rune]bool, bag2 map[rune]bool, bag3 map[rune]bool) rune {
	for key, _ := range bag {
		if bag2[key] && bag3[key] {
			return key
		}
	}
	return 0
}

func calculatePriority(items []rune) int {
	totalPriority := 0
	for _, item := range items {
		charIndex := int(item)
		if charIndex >= 97 {
			priority := charIndex - 96
			totalPriority += priority
		} else {
			priority := charIndex - 38
			totalPriority += priority
		}
	}
	return totalPriority
}

func findIntersectingValues(left map[rune]bool, right map[rune]bool) []rune {
	var intersect []rune

	for k, _ := range left {
		if right[k] {
			intersect = append(intersect, k)
		}
	}
	return intersect
}

func splitItemsByCompartment(items []rune) (map[rune]bool, map[rune]bool) {
	firstCompartmentItems := items[0 : len(items)/2]
	secondCompartmentItems := items[len(items)/2:]

	return itemListToMap(firstCompartmentItems), itemListToMap(secondCompartmentItems)
}

func itemListToMap(items []rune) map[rune]bool {
	itemsMap := make(map[rune]bool)
	for _, item := range items {
		itemsMap[item] = true
	}
	return itemsMap
}
