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
	times := []int{}
	distances := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if (strings.HasPrefix(line, "Time: ")) {
			for _, time := range strings.Fields(strings.Split(line, ": ")[1]) {
				timeInt, _ := strconv.Atoi(time)
				times = append(times, timeInt)
			}
		} 
		if (strings.HasPrefix(line, "Distance: ")) {
			for _, dist := range strings.Fields(strings.Split(line, ": ")[1]) {
				distInt, _ := strconv.Atoi(dist)
				distances = append(distances, distInt)
			}
		}
	}	

	waysToWin := make([]int, len(distances))
	for index, dist := range distances {
		for i := 0; i < (times[index]); i++ {
			if ((i * (times[index]-i)) > dist) {
				waysToWin[index] += 1;
			}
		} 
	}

	total := 1
	for _, ways := range waysToWin {
		total *= ways
	}
	fmt.Println("Ways to win: ", total)
}
