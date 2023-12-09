package day04

import (
	"fmt"
	"testing"
)

type testCase struct {
	inputFile             string
	expectedPartOneAnswer int
	expectedPartTwoAnswer int
}

func TestDay04(t *testing.T) {
	cases := []testCase{
		{"input_test.txt", 13, 30},
		{"input.txt", 23678, 15455663},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("testing with input file %s", tc.inputFile), func(t *testing.T) {
			actualPartOneAns, actualPartTwoAns := generateAnswers(tc.inputFile)

			if actualPartOneAns != tc.expectedPartOneAnswer {
				t.Errorf("Failed running part one! Expected %v, but got %v", tc.expectedPartOneAnswer, actualPartOneAns)
			}

			if actualPartTwoAns != tc.expectedPartTwoAnswer {
				t.Errorf("Failed running part two! Expected %v, but got %v", tc.expectedPartTwoAnswer, actualPartTwoAns)
			}
		})
	}
}
