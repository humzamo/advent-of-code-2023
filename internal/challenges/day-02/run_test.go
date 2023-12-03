package day02

import (
	"fmt"
	"testing"

	"github.com/humzamo/advent-of-code-2023/internal/helpers"
)

type testCase struct {
	inputFile      string
	partNumber     int
	expectedAnswer int
}

func TestDay02(t *testing.T) {
	cases := []testCase{
		{"input_test.txt", 1, 8},
		{"input_test.txt", 2, 2286},
		{"input.txt", 1, 2447},
		{"input.txt", 2, 56322},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("testing part %d with input file %s", tc.partNumber, tc.inputFile), func(t *testing.T) {
			list := helpers.LoadStringList(tc.inputFile)
			games := parseList(list)

			var answer int
			switch tc.partNumber {
			case 1:
				answer = partOne(games)
			case 2:
				answer = partTwo(games)
			}

			if answer != tc.expectedAnswer {
				t.Errorf("Expected %v, but got %v", tc.expectedAnswer, answer)
			}
		})
	}
}
