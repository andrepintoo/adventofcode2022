package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Directory struct {
	name     string
	size     int
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

// TODO: todo
// FIXME: fixme
// WARNING: warning

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result int
	sizeMap := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "$ cd") {

		}
	}

	fmt.Printf("Result: %v\n", result)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Initial directory setup:")
	// printCurrentPath() // Should show "/"
	//
	// // Add some directories for testing
	// current.children["subdir1"] = NewDirectory("subdir1", current)
	// current.children["subdir2"] = NewDirectory("subdir2", current)
	//
	// // Move into "subdir1"
	// changeDirectory("subdir1")
	// printCurrentPath() // Should show "/subdir1"
	//
	// current.children["novo"] = NewDirectory("novo", current)
	// changeDirectory("novo")
	// printCurrentPath() // Should show "/"
	// // Move back to root
	// changeDirectory("/")
	// printCurrentPath() // Should show "/"
	//
	// // Move into "subdir2"
	// changeDirectory("subdir2")
	// printCurrentPath() // Should show "/subdir2"
}
