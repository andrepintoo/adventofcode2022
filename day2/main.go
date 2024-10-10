package main

import (
	"fmt"
)

// A -> rock ; B -> paper ; C -> scissors
//Sum = +1 ; +2; +3
//Sum+= 0 lost, 3 draw, 6 win

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
