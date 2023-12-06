package main

import (
	"bufio"
	"fmt"
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

func countTotalTickets() int {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()
	scratches := make(map[int]int)

	scanner := bufio.NewScanner(file)
	games := []string{}
	index := 1
	for scanner.Scan() {
		ticket := scanner.Text()
		games = append(games, strings.Split(ticket, ": ")[1])
		scratches[index] = 1
		index++
	}

	for index, game := range games {
		winningNums := make(Set)
		matches := 1
		numbers := strings.Split(game, " | ")
		for _, num := range strings.Split(numbers[0], " ") {
			winningNums.Add(num)
		}
		winningNums.Remove("")

		for _, num := range strings.Split(numbers[1], " ") {
			if(winningNums.Contains(num)) {
				scratches[index + 1 + matches] += (1 * scratches[index+1])
				matches += 1
			}
		}
	}
	total := 0
	for _, num := range scratches {
		total += num
	}
	return total
}

func main() {
	fmt.Println(countTotalTickets())
}
