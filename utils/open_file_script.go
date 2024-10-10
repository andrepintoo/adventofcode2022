package main

import (
	"fmt"
)

func main(){
	file, err := os.Open("input.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
}
