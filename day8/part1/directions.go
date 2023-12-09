package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
	for scanner.Scan() {
		line := scanner.Text()
		if (line == "") {
			continue
		} else if (strings.Contains(line, " = ")) {
			elements := strings.Split(line, " = ")
			key := elements[0]
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

	pos := "AAA"
	size := len(directions)
	i := 0
	for {
		if pos == "ZZZ" {
			break
		} else if string(directions[i%size]) == "R" {
			pos = directionMap[pos].right
			steps++
		} else if string(directions[i%size]) == "L" {
			pos = directionMap[pos].left
			steps++
		}
		i++
	}

	fmt.Println("Step made: ", steps)
}
