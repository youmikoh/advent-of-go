package main

import (
	"fmt"
	"os"
)

func PacketMarker(inputs []byte, uniques int) int {
	for i := range inputs {
		hash := make(map[byte]byte)
		for _, prev := range inputs[i : i+uniques] {
			hash[prev] = prev
		}
		if len(hash) == uniques {
			return i + uniques
		}
	}
	return -1
}

func Part1(inputs []byte) int {
	return PacketMarker(inputs, 4)
}

func Part2(inputs []byte) int {
	return PacketMarker(inputs, 14)
}

func main() {
	inputs, _ := os.ReadFile("input.txt")
	fmt.Printf("Part 1, start-of-packet marker on %v\n", Part1(inputs))
	fmt.Printf("Part 2, start-of-packet marker on %v\n", Part2(inputs))
}
