package main

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

type Operation struct {
	Duration  int
	Increment int
}

func (op *Operation) Run(acc *int) bool {
	op.Duration -= 1
	if op.Duration == 0 {
		*acc = *acc + op.Increment
		return true
	}
	return false
}

func Part1(inputs <-chan Operation) int {
	signal := 0
	x := 1
	op, ok := <-inputs
	if !ok {
		panic("uh oh")
	}
	for cycle := 1; ; cycle++ {
		if cycle%40 == 20 {
			prod := cycle * x
			signal += prod
		}
		done := op.Run(&x)
		if !done {
			continue
		}
		op, ok = <-inputs
		if !ok {
			break
		}
	}
	return signal
}

func CycleByte(x, c int) byte {
	sprite := []int{x - 1, x, x + 1}
	for _, i := range sprite {
		if i == c {
			return '#'
		}
	}
	return ' '
}

func Part2(inputs <-chan Operation) [][]byte {
	var crt [][]byte
	var line []byte
	x, index := 1, 0

	op, ok := <-inputs
	if !ok {
		panic("uh oh")
	}
	for cycle := 1; ; cycle++ {
		if index%40 == 0 {
			crt = append(crt, line)
			line, index = []byte{}, 0
		}
		line = append(line, CycleByte(x, index))
		index += 1

		done := op.Run(&x)
		if !done {
			continue
		}
		op, ok = <-inputs
		if !ok {
			break
		}
	}
	return append(crt, line)
}

func main() {
	inputs := utils.StreamInput(
		func(line string) Operation {
			op := strings.Fields(line)
			if len(op) > 1 {
				v, _ := strconv.Atoi(op[1])
				return Operation{2, v}
			}
			return Operation{1, 0}
		},
	)

	// fmt.Printf("Part 1: %v\n", Part1(inputs))
	fmt.Println("Part 2:")
	for _, line := range Part2(inputs) {
		fmt.Println(string(line))
	}
}
