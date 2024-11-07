package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regX := 1
	cycles := make([]int, 0, 250)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "noop") {
			cycles = append(cycles, regX)
		} else {
			cycles = append(cycles, regX, regX)
			parts := strings.Split(line, " ")
			if len(parts) != 2 {
				log.Fatal("error parsing the string")
			}
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal("error converting a number to int")
			}
			regX += val
		}
	}

	var strength int
	for i := 19; i <= 219; i += 40 {
		strength += (i + 1) * cycles[i]
		fmt.Printf("cycle[%d] = %d, strenght = %d\n", (i + 1), cycles[i], strength)
	}

	fmt.Printf("Result: %d \n", strength)
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
