package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func getLineCalibration(input string) string {
	var result []string
	for _, char := range input {
		if unicode.IsDigit(char) {
			result = append(result, string(char))
		}
	}
	return result[0] + result[len(result)-1]
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
		line := scanner.Text()
		calibration, err := strconv.Atoi(getLineCalibration(line))
		if err != nil {
			fmt.Printf("Error converting string to integer: %s\n", err)
			os.Exit(1)
		}
		sum += calibration
	}

	fmt.Println(sum)
}
