package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readTextFile() ([]string, error) {
	// Open the file
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var schematic []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		schematic = append(schematic, line)
	}

	return schematic, nil
}

func traverseschematic(schematic []string) int {
	sum := 0
	size := len(schematic)
	number := ""
	gearPosition := -1
	gearMap := make(map[int][]int)

	moves := [][]int{
        {-1, -1}, {-1, 0}, {-1, 1},
        {0, -1},           {0, 1},
        {1, -1}, {1, 0}, {1, 1},
    }

	// travarse through schematic
	for posY := 0; posY < size; posY++ {
		for posX := 0; posX < size; posX++ {
			// check if schematic position is a number
			if (schematic[posY][posX] >= 48 && schematic[posY][posX] <= 57) {
				// build number
				number += string(schematic[posY][posX])
				// check if number is by a symbol, if not see if curr position is next to a symbol
				if (gearPosition == -1) {
					// check surronding position for symbol
					for _, move := range moves {
						newX, newY := posX+move[0], posY+move[1]
						// Check if the new position is within the board boundaries
						if (newX >= 0 && newX < size && newY >= 0 && newY < size) {
							if(schematic[newY][newX] == 42) {
								gearPosition = (newY*1000) + newX
								break
							}
						}
					}
				}
			} else {
				// reset number and symbol
				if (gearPosition != -1) {
					count, _ := strconv.Atoi(number)
					gearMap[gearPosition] = append(gearMap[gearPosition], count)
				}
				number = ""
				gearPosition = -1
			}
		}
	}

	for _, gears := range gearMap { 
		if (len(gears) > 1) {
			count := 1
			for _, number := range gears { 
				count *= number
			}
			sum += count
		}
	}

	return sum
}

func main() {
	schematic, err := readTextFile()
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(traverseschematic(schematic))
}
