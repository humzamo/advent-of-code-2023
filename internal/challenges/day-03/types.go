package day03

type Schematic struct {
	Symbols map[Position]bool // map only the symbols which are not `.`
	Parts   map[Position]int  // map of positions to the value of the part there
}

func NewSchematic() Schematic {
	return Schematic{
		Symbols: map[Position]bool{},
		Parts:   map[Position]int{},
	}
}

type Position struct {
	row        int
	startIndex int
	endIndex   int
}

const (
	fullStop = '.'
)
