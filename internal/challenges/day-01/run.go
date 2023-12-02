package day01

import (
	"fmt"
	"log"
	"strconv"
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

var wordToDigit = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

// calculateSum calculated the sum of all the two digit numbers in the list
// for part one, we only consider digit in the string
// for part two, we also consider words that are valid numbers
func calculateSum(list []string, partTwo bool) int {
	sum := 0
	for _, row := range list {
		var firstDigit, secondDigit string

		// user pointer i to go forward in the string until we reach a valid digit
		i := 0
		for i < len(row) {
			if partTwo {
				foundWord := false
				for word, digit := range wordToDigit {
					// check if the string up to i contains a number word
					if strings.Contains(row[0:i], word) {
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
				firstDigit = string(row[i])
				i++
				break
			}
			i++
		}

		// user pointer j to go backwards in the string until we reach a valid digit
		// or if the pointers meet
		j := len(row) - 1
		for i <= j {
			// check if the char at j is a digit
			if unicode.IsDigit(rune(row[j])) {
				secondDigit = string(row[j])
				break
			}

			if partTwo {
				foundWord := false
				// check if the string from j to the end contains a number word
				for word, digit := range wordToDigit {
					if strings.Contains(row[j:], word) {
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

		val, err := strconv.Atoi(firstDigit + secondDigit)
		if err != nil {
			log.Fatal(err)
		}

		sum += val
	}

	return sum
}
