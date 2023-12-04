package day04

// Card contains details for a card
type Card struct {
	ID                   int
	WinningNumbers       []string
	ChosenNumbers        []string
	MatchingNumbers      []string
	MatchingNumbersCount int
}

// cardToPoints maps the matching numbers to the points value
var cardToPoints = map[int]int{
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

// Copies is an alias to map the ID of a card to the number of copies
type Copies map[int]int
