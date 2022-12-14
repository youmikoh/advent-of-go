package main

import (
	"advent-of-go/utils"
	"container/heap"
	"fmt"
	"math"
)

type Node struct {
	Height     rune
	Row, Col   int
	Neighbours []*Node
}

type Path struct {
	Distance int
	LastNode *Node
}

type MinPath []Path

func (start *Node) Dijkstra(end *Node) int {
	visited := make(map[*Node]bool)
	h := &MinPath{Path{Distance: 0, LastNode: start}}
	for h.Len() > 0 {
		p := heap.Pop(h).(Path)
		if visited[p.LastNode] {
			continue
		}
		if p.LastNode == end {
			return p.Distance
		}
		for _, node := range p.LastNode.Neighbours {
			if !visited[node] {
				heap.Push(h, Path{Distance: p.Distance + 1, LastNode: node})
			}
		}
		visited[p.LastNode] = true
	}
	return math.MaxInt
}

func Part1(inputs <-chan []rune) int {
	var nodes []*Node
	var start, end *Node
	r := 0
	for input := range inputs {
		for c, h := range input {
			node := Node{Height: h, Row: r, Col: c}
			nodes = append(nodes, &node)
			if h == 'S' {
				node.Height = 'a'
				start = &node
			}
			if h == 'E' {
				node.Height = 'z'
				end = &node
			}
		}
		r += 1
	}
	AddNeighbours(nodes, r, len(nodes)/r)
	return start.Dijkstra(end)
}

func Part2(inputs <-chan []rune) int {
	var nodes, starts []*Node
	var end *Node
	r := 0
	for input := range inputs {
		for c, h := range input {
			node := Node{Height: h, Row: r, Col: c}
			nodes = append(nodes, &node)
			if h == 'a' {
				starts = append(starts, &node)
			}
			if h == 'S' {
				node.Height = 'a'
				starts = append(starts, &node)
			}
			if h == 'E' {
				node.Height = 'z'
				end = &node
			}
		}
		r += 1
	}
	AddNeighbours(nodes, r, len(nodes)/r)
	shortest := math.MaxInt
	for _, start := range starts {
		if d := start.Dijkstra(end); d < shortest {
			shortest = d
		}
	}
	return shortest
}

func (n *Node) AddIfNeighbour(nbr *Node) {
	if nbr.Height-n.Height < 2 {
		n.Neighbours = append(n.Neighbours, nbr)
	}
}

func AddNeighbours(nodes []*Node, Rows, Cols int) {
	for _, n := range nodes {
		if i := n.Col - 1; i >= 0 {
			n.AddIfNeighbour(nodes[n.Row*Cols+i])
		}
		if i := n.Col + 1; i < Cols {
			n.AddIfNeighbour(nodes[n.Row*Cols+i])
		}
		if i := n.Row - 1; i >= 0 {
			n.AddIfNeighbour(nodes[i*Cols+n.Col])
		}
		if i := n.Row + 1; i < Rows {
			n.AddIfNeighbour(nodes[i*Cols+n.Col])
		}
	}
}
func main() {
	inputs := utils.StreamInput(
		func(line string) []rune {
			return []rune(line)
		},
	)

	fmt.Printf("Part 1: %v\n", Part1(inputs))
	fmt.Printf("Part 2: %v\n", Part2(inputs))
}
