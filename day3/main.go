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

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
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
	}


	fmt.Printf("Total: %d\n", sum)
	if err = scanner.Err(); err != nil{
		log.Fatal(err)
	}

}

func getPriorityNumber(character string) int {
	asciiValue := int(character[0])
	if asciiValue > 97 {
		return asciiValue - 96
	}
	return asciiValue - 65 + 27
}

