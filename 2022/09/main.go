package main

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
	X int
	Y int
}

type Move struct {
	Direction string
	Units     int
}

func (m Move) Delta() (int, int) {
	switch m.Direction {
	case "R":
		return 1, 0
	case "L":
		return -1, 0
	case "U":
		return 0, 1
	case "D":
		return 0, -1
	}
	panic("uh oh")
}

func (tail *Position) Pull(head Position) {
	dx, dy := head.X-tail.X, head.Y-tail.Y
	adx, ady := utils.Abs(dx), utils.Abs(dy)

	if dx == 0 || dy == 0 { // not diag
		if adx > 1 {
			tail.X += dx / adx
		}
		if ady > 1 {
			tail.Y += dy / ady
		}
	} else if adx+ady > 2 { // diag, no touch
		tail.X += dx / adx
		tail.Y += dy / ady
	} // else touch
}

func Part1(inputs <-chan Move) int {
	return Roped(inputs, 1)
}

func Part2(inputs <-chan Move) int {
	return Roped(inputs, 9)
}

func Roped(moves <-chan Move, length int) int {
	head, trail := Position{}, make([]Position, length)
	visited := make(map[Position]bool)

	for move := range moves {
		dx, dy := move.Delta()
		for move.Units > 0 {
			head.X += dx
			head.Y += dy

			prev := head
			for i := range trail {
				trail[i].Pull(prev)
				prev = trail[i]
			}
			visited[prev] = true
			move.Units -= 1
		}
	}
	count := 0
	for _, v := range visited {
		if v {
			count += 1
		}
	}
	return count
}

func main() {
	inputs := utils.StreamInput(
		func(line string) Move {
			m := strings.Fields(line)
			u, _ := strconv.Atoi(m[1])
			return Move{
				Direction: m[0],
				Units:     u,
			}
		},
	)
	fmt.Printf("Part 1:  %v\n", Part1(inputs))
	// fmt.Printf("Part 2:  %v\n", Part2(inputs))
}
