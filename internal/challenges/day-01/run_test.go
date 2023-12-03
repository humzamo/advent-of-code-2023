package day01

import (
	"fmt"
	"testing"

	"github.com/humzamo/advent-of-code-2023/internal/helpers"
)

type testCase struct {
	inputFile      string
	partNumber     int
	partTwo        bool
	expectedAnswer int
}

func TestDay01(t *testing.T) {
	cases := []testCase{
		{"input_test_part_1.txt", 1, false, 142},
		{"input_test_part_2.txt", 2, true, 281},
		{"input.txt", 1, false, 56049},
		{"input.txt", 2, true, 54530},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("testing part %d with input file %s", tc.partNumber, tc.inputFile), func(t *testing.T) {
			list := helpers.LoadStringList(tc.inputFile)
			answer := calculateSum(list, tc.partTwo)

			if answer != tc.expectedAnswer {
				t.Errorf("Expected %v, but got %v", tc.expectedAnswer, answer)
			}
		})
	}
}
