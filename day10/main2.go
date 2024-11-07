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
	spriteMiddlePosition := 1
	cycles := make([]string, 0, 250)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "noop") {
			// 1 cycle per "noop" operation
			drawPixel(&cycles, spriteMiddlePosition)
		} else {
			// 2 cycles per "addx" operation"
			drawPixel(&cycles, spriteMiddlePosition)
			drawPixel(&cycles, spriteMiddlePosition)
			parts := strings.Split(line, " ")
			if len(parts) != 2 {
				log.Fatal("error parsing the string")
			}
			val, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatal("error converting a number to int")
			}
			spriteMiddlePosition += val
		}
	}

	for i := 0; i < len(cycles); i += 40 {
		end := i + 40
		if end > len(cycles) {
			end = len(cycles)
		}
		fmt.Printf("%v\n", strings.Join(cycles[i:end], ""))
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func drawPixel(cycles *[]string, registerValue int) {
	// gives the position within a row (from 0 to 39)
	// the sprite will only span 40 pixels
	currentCycle := len(*cycles) % 40
	if currentCycle >= (registerValue-1) &&
		currentCycle <= (registerValue+1) {
		*cycles = append(*cycles, "#")
	} else {
		*cycles = append(*cycles, ".")
	}
}
