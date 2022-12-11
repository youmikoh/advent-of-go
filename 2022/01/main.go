package main

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Part1(inputs <-chan int) int {
	max, sum := 0, 0
	for input := range inputs {
		if input > 0 {
			sum += input
		} else {
			if max < sum {
				max = sum
			}
			sum = 0
		}
	}
	return max
}

func Part2(inputs <-chan int) int {
	top, sum := [3]int{}, 0
	for input := range inputs {
		if input > 0 {
			sum += input
		} else {
			switch {
			case top[0] < sum:
				top[0], top[1], top[2] = sum, top[0], top[1]
			case top[1] < sum:
				top[1], top[2] = sum, top[1]
			case top[2] < sum:
				top[2] = sum
			}
			sum = 0
		}
	}
	return top[0] + top[1] + top[2]
}

func main() {
	inputs := utils.StreamInput(
		func(line string) int {
			if line == "" {
				return 0
			}
			n, _ := strconv.Atoi(line)
			return n
		},
	)
	fmt.Printf("Part 1: %v\n", Part1(inputs))
	fmt.Printf("Part 2: %v\n", Part2(inputs))
}
