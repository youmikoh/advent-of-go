package main

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func PrintCave(cave [][]rune) {
	for _, c := range cave {
		// fmt.Printf("%v", i)
		for _, r := range c {
			if r == 1 {
				fmt.Printf("%v", string('X'))
			}
			if r == 0 {
				fmt.Printf("%v", string('-'))
			}
			if r == 2 {
				fmt.Printf("%v", string('O'))
			}
			if r == 3 {
				fmt.Printf("%v", string('+'))
			}
		}
		fmt.Println("")
	}
}

func StreamInput() <-chan [][2]int {
	return utils.StreamInput(
		func(line string) [][2]int {
			line = strings.ReplaceAll(line, "->", "")
			s := strings.Fields(line)
			pairs := make([][2]int, 0, len(s))
			for _, pair := range s {
				p := strings.Split(pair, ",")
				x, _ := strconv.Atoi(p[0])
				y, _ := strconv.Atoi(p[1])
				pairs = append(pairs, [2]int{x, y})
			}
			return pairs
		},
	)
}

func (b *Bounds) Stretch(rock [2]int) {
	if rock[0] < b.Left {
		b.Left = rock[0]
	}
	if rock[0] > b.Right {
		b.Right = rock[0]
	}
	if rock[1] < b.Bottom {
		b.Bottom = rock[1]
	}
	if rock[1] > b.Top {
		b.Top = rock[1]
	}
}

func Delta(h, t [2]int) (int, int) {
	return utils.Sign(t[0] - h[0]), utils.Sign(t[1] - h[1])
}
