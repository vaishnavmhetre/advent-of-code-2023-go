package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strings"
	"unicode"
)

var NumberWords = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	// file, err := os.Open("d1\\p2\\d1p2.initial.input")
	file, err := os.Open("d1\\p2\\d1p2.input")
	if err != nil {
		log.Fatalln("Failed to open input file", err)
	}

	scanner := bufio.NewScanner(file)

	var sum int64 = 0

	var lineNum = 0

	for scanner.Scan() {
		line := scanner.Text()
		lineNum++

		firstDigit, lastDigit := 0, 0

		firstDigitIdx := strings.IndexFunc(line, unicode.IsDigit)
		firstNumTxt, firstNumIdx, firstNumExists := FindFirstOccurance(line, NumberWords...)
		if firstNumExists && (firstDigitIdx == -1 || firstNumIdx < firstDigitIdx) {
			firstDigit = TextToByte(firstNumTxt)
		} else if firstDigitIdx != -1 {
			firstDigit = int(line[firstDigitIdx] - '0')
		}

		lastDigitIdx := strings.LastIndexFunc(line, unicode.IsDigit)
		lastNumTxt, lastNumIdx, lastNumExists := FindLastOccurance(line, NumberWords...)
		if lastNumExists && (lastDigitIdx == -1 || lastNumIdx > lastDigitIdx) {
			lastDigit = TextToByte(lastNumTxt)
		} else if lastDigitIdx != -1 {
			lastDigit = int(line[lastDigitIdx] - '0')
		}

		log.Printf("Line: %d :: Num %d%d", lineNum, firstDigit, lastDigit)

		num := (firstDigit * 10) + lastDigit
		sum += int64(num)

	}

	log.Printf("Sum: %d", sum)
}

func TextToByte(text string) int {
	switch text {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0 // Return 0 if the input is not "one" to "nine"
	}
}

func FindFirstOccurance(haystack string, needles ...string) (occured_str string, idx int, exists bool) {
	// mux as 1 sets it as ascending
	return FindOccurance(haystack, 1, strings.Index, needles...)
}

func FindLastOccurance(haystack string, needles ...string) (occured_str string, idx int, exists bool) {
	// mux as -1 sets it as descending
	return FindOccurance(haystack, -1, strings.LastIndex, needles...)
}

func FindOccurance(haystack string, mux int, indexFinder func(s string, substr string) int, needles ...string) (occured_str string, idx int, exists bool) {
	needleToFirstIndexMap := make(map[string]int, len(needles)/2)
	for _, needle := range needles {
		if i := indexFinder(haystack, needle); i != -1 {
			needleToFirstIndexMap[needle] = i
		}
	}

	needles = Filter(needles, func(a string) bool {
		_, exists := needleToFirstIndexMap[a]
		return exists
	})

	switch len(needles) {
	case 0:
		return occured_str, -1, false
	case 1:
		break // anyways first element is the one to return if its the only one in slice, no need sorting
	default:
		slices.SortFunc(needles, func(a string, b string) int {
			aidx := needleToFirstIndexMap[a]
			bidx := needleToFirstIndexMap[b]
			return (aidx - bidx) * mux // if mux is -1, it descends, if 1 it ascends
		})
	}
	return needles[0], needleToFirstIndexMap[needles[0]], true
}

func Filter(data []string, closure func(string) bool) []string {
	var newSlice []string

	for _, v := range data {
		if closure(v) {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}

// func Filter[T any](data []T, closure func(a T) bool) []T {
// 	newSlice := make([]T, 0)

// 	for _, v := range data {
// 		if closure(v) {
// 			newSlice = append(newSlice, v)
// 		}
// 	}

// 	return newSlice
// }
