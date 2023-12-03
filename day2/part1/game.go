package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isGamePossible(input string, gameId int) bool {
	input = strings.Split(input, ":")[1]
	sets := strings.Split(input, ";")
	pieces := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}   
    colors := [3]string{"red", "green", "blue"}
	
	for _, set := range sets {
		for _, color := range colors {
			pattern := regexp.MustCompile("(\\d+) " + color)
			matches := pattern.FindAllString(set, -1)
			for _, subValue := range matches {
				count, _ := strconv.Atoi(strings.Split(subValue, " ")[0])
				if (pieces[color] < count) {
					return false
				}
			}
		}
	}
	return true
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
	gameId := 1
	for scanner.Scan() {
		game := scanner.Text()
		if (isGamePossible(game, gameId)) {
			sum += gameId
		}
		gameId += 1
	}

	fmt.Println("total sum passed:",sum)
}
