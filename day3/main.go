package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
)

func main(){
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum, count, group := 0,0,3 
	scanner := bufio.NewScanner(file)
	letters := make(map[string]int)
	for scanner.Scan(){
		line := scanner.Text()
		count++ 
		//sum += partOne(line)
		if count % group == 1 {
			count = 1
			letters = make(map[string]int)
		}
		splice := strings.Split(line,"")
		seenInLine := make(map[string]bool)

		for _, char := range splice {
			if !seenInLine[char] {
				letters[char]++
				seenInLine[char] = true
			}	
		}

		// when we reach the last line of the current group
		if count % group == 0{
			for key, val := range letters {
				if val == group {
					sum += getPriorityNumber(key)
					break // since there is only one 'key' per group
				}
			}
		}
	}

	fmt.Printf("Total: %d\n", sum)
	if err = scanner.Err(); err != nil{
		log.Fatal(err)
	}

}

func partOne(line string) int {
	sum := 0
	letters := make(map[string]bool)
	// 1st divide the string in half
	charSplice := strings.Split(line,"")
	firstHalf := charSplice[:len(charSplice)/2]
	secondHalf := charSplice[len(charSplice)/2:]
	// 2nd check which char is similar between the strings

	// First half insertion
	for _, char := range firstHalf {
		letters[char] = true
	}
	for _, char := range secondHalf {
		if letters[char] {
			// check priority level
			priority := getPriorityNumber(char)
			fmt.Printf("%s -> %d\n", char, priority)
			sum += priority
			break
		}
	}

	return sum
}

func getPriorityNumber(character string) int {
	asciiValue := int(character[0])
	if asciiValue > 97 {
		return asciiValue - 96
	}
	return asciiValue - 65 + 27
}

