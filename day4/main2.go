package main

import (
	"fmt"
	"bufio"
	"strings"
	"errors"
	"log"
	"os"
	"strconv"
)

func main(){
	file, err := os.Open("input.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	
	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		elves := strings.Split(line,",")
		if len(elves) != 2 {
			log.Fatal(errors.New("incorrect parsing of the pair of elves"))
		}
		e1, e2 := elves[0], elves[1]
		leftE1, rightE1 := parseNumbers(e1)
		leftE2, rightE2 := parseNumbers(e2)

		if leftE1 <= rightE2 && rightE1 >= leftE2 {
		    count++
		}
	}

	fmt.Printf("Number of pairs intersect each range: %d\n", count)

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
}

// Func to return first and second number of each elf's section
func parseNumbers(str string) (int,int){
	line := strings.Split(str,"-")
	if len(line) != 2 {
		log.Fatal(errors.New("incorrect parsing of the pair of elves"))
	}
	num1, err := strconv.Atoi(line[0])
	if err != nil {
		log.Fatal(err)
	}
	num2, err := strconv.Atoi(line[1])
	if err != nil {
		log.Fatal(err)
	}
	return num1, num2
}

