package day07

// Hand contains details for a hand
type Hand struct {
	Cards []string
	Bid   int
}

// cardToStrength maps a card to its relative strength
var cardToStrength = map[string]int{
	"2": 1,
	"3": 2,
	"4": 3,
	"5": 4,
	"6": 5,
	"7": 6,
	"8": 7,
	"9": 8,
	"T": 9,
	"J": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

// cardToStrength maps a card to its relative strength
var cardToStrengthPartTwo = map[string]int{
	"J": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"T": 10,
	"Q": 11,
	"K": 12,
	"A": 13,
}

type HandType string

// Ranked from strongest to weakest
var (
	// HandTypeFiveKind is where all five cards have the same label: AAAAA
	HandTypeFiveKind HandType = "Five of a kind"

	// HandTypeFourKind is where four cards have the same label and one card has a different label: AA8AA
	HandTypeFourKind HandType = "Four of a kind"

	// HandTypeFullHouse is where three cards have the same label, and the remaining two cards share a different label: 23332
	HandTypeFullHouse HandType = "Full house"

	// HandTypeThreeKind is where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
	HandTypeThreeKind HandType = "Three of a kind"

	// HandTypeTwoPair is where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
	HandTypeTwoPair HandType = "Two pair"

	// HandTypeTwoPair is where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
	HandTypeOnePair HandType = "One pair"

	// HandTypeHighCard is where all cards' labels are distinct: 23456
	HandTypeHighCard HandType = "High card"
)

var AllHandTypes = []HandType{HandTypeFiveKind, HandTypeFourKind, HandTypeFullHouse, HandTypeThreeKind, HandTypeTwoPair, HandTypeOnePair, HandTypeHighCard}

// HandByType is an alias to map the ID of a card to the number of copies
type HandByType map[HandType][]Hand
