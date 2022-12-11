package main

import (
	"advent-of-go/utils"
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Directory struct {
	Parent   *Directory
	Children map[string]*Directory
	Files    map[string]int
	Size     int
}

func (d *Directory) AddChild(name string) Directory {
	child := Directory{
		Parent:   d,
		Children: make(map[string]*Directory),
		Files:    make(map[string]int),
	}
	d.Children[name] = &child
	return child
}

func (d *Directory) CalculateSize() {
	for _, c := range d.Children {
		if c.Size == 0 {
			c.CalculateSize()
		}
		d.Size += c.Size
	}
	for _, size := range d.Files {
		d.Size += size
	}
}

func SumSmallerThan(d Directory, cap int) int {
	sum := 0
	if d.Size < cap {
		sum += d.Size
	}
	for _, c := range d.Children {
		sum += SumSmallerThan(*c, cap)
	}
	return sum
}

func SmallestGreaterThan(d Directory, floor int) int {
	sizes := []int{}
	if d.Size > floor {
		sizes = append(sizes, d.Size)
	}
	for _, c := range d.Children {
		s := SmallestGreaterThan(*c, floor)
		if s > 0 {
			sizes = append(sizes, s)
		}
	}
	if len(sizes) > 0 {
		sort.Ints(sizes)
		return sizes[0]
	} else {
		return -1
	}
}

func main() {
	file := utils.LoadFile("input.txt")
	scanner := bufio.NewScanner(file)

	root := &Directory{
		Children: make(map[string]*Directory),
		Files:    make(map[string]int),
	}
	current := root

	for scanner.Scan() {
		c := strings.Fields(scanner.Text())
		if c[0] == "$" {
			if c[1] == "ls" { // $ ls
				continue
			} else if c[2] == ".." { // $ cd ..
				current = current.Parent
			} else { // $ cd {dir-name}
				child, exists := current.Children[c[1]]
				if exists {
					current = child
				} else {
					child := current.AddChild(c[2])
					current = &child
				}
			}
		} else if c[0] == "dir" { // $ dir {dir-name}
			_, exists := current.Children[c[1]]
			if !exists {
				_ = current.AddChild(c[1])
			}
		} else { // {file-size} {file-name}
			_, exists := current.Files[c[1]]
			if !exists {
				size, _ := strconv.Atoi(c[0])
				current.Files[c[1]] = size
			}
		}
	}

	root.CalculateSize()

	fmt.Printf("Part 1: %v\n", SumSmallerThan(*root, 100000))
	requiredSpace := root.Size + 30000000 - 70000000
	fmt.Printf("Part 2: %v\n", SmallestGreaterThan(*root, requiredSpace))
}
