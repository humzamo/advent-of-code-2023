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

	return calculateAnswer(hands, false), calculateAnswer(hands, true)
}

// calculateAnswer sums the rank of each card mulitplied by its rank
func calculateAnswer(hands []Hand, partTwo bool) int {
	sum := 0
	handByType := HandByType{}
	partMap := cardToStrengthPartOne

	if partTwo {
		partMap = cardToStrengthPartTwo
	}

	for _, hand := range hands {
		cardCounts := make(map[string]int)
		for _, card := range hand.Cards {
			cardCounts[card]++
		}

		var handType HandType
		uniqueCardCount := len(cardCounts)

		switch uniqueCardCount {
		case 1:
			// JJJJJ or AAAAA
			handType = HandTypeFiveKind
		case 2:
			// AAAAB, AAABB, JJJJA, AAAJJ, JJJAA, or AAAAJ
			_, ok := cardCounts["J"]
			if ok && partTwo {
				handType = HandTypeFiveKind
				break
			}

			// if there is 1 or 4 of any card, the type must be four of a kind
			for _, v := range cardCounts {
				if v == 1 || v == 4 {
					handType = HandTypeFourKind
				} else {
					handType = HandTypeFullHouse
				}
				break
			}
		case 3:
			// AAABC, AABCC, JJAAB, JJJAB, AAABJ, AABBJ
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

			count, ok := cardCounts["J"]
			if ok && partTwo {
				// AAABJ or ABJJJ
				if handType == HandTypeThreeKind {
					handType = HandTypeFourKind
					break
				}

				// AABBJ or AABJJ
				if handType == HandTypeTwoPair {
					if count == 1 {
						handType = HandTypeFullHouse
					}
					if count == 2 {
						handType = HandTypeFourKind
					}
				}
			}
		case 4:
			// ABCDD, ABCJJ, or AABCJ
			_, ok := cardCounts["J"]
			if ok && partTwo {
				handType = HandTypeThreeKind
				break
			}
			handType = HandTypeOnePair
		case 5:
			// ABCDE or ABCDJ
			_, ok := cardCounts["J"]
			if ok && partTwo {
				handType = HandTypeOnePair
				break
			}
			handType = HandTypeHighCard
		default:
			log.Fatal("unable to determine hand type")
		}

		handByType[handType] = append(handByType[handType], hand)
	}

	// sort the cards from strongest to weakest based on the hand type
	allHandsSorted := []Hand{}
	for _, handType := range AllHandTypes {
		hands := sortByKind(handByType[handType], partMap)
		allHandsSorted = append(allHandsSorted, hands...)
	}

	numberOfHands := len(allHandsSorted)
	for i, h := range allHandsSorted {
		sum += h.Bid * (numberOfHands - i)
	}

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

		hands = append(hands, Hand{
			Cards: cards,
			Bid:   bid,
		})
	}

	return hands
}

// sortByKind ranks a slice of hands by the strength distinct card
func sortByKind(hands []Hand, partMap map[string]int) []Hand {
	slices.SortFunc(hands, func(a Hand, b Hand) int {
		if partMap[a.Cards[0]] > partMap[b.Cards[0]] {
			return -1
		}

		if a.Cards[0] == b.Cards[0] &&
			partMap[a.Cards[1]] > partMap[b.Cards[1]] {
			return -1
		}
		if a.Cards[0] == b.Cards[0] &&
			a.Cards[1] == b.Cards[1] &&
			partMap[a.Cards[2]] > partMap[b.Cards[2]] {
			return -1
		}
		if a.Cards[0] == b.Cards[0] &&
			a.Cards[1] == b.Cards[1] &&
			a.Cards[2] == b.Cards[2] &&
			partMap[a.Cards[3]] > partMap[b.Cards[3]] {
			return -1
		}
		if a.Cards[0] == b.Cards[0] &&
			a.Cards[1] == b.Cards[1] &&
			a.Cards[2] == b.Cards[2] &&
			a.Cards[3] == b.Cards[3] &&
			partMap[a.Cards[4]] > partMap[b.Cards[4]] {
			return -1
		}
		return 1
	})
	return hands
}
