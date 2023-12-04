package day04

// Game contains details for a game
type Game struct {
	ID                   int
	WinningNumbers       []string
	ChosenNumbers        []string
	MatchingNumbers      []string
	MatchingNumbersCount int
}
