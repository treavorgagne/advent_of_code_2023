package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	times := 0
	distances := 0
	for scanner.Scan() {
		line := scanner.Text()
		if (strings.HasPrefix(line, "Time: ")) {
			time, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(line, ": ")[1], " ", ""))
			times = time
		} 
		if (strings.HasPrefix(line, "Distance: ")) {
			dist, _ := strconv.Atoi(strings.ReplaceAll(strings.Split(line, ": ")[1], " ", ""))
			distances = dist
		}
	}	

	waysToWin := 0
	for i := 0; i < (times); i++ {
		if ((i * (times-i)) > distances) {
			waysToWin += 1;
		}
	} 

	fmt.Println("Ways to win: ", waysToWin)
}
