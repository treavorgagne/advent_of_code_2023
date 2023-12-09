package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func allZeros(arr []int) bool {
	for _, value := range arr {
		if value != 0 {
			return false
		}
	}
	return true
}

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		layer := 0
		var history [][]int // Declare a slice of slices
		row := []int{}
		for _, num := range strings.Split(line, " ") {
			numInt, _ := strconv.Atoi(num)
			row = append(row, numInt)
		}		
		history = append(history, row)
		for ;; {
			row = []int{}
			if (allZeros(history[layer])) {
				// find history
				histValue := 0
				for i := len(history)-1; i >= 0; i-- {
					histValue += history[i][len(history[i])-1]
				}
				total += histValue
				break
			} else {
				// add history value
				layer++
				for i := 1; i < len(history[layer-1]); i++ {
					row = append(row, history[layer-1][i] - history[layer-1][i-1])
				}
				history = append(history, row)
			}
		}
	}	

	fmt.Println("Sum extrapolated values: ", total)
}
