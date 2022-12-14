package main

import (
	"advent-of-go/utils"
	"fmt"
)

type Bounds struct {
	Left   int
	Right  int
	Top    int
	Bottom int
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

func MapRocks() [][2]int {
	inputs := StreamInput()
	// origin := [2]int{500, 0}
	// bounds := Bounds{500, 500, 0, 0}

	var rocks [][2]int
	for input := range inputs {
		head := input[0]
		rocks = append(rocks, head)

		for _, tail := range input[1:] {
			// bounds.Stretch(head)
			rocks = append(rocks, tail)

			dx, dy := Delta(head, tail)
			for x := head[0]; x != tail[0]; x += dx {
				rocks = append(rocks, [2]int{x, head[1]})
			}
			for y := head[1]; y != tail[1]; y += dy {
				rocks = append(rocks, [2]int{head[0], y})
			}
			head = tail
		}
	}
	return rocks
}

func BuildCave(rocks [][2]int, bounds Bounds) [][]rune {
	cave := make([][]rune, bounds.Top-bounds.Bottom+1)
	for i := range cave {
		cave[i] = make([]rune, bounds.Right-bounds.Left+1)
	}
	cave[0][500-bounds.Left] = 3

	for _, rock := range rocks {
		r, c := rock[1]-bounds.Bottom, rock[0]-bounds.Left
		cave[r][c] = 1
	}
	return cave
}

func Trickle(cave [][]rune, b Bounds) (r, c int) {
	x, y := 500, 0
	cols := b.Right - b.Left

	for x >= b.Left && x <= b.Right && y >= b.Bottom && y <= b.Top {
		y += 1
		r, c := y-b.Bottom, x-b.Left
		switch {
		case cave[r][c] == 0: // down
			continue
		case c > 0 && cave[r][c-1] == 0: // left down
			x -= 1
		case c < cols && cave[r][c+1] == 0: // right down
			x += 1
		case c == 0 || c == cols:
			return -1, -1
		default: // at rest
			return r - 1, c
		}
	}
	return -1, -1
}

func CountSand(cave [][]rune) int {
	count := 0
	for _, c := range cave {
		for _, i := range c {
			if i == 2 {
				count += 1
			}
		}
	}
	return count
}

func Part1() int {
	rocks := MapRocks()
	bounds := Bounds{500, 500, 0, 0}
	for _, r := range rocks {
		bounds.Stretch(r)
	}

	cave := BuildCave(rocks, bounds)
	for {
		// fmt.Printf("\033[H\033[2J")
		// PrintCave(cave)
		r, c := Trickle(cave, bounds)
		if r == -1 || c == -1 {
			break
		}
		cave[r][c] = 2
		// time.Sleep(500 * time.Millisecond)
	}
	return CountSand(cave)

}

func Part2() int {
	rocks := MapRocks()
	bounds := Bounds{500, 500, 0, 0}
	for _, r := range rocks {
		bounds.Stretch(r)
	}
	bounds.Left -= 100000
	bounds.Right += 100000
	bounds.Top += 2

	cave := BuildCave(rocks, bounds)
	for i := range cave[bounds.Top] {
		cave[bounds.Top][i] = 1
	}

	for {
		// fmt.Printf("\033[H\033[2J")
		// PrintCave(cave)
		r, c := Trickle(cave, bounds)
		if r == -1 || c == -1 {
			break
		}
		if r == 0 && c == (500-bounds.Left) {
			cave[r][c] = 2
			break
		}
		cave[r][c] = 2
		// time.Sleep(100 * time.Millisecond)
	}

	return CountSand(cave)
}

func main() {
	fmt.Printf("Part 1: %v\n", Part1())
	fmt.Printf("Part 2: %v\n", Part2())
}
