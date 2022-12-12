package main

import "fmt"

func (n Node) Print() {
	neighbours := []string{}
	for _, node := range n.Neighbours {
		neighbours = append(neighbours, string(node.Height))
	}
	fmt.Printf("NODE %v: (%v,%v) > Neighbours: %v\n", string(n.Height), n.Row, n.Col, neighbours)
}

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Distance < h[j].Distance }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(p interface{}) { *h = append(*h, p.(Path)) }

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	p := old[n-1]
	*h = old[0 : n-1]
	return p
}
