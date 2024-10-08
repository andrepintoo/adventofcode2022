package main

import (
	 "fmt"
	 "log"
	 "os"
	 "bufio"
	 "strconv"
)

func main(){
	fmt.Println("Day 1")
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	max,max2,max3 := 0,0,0
	cur := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		
		if len(line)==0{
			if cur > max {
				max3 = max2
				max2 = max
				max = cur
			} else if cur > max2 {
				max3 = max2
				max2 = cur
			} else if cur > max3 { 
				max3 = cur
			}
			cur = 0	// reset to the next elf	
			continue
		}
		
		aux, err := strconv.Atoi(line)
		if err != nil{
			log.Fatal(err)
		}
		cur += aux
	}
	
	if cur > 0 {
		if cur > max {
			max3 = max2
			max2 = max
			max = cur
		} else if cur > max2 {
			max3 = max2
			max2 = cur
		} else if max3 > cur { 
			max3 = cur
		}
	}
			
	fmt.Printf("Max calories is %d\n", max)
	fmt.Printf("Total of top 3 elves is %d\n", max+max2+max3)

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
