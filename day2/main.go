package main

import (
	"fmt"
	"errors"
	"os"
	"log"
	"strings"
	"bufio"
)


func main(){
	file, err := os.Open("input.txt")
	if err != nil{
		log.Fatal(err)
	}

	defer file.Close()
	
	score := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, " ")
		if len(chars) != 2 {
			log.Fatal(errors.New("invalid line"))
		}
		// A -> rock ; B -> paper ; C -> scissors
		//Sum = +1 ; +2; +3
		//Sum+= 0 lost, 3 draw, 6 win

		// Strategy guide: Y -> Paper, X -> Rock, Z -> scissors
		opponent, me := chars[0], chars[1]
		if opponent == "A" {
			if me == "Y" {
				score += (2+6)
				continue
			}
			if me == "X" {
				score += (1 + 3)
				continue
			}
			score += (3 + 0)
			continue
		}
		if opponent == "B" {
			if me == "Y" {
				score += (2+3)
				continue
			}
			if me == "X" {
				score += (1 + 0) 
				continue 
			}
			score += (3 + 6)
			continue
		}
		// scissors
		if me == "Y" {
			score += (2+0)
			continue
		}
		if me == "X" {
			score += (1 + 6) 
			continue 
		}
		score += (3 + 3)
	}

	fmt.Printf("Score: %d\n", score)

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
}
