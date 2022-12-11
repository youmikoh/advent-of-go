package main

import (
	"advent-of-go/utils"
	"fmt"
	"strings"
)

func Hand(hand string) int {
	switch {
	case hand == "A" || hand == "X":
		return 1
	case hand == "B" || hand == "Y":
		return 2
	case hand == "C" || hand == "Z":
		return 3
	default:
		panic("sad")
	}
}

func RoundScore(theirs int, mine int) int {
	switch {
	case mine == theirs:
		return 3
	case mine == theirs%3+1:
		return 6
	default:
		return 0
	}
}

func MyHand(theirs int, outcome int) int {
	switch {
	case outcome == 3:
		return theirs
	case outcome == 6:
		return theirs%3 + 1
	default:
		return ((theirs + 1) % 3) + 1
	}
}

func Part1(inputs <-chan []string) int {
	var totalScore int
	for input := range inputs {
		theirs, mine := Hand(input[0]), Hand(input[1])
		totalScore += RoundScore(theirs, mine) + mine
	}
	return totalScore
}

func Part2(inputs <-chan []string) int {
	var totalScore int
	for input := range inputs {
		theirs, outcome := Hand(input[0]), (Hand(input[1])-1)*3
		totalScore += outcome + MyHand(theirs, outcome)
	}
	return totalScore
}

func main() {
	inputs := utils.StreamInput(
		func(line string) []string {
			return strings.Split(line, " ")
		},
	)
	fmt.Printf("Part 1: %v\n", Part1(inputs))
	fmt.Printf("Part 2: %v\n", Part2(inputs))
}
