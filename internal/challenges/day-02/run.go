package day02

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/humzamo/advent-of-code-2023/internal/helpers"
)

type Game struct {
	ID       int
	Rounds   []Round
	Possible bool
	Power    int
}

type Round struct {
	Red      int
	Blue     int
	Green    int
	Possible bool
}

var MaxColours = map[Colour]int{
	ColourRed:   12,
	ColourGreen: 13,
	ColourBlue:  14,
}

type Colour string

var (
	ColourRed   Colour = "red"
	ColourGreen Colour = "green"
	ColourBlue  Colour = "blue"
)

func Run() {
	fmt.Println("Generating solutions for day 02...")

	list := helpers.LoadStringList("../../internal/challenges/day-02/input.txt")
	games := parseListPartTwo(list)

	fmt.Println("The answer to part one is:", partOne(games))
	fmt.Println("The answer to part two is:", partTwo(games))
}

var gameIDreg = "Game ([0-9]+):"
var blueCountReg = "([0-9]+) blue"
var greenCountReg = "([0-9]+) green"
var redCountReg = "([0-9]+) red"

func parseListPartOne(list []string) []Game {
	var games []Game

	for _, row := range list {
		r, _ := regexp.Compile(gameIDreg)
		idMatch := r.FindAllStringSubmatch(row, -1)
		id, _ := strconv.Atoi(idMatch[0][1])

		game := Game{
			ID:       id,
			Possible: true,
		}
		roundStrings := strings.Split(row, ";")
		var rounds []Round
		for _, r := range roundStrings {
			round := Round{Possible: true}

			blueReg, _ := regexp.Compile(blueCountReg)
			blueMatch := blueReg.FindAllStringSubmatch(r, -1)
			if blueMatch != nil {
				blue, _ := strconv.Atoi(blueMatch[0][1])
				round.Blue = blue
				if blue > MaxColours[ColourBlue] {
					round.Possible = false
				}

			}

			redReg, _ := regexp.Compile(redCountReg)
			redMatch := redReg.FindAllStringSubmatch(r, -1)
			if redMatch != nil {
				red, _ := strconv.Atoi(redMatch[0][1])
				round.Red = red
				if red > MaxColours[ColourRed] {
					round.Possible = false
				}
			}

			greenReg, _ := regexp.Compile(greenCountReg)
			greenMatch := greenReg.FindAllStringSubmatch(r, -1)
			if greenMatch != nil {
				green, _ := strconv.Atoi(greenMatch[0][1])
				round.Green = green
				if green > MaxColours[ColourGreen] {
					round.Possible = false
				}
			}
			if !round.Possible {
				game.Possible = false
			}
			fmt.Println(round)
			rounds = append(rounds, round)
		}
		game.Rounds = rounds

		games = append(games, game)
	}
	return games
}

func partOne(games []Game) int {
	sum := 0
	for _, game := range games {
		if game.Possible {
			sum += game.ID
		}
	}
	return sum
}

func parseListPartTwo(list []string) []Game {
	var games []Game

	for _, row := range list {
		r, _ := regexp.Compile(gameIDreg)
		idMatch := r.FindAllStringSubmatch(row, -1)
		id, _ := strconv.Atoi(idMatch[0][1])

		game := Game{
			ID:       id,
			Possible: true,
		}
		roundStrings := strings.Split(row, ";")
		var rounds []Round
		var minRed, minBlue, minGreen int
		for _, r := range roundStrings {
			round := Round{Possible: true}

			blueReg, _ := regexp.Compile(blueCountReg)
			blueMatch := blueReg.FindAllStringSubmatch(r, -1)
			if blueMatch != nil {
				blue, _ := strconv.Atoi(blueMatch[0][1])
				round.Blue = blue
				if blue > MaxColours[ColourBlue] {
					round.Possible = false
				}
				if blue > minBlue {
					minBlue = blue
				}
			}

			redReg, _ := regexp.Compile(redCountReg)
			redMatch := redReg.FindAllStringSubmatch(r, -1)
			if redMatch != nil {
				red, _ := strconv.Atoi(redMatch[0][1])
				round.Red = red
				if red > MaxColours[ColourRed] {
					round.Possible = false
				}
				if red > minRed {
					minRed = red
				}
			}

			greenReg, _ := regexp.Compile(greenCountReg)
			greenMatch := greenReg.FindAllStringSubmatch(r, -1)
			if greenMatch != nil {
				green, _ := strconv.Atoi(greenMatch[0][1])
				round.Green = green
				if green > MaxColours[ColourGreen] {
					round.Possible = false
				}
				if green > minGreen {
					minGreen = green
				}
			}
			if !round.Possible {
				game.Possible = false
			}
			rounds = append(rounds, round)
		}
		game.Power = minBlue * minGreen * minRed
		game.Rounds = rounds

		games = append(games, game)
	}
	return games
}

func partTwo(games []Game) int {
	sum := 0
	for _, game := range games {
		sum += game.Power
	}
	return sum
}
