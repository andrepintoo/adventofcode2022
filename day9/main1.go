package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	row int
	col int
}

type Move struct {
	direction string
	steps     int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	visited := make(map[Position]bool)
	// 	matrix := [][]string{"H"} // start at 0,0 with H and T overlapped
	headPos := Position{row: 0, col: 0}
	tailPos := Position{row: 0, col: 0}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		move := getMove(line)
		// 		fmt.Printf("move -> %v\n", move)
		// move head according to the input given
		// 		newHeadPosition := getNewPosition(move, headPos)
		// 		setMatrixValue(&matrix, "H", newHeadPosition.row, newHeadPosition.col)
		for step := 0; step < move.steps; step++ {
			headPos = getNewPosition(Move{direction: move.direction, steps: 1}, headPos)

			// update tail to follow head movement
			tailPos = moveTail(headPos, tailPos)

			visited[tailPos] = true
		}

	}

	// the dictionary has the unique positions visited by the tail
	fmt.Printf("Result: %d \n", len(visited))
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

// func setMatrixValue(matrix *[][]string, value string, row, col int){
// 	// expand it if necessary
// 	for len(*matrix) <= row {
// 		*matrix = append(*matrix, []string{})
// 	}
//
// 	for len((*matrix)[row]) <= col {
// 		(*matrix)[row] = append((*matrix)[row], "")
// 	}
//
// 	// set the value at the given position
// 	(*matrix)[row][col] = value
// }

func moveTail(head, tail Position) Position {
	// check if they are not adjacent (only when this happens we move the tail)
	if abs(head.row-tail.row) > 1 || abs(head.col-tail.col) > 1 {
		// this logic allows for a diagonal move if needed
		if head.row > tail.row {
			tail.row++
		} else if head.row < tail.row {
			tail.row--
		}

		if head.col > tail.col {
			tail.col++
		} else if head.col < tail.col {
			tail.col--
		}
	}
	return tail
}

func getMove(line string) Move {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		log.Fatal("wrong conversion")
	}
	steps, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	return Move{direction: parts[0], steps: steps}
}

func getNewPosition(move Move, pos Position) Position {
	// right
	if move.direction == "R" {
		return Position{pos.row, pos.col + move.steps}
	}

	// left
	if move.direction == "L" {
		return Position{pos.row, pos.col - move.steps}
	}

	// up
	if move.direction == "U" {
		return Position{pos.row - move.steps, pos.col}
	}

	// down
	return Position{pos.row + move.steps, pos.col}
}

// helper function:
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
