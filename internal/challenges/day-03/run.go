package day03

import (
	"fmt"
	"log"
	"strconv"
	"unicode"

	"github.com/humzamo/advent-of-code-2023/internal/helpers"
)

func Run() {
	fmt.Println("Generating solutions for day 03...")

	list := helpers.LoadStringList("./internal/challenges/day-03/input_test.txt")
	schematic := parseList(list)
	fmt.Printf("%+v\n", schematic)

	// fmt.Println("The answer to part one is:", partOne(schematic))
}

func partOne(sch Schematic) int {
	sum := 0
	// for partPos, num := range sch.Parts {
	// 	indexes := make([]int, partPos.endIndex-partPos.startIndex)

	// 	ok := sch.Symbols[Position{row: partPos.row, startIndex: partPos.startIndex - 1}]

	// 	for _, index := range indexes {

	// 	}
	// }

	return sum
}

// parseList parses the string list into a Schematic
func parseList(list []string) Schematic {
	sch := NewSchematic()

	for j, row := range list {
		var numStr string
		var indexes []int
		for i, c := range row {
			if unicode.IsDigit(c) {
				numStr += string(c)
				indexes = append(indexes, i)

				if i == len(row)-1 {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						log.Fatal(err)
					}
					sch.Parts[Position{row: j, startIndex: indexes[0], endIndex: indexes[len(indexes)-1]}] = num
					numStr = ""
					indexes = []int{}
				}
				continue
			}
			if c == fullStop {
				if numStr != "" {
					num, err := strconv.Atoi(numStr)
					if err != nil {
						log.Fatal(err)
					}
					sch.Parts[Position{row: j, startIndex: indexes[0], endIndex: indexes[len(indexes)-1]}] = num
					numStr = ""
					indexes = []int{}
				}
				continue
			}

			sch.Symbols[Position{row: j, startIndex: i, endIndex: i}] = true
		}
	}
	return sch
}
