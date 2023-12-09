package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Function to calculate the Greatest Common Divisor (GCD) using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Function to calculate the Least Common Multiple (LCM) using GCD
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// Function to calculate the Least Common Multiple (LCM) for an array of integers
func lcmOfArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = lcm(result, arr[i])
	}
	return result
}

type Direction struct {
    right  string
    left string
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	directions := ""
	directionMap := make(map[string]Direction)
	positions := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if (line == "") {
			continue
		} else if (strings.Contains(line, " = ")) {
			elements := strings.Split(line, " = ")
			key := elements[0]
			if (strings.HasSuffix(key, "A")) {
				positions = append(positions, key)
			}
			dirs := strings.Split(elements[1], ", ")
			leftDir := dirs[0][1:]
			rightDir := dirs[1][:3]
			mapDir := Direction{right: rightDir, left: leftDir}
			directionMap[key] = mapDir;
		} else {
			directions = line
		}
	}	

	steps := 0
	lcmValues := []int{}
	for x := 1; x < len(positions); x++ {
		for i:=0;;i++ {
			i %= len(directions)
			rd := true
			if (!strings.HasSuffix(positions[x], "Z") || !strings.HasSuffix(positions[x-1], "Z")) {
				rd = false
			}
			if (rd) {
				break
			}
			if string(directions[i]) == "R" {
				positions[x] = directionMap[positions[x]].right
				positions[x-1] = directionMap[positions[x-1]].right
			} else if string(directions[i]) == "L" {
				positions[x] = directionMap[positions[x]].left
				positions[x-1] = directionMap[positions[x-1]].left
			}
			steps++
		}

		lcmValues = append(lcmValues, steps)
		steps = 0
	}

	fmt.Println("Step made: ", lcmOfArray(lcmValues))
}
