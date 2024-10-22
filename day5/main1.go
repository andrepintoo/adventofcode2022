package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
	"regexp"
	"github.com/andrepintoo/adventofcode2022/utils"	
)

type MoveBoxes struct {
	moves int
	from int
	to int
}
//"github.com/andrepintoo/adventofcode2022/utils"
func main(){
	file, err := os.Open("input.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	
	numRows := 9
	matrix := make([][]string, numRows)
	moves := make([]MoveBoxes, 0, 30)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		runeLine := []rune(line)
		// Get the matrix input:
		if !strings.HasPrefix(line,"move") {
			// each character appears, in each row, in a (1 + column * 4) pattern
			for i := 0; i < numRows; i++ {
				pattern := 1 + (i*4)
				char := runeLine[pattern]
				// check if it's not empty space and if it's between A-Z in ascii values
				if char != ' ' && char >= 'A' && char <= 'Z' {
					matrix[i] = append(matrix[i], string(char))
				}	
			}
		}else{
		// Get the moves needed
			err = fillMovements(line, &moves)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}
	}

	// for i:=0; i < numRows; i++ {
	// 	fmt.Println(matrix[i])
	// }
	// for i:=0; i<len(moves);i++ {
	// 	fmt.Println(moves[i])
	// }

	// Do the moves onto the crates
	for i:=0; i < len(moves); i++ {
		mbox := moves[i]
		m,f,t := mbox.moves, mbox.from, mbox.to
		var aux string
		for j:=0; j < m; j++ {
			if len(matrix[f]) == 0 {
				fmt.Printf("Error: Stack %d is empty, cannot move more crates\n", f)
				return
			}
			aux = matrix[f][0] // get the first element
			matrix[f] = utils.RemoveIndex[string](matrix[f], 0) // remove the first element
			matrix[t] = append([]string{aux}, matrix[t]...) // append the removed element to the 'to' slice
		}
	}

	fmt.Println("finished")
	var result string
	for i := 0; i < numRows; i++{
		result += matrix[i][0]
	}

	fmt.Printf("Final result is: %s\n", result)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func fillMovements(line string, moves *[]MoveBoxes) error {
	re := regexp.MustCompile("[0-9]+")
	numbers := re.FindAllString(line, -1)
	if len(numbers) != 3 {
		return fmt.Errorf("line '%s' doesn't contain correct instruction", line)
	}
	
	moveCount, err := strconv.Atoi(numbers[0])
	if err != nil {
		return fmt.Errorf("invalid move count: %s", numbers[0])
	}
	from, err := strconv.Atoi(numbers[1])
	if err != nil {
		return fmt.Errorf("invalid 'from' value: %s", numbers[1])
	}
	to, err := strconv.Atoi(numbers[2])
	if err != nil {
		return fmt.Errorf("invalid 'to' value: %s", numbers[2])
	}

	// -1 since the problem description is 1-based and not 0-based
	newMove := MoveBoxes{moves: moveCount, from: from - 1, to: to - 1}
	*moves = append(*moves, newMove)
	return nil
}
