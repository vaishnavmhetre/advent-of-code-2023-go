package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var maxQuant map[string]int = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	// file, err := os.Open("d2\\p1\\d2p1.initial.input")
	file, err := os.Open("d2\\p1\\d2p1.input")
	// file, err := os.Open("d2p1.input")
	if err != nil {
		log.Fatalln("Failed to open input file", err)
	}

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		gameNo, failed, err := ParseAndEval(line)
		if err != nil {
			log.Fatalf("Failed to parse. Reported %+v", err)
			os.Exit(1)
		} else if !failed {
			sum += gameNo
		} else {
			log.Printf("Failed Game #%02d :: %s", gameNo, line)
		}
	}

	log.Printf("Sum: %d", sum)
}

func ParseAndEval(line string) (gameNo int, failed bool, err error) {
	subsetsStr := strings.TrimPrefix(line, "Game ") // Result - "21: 4 red, ....."

	gameNumStr, subsetsStr, _ := strings.Cut(subsetsStr, ": ") // Result - "21" , "4 red, ...."
	gameNo, err = strconv.Atoi(gameNumStr)
	if err != nil {
		log.Fatalf("Failed to parse Game Number from %s", gameNumStr)
		return -1, true, err
	}

	subsets := strings.Split(subsetsStr, "; ") // Result - "4 red, 8 blue" , "6 red, 13 green"

	for _, subset := range subsets {
		for _, grab := range strings.Split(subset, ", ") { // Result - "4 red" , "8 blue"
			countStr, color, _ := strings.Cut(grab, " ") // Result - "4" , "red"
			count, _ := strconv.Atoi(countStr)
			if count > maxQuant[color] {
				return gameNo, true, nil
			}
		}

	}

	return gameNo, false, nil

}
