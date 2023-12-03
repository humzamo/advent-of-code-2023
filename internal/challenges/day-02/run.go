package day02

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/humzamo/advent-of-code-2023/internal/helpers"
)

func Run() {
	fmt.Println("Generating solutions for day 02...")

	list := helpers.LoadStringList("../../internal/challenges/day-02/input.txt")
	games := parseList(list)

	fmt.Println("The answer to part one is:", partOne(games))
	fmt.Println("The answer to part two is:", partTwo(games))
}

// partOne sums the IDs of all the games which are possible
func partOne(games []Game) int {
	sum := 0
	for _, game := range games {
		if game.Possible {
			sum += game.ID
		}
	}
	return sum
}

// partTwo sums the powers of all the games
func partTwo(games []Game) int {
	sum := 0
	for _, game := range games {
		sum += game.Power
	}
	return sum
}

// parseList parses the string list into a slice of completed games
func parseList(list []string) []Game {
	var games []Game

	digitsRegex := "([0-9]+)"

	gameIDRegex, err := regexp.Compile("Game " + digitsRegex)
	if err != nil {
		log.Fatal(err)
	}

	greenRegex, err := regexp.Compile(digitsRegex + " " + string(ColourGreen))
	if err != nil {
		log.Fatal(err)
	}

	redRegex, err := regexp.Compile(digitsRegex + " " + string(ColourRed))
	if err != nil {
		log.Fatal(err)
	}

	blueRegex, err := regexp.Compile(digitsRegex + " " + string(ColourBlue))
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range list {
		game := NewGame()

		// the ID corresponds to the item number of the list but this
		// is a useful step in case the input didn't follow this logic
		idMatch := gameIDRegex.FindAllStringSubmatch(row, -1)
		id, err := strconv.Atoi(idMatch[0][1])
		if err != nil {
			log.Fatal(err)
		}
		game.ID = id

		// all rounds of a game are split by the `;` char in the string
		roundStrings := strings.Split(row, ";")

		var rounds []Round

		for _, r := range roundStrings {
			var blue, red, green int

			blueMatch := blueRegex.FindAllStringSubmatch(r, -1)
			if blueMatch != nil {
				blue, err = strconv.Atoi(blueMatch[0][1])
				if err != nil {
					log.Fatal(err)
				}
			}

			redMatch := redRegex.FindAllStringSubmatch(r, -1)
			if redMatch != nil {
				red, err = strconv.Atoi(redMatch[0][1])
				if err != nil {
					log.Fatal(err)
				}
			}

			greenMatch := greenRegex.FindAllStringSubmatch(r, -1)
			if greenMatch != nil {
				green, err = strconv.Atoi(greenMatch[0][1])
				if err != nil {
					log.Fatal(err)
				}
			}

			round := NewRound(red, green, blue)
			game.checkPossible(round)
			game.setMinValues(round)
			rounds = append(rounds, round)
		}
		game.setPower()
		game.Rounds = rounds

		games = append(games, game)
	}
	return games
}

// checkPossible verifies if each round has at least the minimum items
// for each colour
func (g *Game) checkPossible(r Round) {
	if !g.Possible {
		// already verified that the game is impossible
		return
	}
	g.Possible = r.Blue <= MaxColours[ColourBlue] &&
		r.Red <= MaxColours[ColourRed] &&
		r.Green <= MaxColours[ColourGreen]
}

// setMinValues sets the minimum values for the game based on each round
func (g *Game) setMinValues(r Round) {
	if r.Red > 0 && r.Red > g.MinRed {
		g.MinRed = r.Red
	}

	if r.Green > 0 && r.Green > g.MinGreen {
		g.MinGreen = r.Green
	}

	if r.Blue > 0 && r.Blue > g.MinBlue {
		g.MinBlue = r.Blue
	}
}

// setPower sets the power of the game by multiplying all the min values
func (g *Game) setPower() {
	g.Power = g.MinRed * g.MinGreen * g.MinBlue
}
