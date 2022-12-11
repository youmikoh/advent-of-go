package main

import (
	"fmt"
	"sort"
)

type Monkey struct {
	Items       []uint64
	Operate     func(uint64) uint64
	Test        func(uint64) int
	Mod         uint64
	Inspections int
}

func Part1(monkeys []Monkey) int {
	return KeepAway(
		monkeys,
		20,
		func(item uint64) uint64 {
			return item / 3
		},
	)
}

func Part2(monkeys []Monkey) int {
	var lcm uint64 = 1 // least common multiple, mods are all prime
	for _, m := range monkeys {
		lcm *= uint64(m.Mod)
	}

	return KeepAway(
		monkeys,
		10000,
		func(item uint64) uint64 {
			return item % lcm
		},
	)
}

func KeepAway(monkeys []Monkey, rounds int, worry func(uint64) uint64) int {
	for r := 0; r < rounds; r++ {
		for m := 0; m < len(monkeys); m++ {
			for _, item := range monkeys[m].Items {
				item := worry(monkeys[m].Operate(item))
				monkeys[m].Inspections += 1

				next := monkeys[m].Test(item)
				monkeys[next].Items = append(monkeys[next].Items, item)
			}
			monkeys[m].Items = nil
		}
	}
	return MonkeyBusiness(monkeys)
}

func MonkeyBusiness(monkeys []Monkey) int {
	var inspections []int
	for _, monkey := range monkeys {
		inspections = append(inspections, monkey.Inspections)
	}
	sort.Slice(
		inspections,
		func(i, j int) bool {
			return inspections[i] > inspections[j]
		},
	)
	return inspections[0] * inspections[1]
}

func main() {
	fmt.Println("golang doesn't have eval 🙈")
	fmt.Printf("Part 1: %v\n", Part1(InputMonkeys))
	fmt.Printf("Part 2: %v\n", Part2(InputMonkeys))
}
