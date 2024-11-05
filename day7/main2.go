package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Directory struct {
	name     string
	size     int64
	parent   *Directory
	children map[string]*Directory
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		name:     name,
		size:     0,
		parent:   parent,
		children: make(map[string]*Directory),
	}
}

var root = NewDirectory("/", nil)
var current = root

func changeDirectory(path string) {
	switch path {
	case "/":
		current = root
	case "..":
		if current.parent != nil {
			current = current.parent
		}
	default:
		if child, ok := current.children[path]; ok {
			current = child
		} else {
			fmt.Printf("Directory %s does not exist.\n", path)
		}
	}
}

func printCurrentPath() {
	var path []string
	for dir := current; dir != nil; dir = dir.parent {
		path = append([]string{dir.name}, path...)
	}
	fmt.Println("Current path: ", strings.Join(path, "/"))
}

func getCommandInput(line string, command string) string {
	parts := strings.SplitN(line, command, 2)
	if len(parts) != 2 {
		fmt.Printf("couldn't parse command: %v\n", line)
		os.Exit(1)
	}
	return strings.TrimSpace(parts[1])
}

func traverseDirectory(dir *Directory, mm map[string]int64) int64 {
	totalSize := int64(dir.size)
	for _, child := range dir.children {
		totalSize += traverseDirectory(child, mm)
	}
	dir.size = totalSize
	mm[dir.name] = totalSize
	fmt.Printf("Directory %s, total size: %d \n", dir.name, dir.size)
	return totalSize
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	firstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		if firstLine {
			firstLine = false
			continue
		}
		if strings.HasPrefix(line, "$ cd") {
			dir := getCommandInput(line, "$ cd")
			changeDirectory(dir)
			continue
		}
		if strings.HasPrefix(line, "$ ls") {
			continue
		}
		if strings.HasPrefix(line, "dir") {
			dir := getCommandInput(line, "dir")
			current.children[dir] = NewDirectory(dir, current)
			continue
		}
		// we get to a file
		lsResult := strings.Split(line, " ")
		if len(lsResult) != 2 {
			fmt.Printf("error in line: %s\n", line)
			os.Exit(1)
		}
		size, err := strconv.ParseInt(lsResult[0], 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		current.size += size
	}

	dirMap := make(map[string]int64)
	used := traverseDirectory(root, dirMap)

	available := int64(70000000)
	required := int64(30000000)

	freeSpace := available - used
	needed := required - freeSpace

	smallestSufficientDirSize := int64(math.MaxInt64)

	for _, dirSize := range dirMap {
		// I need to delete a directory that would free up the needed size to meet the required space for the update
		if dirSize >= needed && dirSize < smallestSufficientDirSize {
			smallestSufficientDirSize = dirSize
		}
	}

	fmt.Printf("Smallest difference: %v\n", smallestSufficientDirSize)

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func absolute(val int64) int64 {
	if val < 0 {
		return -val
	}
	return val
}
