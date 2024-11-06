package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Advent of Code 2022 - Day 8")
	var matrix [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		trees := strings.Split(line, "")
		for _, v := range trees {
			v, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal("wrong conversion")
			}
			row = append(row, v)
		}
		matrix = append(matrix, row)
	}

	var visibleTrees int
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			tree := matrix[i][j]

			// Check if visible from the left
			visibleLeft := true
			for k := 0; k < j; k++ {
				if matrix[i][k] >= tree {
					visibleLeft = false
					break
				}
			}

			// Check if visible from the right
			visibleRight := true
			for k := j + 1; k < len(matrix[i]); k++ {
				if matrix[i][k] >= tree {
					visibleRight = false
					break
				}
			}

			// Check if visible from the top
			visibleTop := true
			for k := 0; k < i; k++ {
				if matrix[k][j] >= tree {
					visibleTop = false
					break
				}
			}

			// Check if visible from the bottom
			visibleBottom := true
			for k := i + 1; k < len(matrix); k++ {
				if matrix[k][j] >= tree {
					visibleBottom = false
					break
				}
			}

			// If visible from any direction, count it
			if visibleLeft || visibleRight || visibleTop || visibleBottom {
				visibleTrees += 1
			}
		}
	}

	// add edges and subtract 4 for the corners (that are being counted twice)
	visibleTrees += len(matrix)*2 + len(matrix[0])*2 - 4

	fmt.Printf("Result: %d \n", visibleTrees)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
