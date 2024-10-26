package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

type Directory struct {
	name string
	size int
	children map[string]*Directory
}

func main(){
	file, err := os.Open("example.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()

	sizeM := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
}
