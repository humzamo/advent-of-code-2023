package day01

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/humzamo/advent-of-code-2023/internal/helpers"
)

func Run() {
	fmt.Println("Generating solutions for day 01...")

	list := helpers.LoadStringList("../../internal/challenges/day-01/input.txt")

	fmt.Println("The answer to part one is:", calculateSum(list, false))
	fmt.Println("The answer to part two is:", calculateSum(list, true))
}

var wordToDigit = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// calculateSum calculated the sum of all the two digit numbers in the list
// for part one, we only consider digit in the string
// for part two, we also consider words that are valid numbers
func calculateSum(list []string, partTwo bool) int {
	sum := 0
	for _, row := range list {
		var firstDigit, secondDigit int

		// use pointer i to go forward in the string until we reach a valid digit
		i := 0
		for i < len(row) {
			if partTwo {
				foundWord := false
				for word, digit := range wordToDigit {
					// check if the string from the start index to i contains a number word
					if strings.Contains(row[startIndex(i):i], word) {
						firstDigit = digit
						foundWord = true
						break
					}
				}
				if foundWord {
					break
				}
			}

			// check if the char at i is a digit
			if unicode.IsDigit(rune(row[i])) {
				firstDigit = byteToDigit(row[i])
				i++
				break
			}
			i++
		}

		// use pointer j to go backwards in the string until we reach a valid digit
		// or if the pointers meet
		j := len(row) - 1
		for i <= j {
			// check if the char at j is a digit
			if unicode.IsDigit(rune(row[j])) {
				secondDigit = byteToDigit(row[j])
				break
			}

			if partTwo {
				foundWord := false
				// check if the string from j to the end index contains a number word
				for word, digit := range wordToDigit {
					if strings.Contains(row[j:endIndex(j, len(row))], word) {
						secondDigit = digit
						foundWord = true
						break
					}
				}
				if foundWord {
					break
				}
			}
			j--
		}

		// if the pointers meet, there was no valid digit when traversing backwards,
		// so the second digit should be the same as the first digit
		if i == j+1 {
			secondDigit = firstDigit
		}

		sum += 10*firstDigit + secondDigit
	}

	return sum
}

// startIndex returns an index to start position of the search
// since the longest word is 5 chars, we only need to check range of i to five indexes before i
func startIndex(i int) int {
	if i < 5 {
		return 0
	}
	return i - 5
}

// endIndex returns an index to end position of the search
// since the longest word is 5 chars, we only need to check range of j to five indexes after j
func endIndex(j, stringLen int) int {
	if j-5 < stringLen {
		return stringLen
	}

	return j - 5
}

// byteToInt converts a byte to a digit
func byteToDigit(b byte) int {
	return int(b - 48)
}
