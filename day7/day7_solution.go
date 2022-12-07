package main

import (
	"fmt"
	"math"
	"ravishrathod/adventofcode-2022/commons"
	"strconv"
	"strings"
	"time"
)

func main() {
	lines, err := commons.ReadFile("input/day7.txt")
	if err != nil {
		panic(err)
	}
	rootDir := parseInput(lines)
	sizeMap := calculateTotalSizes(rootDir)
	part1(sizeMap)
	part2(sizeMap)
}

func part2(sizeMap map[string]int64) {
	const totalSize = 70000000
	const minSpaceRequired = 30000000
	totalUsed := sizeMap["/"]
	freeSpaceAvailable := totalSize - totalUsed
	moreSpaceNeeded := minSpaceRequired - freeSpaceAvailable

	var sizeToPurge int64 = math.MaxInt

	for _, size := range sizeMap {
		if size >= moreSpaceNeeded {
			if size < sizeToPurge {
				sizeToPurge = size
			}
		}
	}
	println(sizeToPurge)
}

func part1(sizeMap map[string]int64) {
	var sizeToPurge int64 = 0
	for _, size := range sizeMap {
		if size <= 100000 {
			sizeToPurge += size
		}
	}
	println(sizeToPurge)
}

func calculateTotalSizes(root *directory) map[string]int64 {
	start := time.Now()
	sizeMap := make(map[string]int64)
	stack := &DirectoryStack{}
	populateStack(root, stack)
	for !stack.IsEmpty() {
		dir, _ := stack.Pop()
		totalSize := dir.GetSize()
		for _, child := range dir.Directories {
			totalSize += child.TotalSize
		}
		dir.TotalSize = totalSize
		key := getCanonicalName(dir)
		sizeMap[key] = dir.TotalSize
	}
	fmt.Printf("\nTime taken to calculate sizes: %v\n", time.Since(start))
	return sizeMap
}

func populateStack(dir *directory, stack *DirectoryStack) {
	stack.Push(dir)
	for _, child := range dir.Directories {
		populateStack(child, stack)
	}
}

func parseInput(lines []string) *directory {
	start := time.Now()
	rootDir := &directory{
		Name:        "/",
		Files:       []*file{},
		Directories: []*directory{},
	}
	directoryMap := make(map[string]*directory)
	directoryMap["/"] = rootDir
	currentDir := rootDir
	for index, line := range lines {
		if index == 0 {
			continue
		}
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			command := parts[1]
			if command == "cd" {
				destinationDirectoryName := parts[2]
				if destinationDirectoryName == ".." {
					currentDir = currentDir.Parent
					continue
				}
				destinationDir := createDirectory(destinationDirectoryName, currentDir)
				mapKey := getCanonicalName(destinationDir) //currentDir.Name + "-" + destinationDirectoryName
				if directoryMap[mapKey] == nil {
					currentDir.Directories = append(currentDir.Directories, destinationDir)
					directoryMap[mapKey] = destinationDir
				}
				currentDir = directoryMap[mapKey]
			} else if command == "ls" {
				continue
			}
		} else {
			if parts[0] == "dir" {
				dir := createDirectory(parts[1], currentDir)
				mapKey := getCanonicalName(dir)
				directoryMap[mapKey] = dir
				currentDir.Directories = append(currentDir.Directories, dir)
			} else {
				size, _ := strconv.Atoi(parts[0])
				f := &file{
					Name: parts[1],
					Size: int64(size),
				}
				currentDir.Files = append(currentDir.Files, f)
			}
		}
	}
	fmt.Printf("\nTime taken to parse: %v\n", time.Since(start))
	return rootDir
}

func createDirectory(name string, parent *directory) *directory {
	return &directory{
		Name:        name,
		Files:       []*file{},
		Directories: []*directory{},
		Parent: parent,
		TotalSize: -1,
	}
}

type directory struct {
	Name string
	Files []*file
	Directories []*directory
	Parent *directory
	TotalSize int64
}

func (d *directory) GetSize() int64 {
	var size int64 = 0
	for _, f := range d.Files {
		size += f.Size
	}
	return size
}

type file struct {
	Name string
	Size int64
}

func getCanonicalName(d *directory) string {
	return prependParentName(d.Name, d.Parent)
}

func prependParentName(input string, d *directory) string {
	name := input
	if d != nil {
		name = d.Name + "/" + input
		if d.Parent != nil {
			return prependParentName(name, d.Parent)
		}
	}
	return name
}
