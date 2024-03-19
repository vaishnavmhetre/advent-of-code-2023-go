package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	// file, err := os.Open("d1\\p1\\d1p1.initial.input")
	file, err := os.Open("d1\\p1\\d1p1.input")
	if err != nil {
		log.Fatalln("Failed to open input file", err)
	}

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		lineBytes := scanner.Text()

		firstDigit, lastDigit := byte('0'), byte('0')

		if firstDigitIdx := strings.IndexFunc(lineBytes, func(r rune) bool { return unicode.IsDigit(r) }); firstDigitIdx != -1 {
			firstDigit = lineBytes[firstDigitIdx]
		}

		if lastDigitIdx := strings.LastIndexFunc(lineBytes, func(r rune) bool { return unicode.IsDigit(r) }); lastDigitIdx != -1 {
			lastDigit = lineBytes[lastDigitIdx]
		}

		numStr := string([]byte{firstDigit, lastDigit})
		num, err := strconv.Atoi(numStr)
		if err != nil {
			log.Fatalln("Something fucked up while parsing number from ", numStr)
		}
		sum += num

	}

	print(sum)
}
