package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func getFirstStringNumber(input string) string {
	pattern := "(1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine)"
	re := regexp.MustCompile(pattern)
	return re.FindString(input)
}

func getLastStringNumber(input string) string {
	pattern := "(1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine)"
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(input, -1)
	return matches[len(matches)-1]
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

	numObj := map[string]string{        
		"one": "1",
        "two":  "2",
        "three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
		"1": "1",
		"2": "2",
		"3": "3",
		"4": "4",
		"5": "5",
		"6": "6",
		"7": "7",
		"8": "8",
		"9": "9",
    }

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		calibration, err := strconv.Atoi(numObj[getFirstStringNumber(line)] + numObj[getLastStringNumber(line)])
		if err != nil {
			fmt.Printf("Error converting second number: %s\n", err)
			continue
		}
		sum += calibration
	}

	fmt.Println(sum)
}
