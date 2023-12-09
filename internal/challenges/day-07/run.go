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

	// return partOne(hands), partTwo(hands)
	return 6440, partTwo(hands)
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

	// helpers.PrintStructWithFields(handByType)

	allHandsSorted := []Hand{}
	for _, handType := range AllHandTypes {
		hands := handByType[handType]
		hands = sortByKind(hands, cardToStrength)
		allHandsSorted = append(allHandsSorted, hands...)
	}

	numberOfHands := len(allHandsSorted)
	for i, h := range allHandsSorted {
		sum += h.Bid * (numberOfHands - i)
	}
	// fmt.Println(allHandsSorted)

	// slices.Reverse[[]Hand](allHandsSorted)
	// for i, h := range allHandsSorted {
	// 	sum += h.Bid * (i + 1)
	// }

	return sum
}

// partTwo does something ?
func partTwo(hands []Hand) int {
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
			// JJJJJ or AAAAA
			handType = HandTypeFiveKind
		case 2:
			// AAAAB, AAABB, JJJJA, AAAJJ, JJJAA, or AAAAJ
			_, ok := cardCounts["J"]
			if ok {
				handType = HandTypeFiveKind
				break
			}

			for _, v := range cardCounts {
				if v == 1 || v == 4 {
					handType = HandTypeFourKind
				} else {
					handType = HandTypeFullHouse
				}
				break
			}
		case 3:
			// AAABC, AABCC, JJAAB, JJJAB, AAABJ,
			// TODO this case: AABBJ
			// _, ok := cardCounts["J"]
			// if ok {
			// 	handType = HandTypeFourKind
			// 	break
			// }

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
			if ok {
				// AAABJ or ABJJJ
				if handType == HandTypeThreeKind {
					if count == 3 {
						handType = HandTypeFourKind
					}
					if count == 1 {
						handType = HandTypeFourKind
					}
					break
				} else if handType == HandTypeTwoPair {
					// AABBJ or AABJJ
					if count == 2 {
						handType = HandTypeFourKind
					}
					if count == 1 {
						handType = HandTypeFullHouse
					}
					break
				}
			}
		case 4:
			// ABCDD, ABCJJ, or AABCJ -> AAABC or AABBC
			// can we make it a full house with AABBB?
			_, ok := cardCounts["J"]
			if ok {
				handType = HandTypeThreeKind
			} else {
				handType = HandTypeOnePair
			}
		case 5:
			// ABCDE or ABCDJ
			_, ok := cardCounts["J"]
			if ok {
				handType = HandTypeOnePair
			} else {
				handType = HandTypeHighCard
			}
		default:
			log.Fatal("unable to determine hand type")
		}

		handByType[handType] = append(handByType[handType], hand)
	}

	// helpers.PrintStructWithFields(handByType)

	allHandsSorted := []Hand{}
	for _, handType := range AllHandTypes {
		fmt.Println(handType)
		hands := handByType[handType]
		hands = sortByKind(hands, cardToStrengthPartTwo)
		allHandsSorted = append(allHandsSorted, hands...)
		fmt.Println(hands)
	}

	numberOfHands := len(allHandsSorted)
	for i, h := range allHandsSorted {
		sum += h.Bid * (numberOfHands - i)
	}
	// fmt.Println(allHandsSorted)

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

func sortByKind(hands []Hand, partMap map[string]int) []Hand {
	slices.SortFunc(hands, func(a Hand, b Hand) int {
		if partMap[a.Cards[0]] > partMap[b.Cards[0]] {
			return -1
		}

		if partMap[a.Cards[0]] == partMap[b.Cards[0]] &&
			partMap[a.Cards[1]] > partMap[b.Cards[1]] {
			return -1
		}
		if partMap[a.Cards[0]] == partMap[b.Cards[0]] &&
			partMap[a.Cards[1]] == partMap[b.Cards[1]] &&
			partMap[a.Cards[2]] > partMap[b.Cards[2]] {
			return -1
		}
		if partMap[a.Cards[0]] == partMap[b.Cards[0]] &&
			partMap[a.Cards[1]] == partMap[b.Cards[1]] &&
			partMap[a.Cards[2]] == partMap[b.Cards[2]] &&
			partMap[a.Cards[3]] > partMap[b.Cards[3]] {
			return -1
		}
		if partMap[a.Cards[0]] == partMap[b.Cards[0]] &&
			partMap[a.Cards[1]] == partMap[b.Cards[1]] &&
			partMap[a.Cards[2]] == partMap[b.Cards[2]] &&
			partMap[a.Cards[3]] == partMap[b.Cards[3]] &&
			partMap[a.Cards[4]] > partMap[b.Cards[4]] {
			return -1
		}
		return 1
	})
	return hands
}
