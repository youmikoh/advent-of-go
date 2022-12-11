package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BuildStacks(s *bufio.Scanner) [][]rune {
	var crates [][]rune
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		var runes []rune
		for i, r := range line {
			if i%4 == 1 {
				runes = append(runes, r)
			}
		}
		crates = append(crates, runes)
	}

	var stacks [][]rune
	for range crates[0] {
		stacks = append(stacks, []rune{})
	}

	crates = crates[:len(crates)-1]
	for _, crate := range Reverse(crates) {
		for i, c := range crate {
			if c != ' ' {
				stacks[i] = append(stacks[i], c)
			}
		}
	}

	return stacks
}

func StreamMoves(s *bufio.Scanner) <-chan []int {
	moves := make(chan []int)
	go func() {
		defer close(moves)
		for s.Scan() {
			line := s.Text()
			m := strings.Fields(line)
			units, _ := strconv.Atoi(m[1])
			from, _ := strconv.Atoi(m[3])
			to, _ := strconv.Atoi(m[5])
			moves <- []int{units, from - 1, to - 1}
		}
	}()
	return moves
}

func Part1(stacks [][]rune, moves <-chan []int) string {
	for m := range moves {
		from, to := m[1], m[2]
		tail := len(stacks[from]) - m[0]
		chunk := stacks[from][tail:]

		stacks[from] = stacks[from][:tail]
		stacks[to] = append(stacks[to], Reverse(chunk)...)
	}
	var lasts string
	for _, stack := range stacks {
		last := stack[len(stack)-1]
		lasts += string(last)
	}
	return lasts
}

func Part2(stacks [][]rune, moves <-chan []int) string {
	for m := range moves {
		from, to := m[1], m[2]
		tail := len(stacks[from]) - m[0]
		chunk := stacks[from][tail:]

		stacks[from] = stacks[from][:tail]
		stacks[to] = append(stacks[to], chunk...)
	}
	var lasts string
	for _, stack := range stacks {
		last := stack[len(stack)-1]
		lasts += string(last)
	}
	return lasts
}

func Reverse[T any](slice []T) []T {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func main() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)
	stacks := BuildStacks(scanner)
	moves := StreamMoves(scanner)

	fmt.Printf("Part 1: %v\n", Part1(stacks, moves))
	fmt.Printf("Part 2: %v\n", Part2(stacks, moves))

}
