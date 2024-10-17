package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"github.com/andrepintoo/adventofcode2022/utils"
)

func main(){
	st := utils.NewStack(int)
	st.Push(1)
	fmt.Printf("done")
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
