package main

import "ravishrathod/adventofcode-2022/commons"

func main() {
	lines, err := commons.ReadFile("input/day6.txt")
	if err != nil {
		panic(err)
	}
	characters := []rune(lines[0])
	part1(characters)
	part2(characters)
}

func part1(characters []rune) {
	println(findFirstUniqueSubStringOfLength(characters, 4))
}

func part2(characters []rune) {
	println(findFirstUniqueSubStringOfLength(characters, 14))
}

func findFirstUniqueSubStringOfLength(characters []rune, length int) int {
	var lastXChars []rune
	for idx, char := range characters {
		if len(lastXChars) == length {
			if isUnique(lastXChars) {
				return idx
			} else {
				lastXChars = lastXChars[1:]
				lastXChars = append(lastXChars, char)
			}
		} else {
			lastXChars = append(lastXChars, char)
		}
	}
	return -1
}

func isUnique(chars []rune) bool {
	charMap := make(map[rune]bool)
	for _, char := range chars {
		if charMap[char] {
			return false
		}
		charMap[char] = true
	}
	return true
}