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

	var maxScenicScore int
	for i := 1; i < len(matrix)-1; i++ {
		for j := 1; j < len(matrix[i])-1; j++ {
			tree := matrix[i][j]

			// left
			leftScore := 0
			for k := j - 1; k >= 0; k-- {
				leftScore += 1
				if matrix[i][k] >= tree {
					break
				}
			}

			// right
			rightScore := 0
			for k := j + 1; k < len(matrix[i]); k++ {
				rightScore += 1
				if matrix[i][k] >= tree {
					break
				}
			}

			// top
			topScore := 0
			for k := i - 1; k >= 0; k-- {
				topScore += 1
				if matrix[k][j] >= tree {
					break
				}
			}

			// bottom
			bottomScore := 0
			for k := i + 1; k < len(matrix); k++ {
				bottomScore += 1
				if matrix[k][j] >= tree {
					break
				}
			}

			treeScore := leftScore * rightScore * topScore * bottomScore
			if treeScore > maxScenicScore {
				maxScenicScore = treeScore
			}
		}
	}

	fmt.Printf("Result: %d \n", maxScenicScore)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
