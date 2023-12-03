package day02

// Game contains details for a game
type Game struct {
	ID       int
	Rounds   []Round
	Possible bool
	MinRed   int
	MinGreen int
	MinBlue  int
	Power    int
}

// NewGame returns an empty new game which is assumed possible
func NewGame() Game {
	return Game{Possible: true}
}

// Round contains the number of cubes of each colour
type Round struct {
	Red   int
	Blue  int
	Green int
}

// NewRound returns a round with completed entries
func NewRound(red, green, blue int) Round {
	return Round{
		Red:   red,
		Green: green,
		Blue:  blue,
	}
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
