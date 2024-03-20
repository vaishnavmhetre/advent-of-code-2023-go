package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// file, err := os.Open("d2\\p2\\d2p2.initial.input")
	file, err := os.Open("d2\\p2\\d2p2.input")
	// file, err := os.Open("d2p2.initial.input")
	if err != nil {
		log.Fatalln("Failed to open input file", err)
	}

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		power := ParseAndEval(line)
		sum += power
	}

	log.Printf("Sum: %d", sum)
}

func ParseAndEval(line string) (power int) {
	maxCountMap := make(map[string]int)
	colonIndex := strings.Index(line, ":")
	subsetsStr := line[colonIndex+2:]          // Remove space too, hence ' ' // Result - "4 red, ...."
	subsets := strings.Split(subsetsStr, "; ") // Result - "4 red, 8 blue" , "6 red, 13 green"

	for _, subset := range subsets {
		for _, grab := range strings.Split(subset, ", ") { // Result - "4 red" , "8 blue"
			countStr, color, _ := strings.Cut(grab, " ") // Result - "4" , "red"
			count, _ := strconv.Atoi(countStr)
			if count > maxCountMap[color] {
				maxCountMap[color] = count
			}
		}

	}

	power = 1

	for _, count := range maxCountMap {
		power *= count
	}

	return power

}
