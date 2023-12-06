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
	seeds := []int64{}
	skip := make(map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if ( line == "" || strings.HasSuffix(line, "map:")){
			skip = make(map[int]bool)
			continue
		} else if( strings.HasPrefix(line, "seeds: ") ) {
			seedsRaw := strings.Split(line, ": ")[1]
			for _, num := range strings.Split(seedsRaw, " "){
				numInt, _ := strconv.ParseInt(num, 10, 64)
				seeds = append(seeds, numInt)
			}
		} else {
			numbers := strings.Split(line, " ")
			srcDestination, _ := strconv.ParseInt(numbers[0], 10, 64)
			srcRange, _ := strconv.ParseInt(numbers[1], 10, 64)
			rangeLen, _ := strconv.ParseInt(numbers[2], 10, 64)
			for index, seed := range seeds {
				if skip[index] {
					continue
				}
				if ((srcRange <= seed && seed < (srcRange + rangeLen))) {
					seeds[index] = srcDestination + (seed - srcRange)
					skip[index] = true
				}
			}
		}
	}	
	minSeed := seeds[0]
	for _, seed := range seeds {
		if (seed < minSeed) { minSeed = seed }
	}
	fmt.Println(minSeed)
}
