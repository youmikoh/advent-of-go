package main

import (
	"advent-of-go/utils"
	"fmt"
	"sync"
)

func Part1(inputs <-chan []int8) int {
	trees, visibility := [][]int8{}, [][]bool{}
	for input := range inputs {
		trees = append(trees, input)
		visibility = append(visibility, make([]bool, len(input)))
	}
	rows, cols := len(trees), len(trees[0])

	for r := 0; r < rows; r++ {
		var max int8 = -1
		for c := 0; c < cols; c++ {
			if trees[r][c] > max {
				max, visibility[r][c] = trees[r][c], true
			}
		}
		max = -1
		for c := cols - 1; c >= 0; c-- {
			if trees[r][c] > max {
				max, visibility[r][c] = trees[r][c], true
			}
		}
	}
	for c := 0; c < cols; c++ {
		var max int8 = -1
		for r := 0; r < rows; r++ {
			if trees[r][c] > max {
				max, visibility[r][c] = trees[r][c], true
			}
		}
		max = -1
		for r := rows - 1; r >= 0; r-- {
			if trees[r][c] > max {
				max, visibility[r][c] = trees[r][c], true
			}
		}
	}

	count := 0
	for _, rows := range visibility {
		for _, visible := range rows {
			if visible {
				count += 1
			}
		}
	}
	return count
}

func Part2(inputs <-chan []int8) int {
	trees := [][]int8{}
	for input := range inputs {
		trees = append(trees, input)
	}
	rows, cols := len(trees), len(trees[0])

	scores := make(chan int)
	go func() {
		var wg sync.WaitGroup
		defer close(scores)
		for r, row := range trees {
			for c := range row {
				r, c := r, c
				wg.Add(1)
				go func() {
					defer wg.Done()
					scores <- Score(trees, r, c, rows, cols)
				}()
			}
		}
		wg.Wait()
	}()

	max := 0
	for s := range scores {
		if max < s {
			max = s
		}
	}
	return max
}

func Score(trees [][]int8, r, c, R, C int) int {
	down, up, right, left := 0, 0, 0, 0
	for i := r + 1; i < R; i++ {
		down += 1
		if trees[i][c] >= trees[r][c] {
			break
		}
	}
	for i := r - 1; i >= 0; i-- {
		up += 1
		if trees[i][c] >= trees[r][c] {
			break
		}
	}
	for i := c + 1; i < C; i++ {
		right += 1
		if trees[r][i] >= trees[r][c] {
			break
		}
	}
	for i := c - 1; i >= 0; i-- {
		left += 1
		if trees[r][i] >= trees[r][c] {
			break
		}
	}
	return down * up * right * left

}

func main() {
	inputs := utils.StreamInput(
		func(line string) []int8 {
			trees := make([]int8, 0, len(line))
			for _, b := range []byte(line) {
				trees = append(trees, int8(b))
			}
			return trees
		},
	)

	fmt.Printf("Part 1: %v\n", Part1(inputs))
	inputs = utils.StreamInput(
		func(line string) []int8 {
			trees := make([]int8, 0, len(line))
			for _, b := range []byte(line) {
				trees = append(trees, int8(b))
			}
			return trees
		},
	)

	fmt.Printf("Part 2: %v\n", Part2(inputs))
}
