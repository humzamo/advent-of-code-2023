package day04

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/humzamo/advent-of-code-2023/internal/helpers"
)

func Run() {
	fmt.Println("Generating solutions for day 04...")

	list := helpers.LoadStringList("./internal/challenges/day-04/input.txt")
	games := parseList(list)

	fmt.Println("The answer to part one is:", partOne(games))
	// fmt.Println("The answer to part two is:", partTwo(games))
}

// partOne sums the points from all the games
func partOne(games []Game) int {
	sum := 0
	for _, game := range games {
		sum += gamesToPoints[game.MatchingNumbersCount]
	}
	return sum
}

// partTwo sums the powers of all the games
func partTwo(games []Game) int {
	sum := 0
	for _, game := range games {
		// TODO
		sum += game.ID
	}
	return sum
}

var gamesToPoints = map[int]int{
	1:  1,
	2:  2,
	3:  4,
	4:  8,
	5:  16,
	6:  32,
	7:  64,
	8:  128,
	9:  256,
	10: 512,
}

// parseList parses the string list into a slice of completed games
func parseList(list []string) []Game {
	var games []Game

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

		game := Game{
			ID:                   id,
			WinningNumbers:       winningNumbers,
			ChosenNumbers:        chosenNumbers,
			MatchingNumbers:      matchingNumbers,
			MatchingNumbersCount: len(matchingNumbers),
		}

		games = append(games, game)
	}
	return games
}
