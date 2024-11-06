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
	// In part2, we have 10 knots
	// 0th will be the head, 9th will be the tail
	knots := make([]Position, 10)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		move := getMove(line)
		for step := 0; step < move.steps; step++ {
			knots[0] = getNewPosition(Move{direction: move.direction, steps: 1}, knots[0])

			knots[1] = moveTail(knots[0], knots[1])
			for i := 2; i < 10; i++ {
				knots[i] = moveTail(knots[i-1], knots[i])
			}

			visited[knots[9]] = true
		}
	}

	// the dictionary has the unique positions visited by the tail
	fmt.Printf("Result: %d \n", len(visited))
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func moveTail(first, second Position) Position {
	// check if they are not adjacent (only when this happens we move the tail)
	if abs(first.row-second.row) > 1 || abs(first.col-second.col) > 1 {
		// this logic allows for a diagonal move if needed
		if first.row > second.row {
			second.row++
		} else if first.row < second.row {
			second.row--
		}

		if first.col > second.col {
			second.col++
		} else if first.col < second.col {
			second.col--
		}
	}
	return second
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
