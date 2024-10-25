package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"errors"
)

func main(){
	file, err := os.Open("input.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	if !scanner.Scan() {
		log.Fatal(errors.New("invalid text file"))
	}
	line := scanner.Text()

	if len(line) == 0 {
		log.Fatal(errors.New("invalid string"))
	}
	marker := 14
	seen := make([]string, 0, marker)
	for i, char := range line {

		seen = append(seen, string(char))	

		if len(seen) > marker {
			seen = seen[1:]
		}

		// check at every 4 characters
		if len(seen) == marker && !hasDuplicates(seen) {
			fmt.Printf("First marker at character %d\n", i+1)
			return
			
		}
	}

	fmt.Println("no marker was found")

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
}

func hasDuplicates(s []string) bool {
	fmt.Println(s)
	m := make(map[string]bool)
	for _, c := range s {
		if _, ok := m[c]; ok {
			return true
		}
		m[c] = true
	}
	return false
}
