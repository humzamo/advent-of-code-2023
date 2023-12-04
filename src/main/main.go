package main

import (
	"fmt"
	"os"
	"strconv"

	day01 "github.com/humzamo/advent-of-code-2023/internal/challenges/day-01"
	day02 "github.com/humzamo/advent-of-code-2023/internal/challenges/day-02"
	day04 "github.com/humzamo/advent-of-code-2023/internal/challenges/day-04"
)

func main() {
	fmt.Print("🎄 Welcome to Advent of Code 2023! 🎄\n\n")

	if len(os.Args) < 2 {
		fmt.Println("Please make sure you enter a day to generate a solution!")
		fmt.Println("Usage: make run day=<day_number>")
		os.Exit(1)
	}

	dayString := os.Args[1]

	dayNumber, err := strconv.Atoi(dayString)
	if err != nil {
		fmt.Printf("invalid day entered! [day=%s]", dayString)
		os.Exit(1)
	}

	if dayNumber > 25 {
		fmt.Println("Sorry there are only 25 days in this challenge!")
		os.Exit(1)
	}

	switch dayNumber {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 4:
		day04.Run()
	default:
		fmt.Printf("Uh oh, there's no solution for day %s yet... 👀\n", dayString)
		os.Exit(1)
	}
}
