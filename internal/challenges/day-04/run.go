package day04

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/humzamo/advent-of-code-2023/internal/helpers"
)

// Run runs the input file to generate the answers for Day 4
func Run() {
	fmt.Println("Generating solutions for day 04...")

	partOneAns, partTwoAns := generateAnswers("./internal/challenges/day-04/input.txt")

	fmt.Println("The answer to part one is:", partOneAns)
	fmt.Println("The answer to part two is:", partTwoAns)
}

// generateAnswers generates the answers for both parts
func generateAnswers(inputFile string) (int, int) {
	list := helpers.LoadStringList(inputFile)
	cards, copies := parseList(list)

	return partOne(cards), partTwo(copies)
}

// partOne sums the points from all the cards
func partOne(cards []Card) int {
	sum := 0
	for _, c := range cards {
		sum += cardToPoints[c.MatchingNumbersCount]
	}
	return sum
}

// partTwo sums the copies of all the cards
func partTwo(copies Copies) int {
	sum := 0
	for _, count := range copies {
		sum += count
	}
	return sum
}

// parseList parses the string list into a slice of cards and copies information
func parseList(list []string) ([]Card, Copies) {
	var cards []Card
	var copies = Copies{}

	for _, row := range list {
		splitByID := strings.Split(row, ":")
		idString := strings.Fields(splitByID[0])[1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			log.Fatal(err)
		}

		splitByNumbers := strings.Split(splitByID[1], "|")

		winningNumbersString := splitByNumbers[0]
		chosenNumbersString := splitByNumbers[1]

		winningNumbers := strings.Fields(winningNumbersString)
		chosenNumbers := strings.Fields(chosenNumbersString)

		matchingNumbers := helpers.Intersection(winningNumbers, chosenNumbers)

		card := Card{
			ID:                   id,
			WinningNumbers:       winningNumbers,
			ChosenNumbers:        chosenNumbers,
			MatchingNumbers:      matchingNumbers,
			MatchingNumbersCount: len(matchingNumbers),
		}

		// this is the first time analysing this new card so add a copy to the map
		copies[id]++

		// for every copy of the card, increase the copies of the corresponding cards
		for i := 0; i < copies[id]; i++ {
			for j := 1; j <= len(matchingNumbers); j++ {
				copies[id+j]++
			}
		}

		cards = append(cards, card)
	}

	return cards, copies
}
