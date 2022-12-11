package main

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part1(inputs <-chan [][]int) int {
	count := 0
	for input := range inputs {
		a, b := input[0], input[1]
		if (a[0] <= b[0] && a[1] >= b[1]) || (b[0] <= a[0] && b[1] >= a[1]) {
			count += 1
		}
	}
	return count
}

func Part2(inputs <-chan [][]int) int {
	count := 0
	for input := range inputs {
		a, b := input[0], input[1]
		if a[0] <= b[0] && a[1] >= b[0] || a[0] >= b[0] && b[1] >= a[0] {
			count += 1
		}
	}
	return count
}

func main() {
	inputs := utils.StreamInput(
		func(line string) [][]int {
			pair := strings.Split(line, ",")
			var ints [][]int
			for _, p := range pair {
				s := strings.Split(p, "-")
				a, _ := strconv.Atoi(s[0])
				b, _ := strconv.Atoi(s[1])
				ints = append(ints, []int{a, b})
			}
			return ints
		},
	)
	fmt.Printf("count %v\n", Part1(inputs))
	fmt.Printf("count %v\n", Part2(inputs))
}
