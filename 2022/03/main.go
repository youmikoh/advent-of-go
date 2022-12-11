package main

import (
	"advent-of-go/utils"
	"fmt"
	"sync"
)

func Priority(b byte) int {
	if b < 91 {
		return int(b - 64 + 26)
	}
	return int(b - 96)
}

func Part1(inputs <-chan []byte) int {
	sum := 0
	for input := range inputs {
		mid := int(len(input) / 2)
		hash := make(map[byte]byte)
		for _, b0 := range input[:mid] {
			hash[b0] = b0
		}
		for _, b1 := range input[mid:] {
			_, err := hash[b1]
			if err {
				sum += Priority(b1)
				break
			}
		}
	}
	return sum
}

func Part2(inputs <-chan []byte) int {
	groups := [][][]byte{}
	g := [][]byte{}
	i := 0
	for elf := range inputs {
		g = append(g, elf)
		if i%3 == 2 {
			groups = append(groups, g)
			g = [][]byte{}
		}
		i += 1
	}

	sum := 0
	for _, g := range groups {
		badge := HashGroup(g)
		sum += Priority(badge)
	}
	return sum
}

func HashGroup(g [][]byte) byte {
	hash0 := make(map[byte]byte)
	for _, b0 := range g[0] {
		hash0[b0] = b0
	}

	hash1 := make(map[byte]byte)
	for _, b1 := range g[1] {
		_, err := hash0[b1]
		if err {
			hash1[b1] = b1
		}
	}

	for _, b2 := range g[2] {
		_, err := hash1[b2]
		if err {
			return b2
		}
	}
	panic(g)
}

func ConcurrentPart2(inputs <-chan []byte) int {
	groups := [][][]byte{}
	g := [][]byte{}
	line := 0
	for elf := range inputs {
		g = append(g, elf)
		if line%3 == 2 {
			groups = append(groups, g)
			g = [][]byte{}
		}
		line += 1
	}

	badges := make(chan byte)

	go func() {
		var wg sync.WaitGroup
		defer close(badges)

		for _, g := range groups {
			wg.Add(1)
			go func() {
				defer wg.Done()
				badges <- HashGroup(g)
			}()
		}

		wg.Wait()
	}()

	priorities := 0
	for b := range badges {
		priorities += Priority(b)
	}

	return priorities
}

func main() {
	inputs := utils.StreamInput(
		func(line string) []byte {
			return []byte(line)
		},
	)
	// fmt.Printf("Part 1: %v\n", Part1(inputs))
	// fmt.Printf("Part 2: %v\n", Part2(inputs))
	fmt.Printf("Goroutined Part 2: %v\n", ConcurrentPart2(inputs))
}
