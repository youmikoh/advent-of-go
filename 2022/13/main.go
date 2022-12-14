package main

import (
	"advent-of-go/utils"
	"bufio"
	"fmt"
	"sort"
	"strconv"
)

func inOrderInts(left int, right int) int {
	return left - right
}

func inOrderSlices(left []interface{}, right []interface{}) int {
	for i, l := range left {
		if i >= len(right) {
			return 1
		}
		if r := inOrder(l, right[i]); r != 0 {
			return r
		}
	}
	return len(left) - len(right)
}

func inOrder(left interface{}, right interface{}) int {
	switch l := left.(type) {
	case int:
		switch r := right.(type) {
		case int:
			return inOrderInts(l, r)
		case []interface{}:
			return inOrderSlices([]interface{}{l}, r)
		}
	case []interface{}:
		switch r := right.(type) {
		case int:
			return inOrderSlices(l, []interface{}{r})

		case []interface{}:
			return inOrderSlices(l, r)
		}
	}
	panic("uh oh")
}

func Part1() int {
	file := utils.LoadFile("input.txt")
	scanner := bufio.NewScanner(file)

	var inputs [][][]interface{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		pair := make([][]interface{}, 0, 2)
		pair = append(pair, BuildList(line))
		scanner.Scan()
		pair = append(pair, BuildList(scanner.Text()))
		inputs = append(inputs, pair)
	}

	sum := 0
	for i, input := range inputs {
		if inOrder(input[0], input[1]) <= 0 {
			sum += i + 1
		}
	}
	return sum
}

func Part2() int {
	file := utils.LoadFile("input.txt")
	scanner := bufio.NewScanner(file)

	var inputs [][]interface{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		inputs = append(inputs, BuildList(line))
	}
	div1 := BuildList("[[2]]")
	div2 := BuildList("[[6]]")

	inputs = append(inputs, div1)
	inputs = append(inputs, div2)

	sort.Slice(
		inputs,
		func(i, j int) bool {
			return inOrder(inputs[i], inputs[j]) < 0
		},
	)

	d1 := sort.Search(
		len(inputs),
		func(i int) bool {
			return inOrder(inputs[i], div1) >= 0
		},
	)
	d2 := sort.Search(
		len(inputs),
		func(i int) bool {
			return inOrder(inputs[i], div2) >= 0
		},
	)

	return (d1 + 1) * (d2 + 1)
}

func BuildList(str string) []interface{} {
	stack := [][]interface{}{}
	num := []byte{}
	last := -1
	for _, s := range []byte(str) {
		if s == '[' {
			stack = append(stack, []interface{}{})
			last += 1
		} else if s == ']' {
			if len(num) > 0 {
				n, _ := strconv.Atoi(string(num))
				stack[last] = append(stack[last], n)
				num = []byte{}
			}
			if last == 0 {
				return stack[last]
			}
			stack[last-1] = append(stack[last-1], stack[last])
			stack = stack[:last]
			last -= 1
		} else if s == ',' {
			if len(num) > 0 {
				n, _ := strconv.Atoi(string(num))
				stack[last] = append(stack[last], n)
				num = []byte{}
			}
		} else { // digit
			num = append(num, s)
		}
	}
	return stack[last]
}

func main() {
	fmt.Printf("Part 1: %v\n", Part1())
	fmt.Printf("Part 2: %v\n", Part2())
}
