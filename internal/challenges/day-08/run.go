package day08

import (
	"fmt"
	"strings"

	"github.com/humzamo/advent-of-code-2023/internal/helpers"
)

const day = "day-08"

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
	instructions, nodes := parseList(list)
	fmt.Println(instructions, nodes)
	return 0, partTwo(instructions, nodes)
}

func partOne(instructions []string, nodes Nodes) int {
	count := 0
	reachedZ := false

	currentNode := "AAA"
	for i := 0; i < len(instructions); i++ {
		count++
		if reachedZ {
			break
		}

		fmt.Println("current:", currentNode)
		fmt.Println(instructions[i])
		fmt.Println(nodes[currentNode])

		if instructions[i] == "L" {
			currentNode = nodes[currentNode][0]
		}

		if instructions[i] == "R" {
			currentNode = nodes[currentNode][1]
		}
		fmt.Println("new:", currentNode)

		if currentNode == "ZZZ" {
			reachedZ = true
			break
		}

		if i == len(instructions)-1 && !reachedZ {
			i = -1
		}
	}

	return count
}

func partTwo(instructions []string, nodes Nodes) int {
	count := 0
	reachedZ := false
	currentNodes := []string{}

	for n := range nodes {
		if n[2] == 'A' {
			currentNodes = append(currentNodes, n)
		}
	}

	for i := 0; i < len(instructions); i++ {
		count++
		if reachedZ {
			break
		}

		fmt.Println("current:", currentNodes)
		// fmt.Println("current:", currentNode)
		fmt.Println(instructions[i])
		// fmt.Println(nodes[currentNode])

		if instructions[i] == "L" {

			for i, n := range currentNodes {
				currentNodes[i] = nodes[n][0]
			}

		}

		if instructions[i] == "R" {
			for i, n := range currentNodes {
				currentNodes[i] = nodes[n][1]
			}
		}
		fmt.Println("new:", currentNodes)

		endLetters := ""
		for _, n := range currentNodes {
			endLetters += string(n[2])
			// endLetter = append(endLetter, n[2])
			// if n[2] != 'Z' {
			// 	reachedZ = false
			// 	break
			// }
			// reachedZ = true
		}
		fmt.Println(endLetters)
		if endLetters == "ZZZZZZ" {
			reachedZ = true
			break
		}

		if reachedZ {
			break
		}

		if i == len(instructions)-1 && !reachedZ {
			i = -1
		}
	}

	return count
}

// parseList parses the string list into a slice of hands
func parseList(list []string) ([]string, Nodes) {
	nodes := Nodes{}
	instructions := strings.Split(list[0], "")

	for i := 2; i < len(list); i++ {
		splitNode := strings.Split(list[i], " = ")

		splitElements := strings.Split(splitNode[1], ", ")
		leftElement := splitElements[0][1:]
		rightElement := splitElements[1][0:3]
		nodes[splitNode[0]] = []string{leftElement, rightElement}
	}

	return instructions, nodes
}
