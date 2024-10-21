package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	u "github.com/andrepintoo/adventofcode2022/utils"
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

	fmt.Println(line)

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}


}

