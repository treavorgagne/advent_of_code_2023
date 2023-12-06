package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Set map[string]struct{}

func (s Set) Add(item string) {
    s[item] = struct{}{}
}

func (s Set) Remove(item string) {
    delete(s, item)
}

func (s Set) Contains(item string) bool {
    _, exists := s[item]
    return exists
}

func scatchTickets(game string) float64 {
	matches := -1

	numbers := strings.Split(strings.Split(game, ": ")[1], " | ")
	winningNums := make(Set)
	for _, num := range strings.Split(numbers[0], " ") {
		winningNums.Add(num)
	}
	winningNums.Remove("")

	for _, num := range strings.Split(numbers[1], " ") {
		if(winningNums.Contains(num)) {
			matches += 1
		}
	}

	if (matches == -1) {
		return 0
	}

	return math.Pow(2, float64(matches))
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	sum := 0.0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scratchTicket := scanner.Text()
		sum += scatchTickets(scratchTicket)
	}
	fmt.Println(sum)
}
