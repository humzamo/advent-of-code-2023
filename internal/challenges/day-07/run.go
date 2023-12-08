package day07

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/humzamo/advent-of-code-2023/internal/helpers"
)

const day = "day-07"

// Run runs the input file to generate the answers for Day 7
func Run() {
	fmt.Printf("Generating solutions for %s...\n", day)

	partOneAns, partTwoAns := generateAnswers(fmt.Sprintf("./internal/challenges/%s/input.txt", day))

	fmt.Println("The answer to part one is:", partOneAns)
	fmt.Println("The answer to part two is:", partTwoAns)
}

// generateAnswers generates the answers for both parts
func generateAnswers(inputFile string) (int, int) {
	list := helpers.LoadStringList(inputFile)
	hands := parseList(list)

	return partOne(hands), partTwo(hands)
}

// partOne sums the rank of each card mulitplied by its rank
func partOne(hands []Hand) int {
	sum := 0
	handByType := HandByType{}

	for _, hand := range hands {
		cardCounts := make(map[string]int)
		for _, card := range hand.Cards {
			cardCounts[card]++
		}

		var handType HandType
		uniqueCardCount := len(cardCounts)

		switch uniqueCardCount {
		case 1:
			handType = HandTypeFiveKind
		case 2:
			for _, v := range cardCounts {
				if v == 1 || v == 4 {
					handType = HandTypeFourKind
				} else {
					handType = HandTypeFullHouse
				}
				break
			}
		case 3:
			found := false
			for _, v := range cardCounts {
				if v == 3 {
					handType = HandTypeThreeKind
					found = true
					break
				}
			}
			if !found {
				handType = HandTypeTwoPair
			}
		case 4:
			handType = HandTypeOnePair
		case 5:
			handType = HandTypeHighCard
		default:
			log.Fatal("unable to determine hand type")
		}

		handByType[handType] = append(handByType[handType], hand)
	}

	helpers.PrintStructWithFields(handByType)

	allHandsSorted := []Hand{}
	for _, handType := range AllHandTypes {
		hands := handByType[handType]
		hands = sortByKind(hands)
		allHandsSorted = append(allHandsSorted, hands...)
	}

	numberOfHands := len(allHandsSorted)
	for i, h := range allHandsSorted {
		sum += h.Bid * (numberOfHands - i)
	}
	fmt.Println(allHandsSorted)

	// slices.Reverse[[]Hand](allHandsSorted)
	// for i, h := range allHandsSorted {
	// 	sum += h.Bid * (i + 1)
	// }

	return sum
}

// partTwo does something ?
func partTwo(hands []Hand) int {
	sum := 0
	// TODO
	return sum
}

// parseList parses the string list into a slice of hands
func parseList(list []string) []Hand {
	var hands []Hand

	for _, row := range list {
		splitHand := strings.Fields(row)
		cards := strings.Split(splitHand[0], "")

		bidString := splitHand[1]
		bid, err := strconv.Atoi(bidString)
		if err != nil {
			log.Fatal(err)
		}

		hand := Hand{
			Cards: cards,
			Bid:   bid,
		}

		hands = append(hands, hand)
	}

	return hands
}

func sortByKind(hands []Hand) []Hand {
	slices.SortFunc(hands, func(a Hand, b Hand) int {
		if cardToStrength[a.Cards[0]] > cardToStrength[b.Cards[0]] {
			return -1
		}

		if cardToStrength[a.Cards[0]] == cardToStrength[b.Cards[0]] &&
			cardToStrength[a.Cards[1]] > cardToStrength[b.Cards[1]] {
			return -1
		}
		if cardToStrength[a.Cards[0]] == cardToStrength[b.Cards[0]] &&
			cardToStrength[a.Cards[1]] == cardToStrength[b.Cards[1]] &&
			cardToStrength[a.Cards[2]] > cardToStrength[b.Cards[2]] {
			return -1
		}
		if cardToStrength[a.Cards[0]] == cardToStrength[b.Cards[0]] &&
			cardToStrength[a.Cards[1]] == cardToStrength[b.Cards[1]] &&
			cardToStrength[a.Cards[2]] == cardToStrength[b.Cards[2]] &&
			cardToStrength[a.Cards[3]] > cardToStrength[b.Cards[3]] {
			return -1
		}
		if cardToStrength[a.Cards[0]] == cardToStrength[b.Cards[0]] &&
			cardToStrength[a.Cards[1]] == cardToStrength[b.Cards[1]] &&
			cardToStrength[a.Cards[2]] == cardToStrength[b.Cards[2]] &&
			cardToStrength[a.Cards[3]] == cardToStrength[b.Cards[3]] &&
			cardToStrength[a.Cards[4]] > cardToStrength[b.Cards[4]] {
			return -1
		}
		return 1
	})
	return hands
}
