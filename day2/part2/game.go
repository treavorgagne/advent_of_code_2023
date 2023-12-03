package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func gamesPower(input string) int {
	input = strings.Split(input, ":")[1]
	sets := strings.Split(input, ";")
	pieces := map[string]int{
		"red": 0,
		"green": 0,
		"blue": 0,
	}   
    colors := [3]string{"red", "green", "blue"}
	
	for _, set := range sets {
		for _, color := range colors {
			pattern := regexp.MustCompile("(\\d+) " + color)
			matches := pattern.FindAllString(set, -1)
			for _, subValue := range matches {
				count, _ := strconv.Atoi(strings.Split(subValue, " ")[0])
				if (pieces[color] < count) {
					pieces[color] = count
				}
			}
		}
	}
	return pieces["red"] * pieces["green"] * pieces["blue"]
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		game := scanner.Text()
		sum += gamesPower(game)
	}

	fmt.Println("total sum of power:",sum)
}
