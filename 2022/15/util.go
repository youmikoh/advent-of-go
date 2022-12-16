package main

import (
	"advent-of-go/utils"
	"regexp"
	"strconv"
)

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func StreamInputs() <-chan [4]int {
	r := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	return utils.StreamInput(
		func(line string) [4]int {
			subs := r.FindStringSubmatch(line)
			ints := [4]int{}
			for i, s := range subs[1:] {
				n, err := strconv.Atoi(s)
				if err != nil {
					panic(err)
				}
				ints[i] = n
			}
			return ints
		},
	)
}
