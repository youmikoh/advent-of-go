package main

import (
	"fmt"
	"strconv"
)

type Network struct {
	valves    map[string]*Valve
	flowable  map[string]int
	flowcache map[Flow]int
}

type Valve struct {
	v        string
	flowrate int
	tunnels  []string
	paths    map[string]int
}

type Path struct {
	v    string
	dist int
}

type Flow struct {
	v      string
	time   int
	opened int
}

func (n Network) ShortestPaths(v string) {
	paths := map[string]int{v: 0, "AA": 0}
	visited := map[string]bool{v: true}
	p := Path{v, 0}
	q := []Path{p}
	for len(q) > 0 {
		p, q = q[0], q[1:]
		fmt.Printf("p=%v\n", p)
		for _, t := range n.valves[p.v].tunnels {
			if visited[t] {
				continue
			}
			visited[t] = true
			if n.valves[t].flowrate > 0 {
				paths[t] = p.dist + 1
			}
			q = append(q, Path{t, p.dist + 1})
		}
	}
	delete(paths, v)
	if v != "AA" {
		delete(paths, "AA")
	}
	n.valves[v].paths = paths
}

func (n Network) MaxOutflow(f Flow) int {
	c, cached := n.flowcache[f]
	if cached {
		return c
	}
	m := 0
	for w, distance := range n.valves[f.v].paths {
		b := 1 << n.flowable[w]
		if f.opened&b != 0 { // already open
			continue
		}
		timeleft := f.time - distance - 1
		if timeleft <= 0 { // out of time
			continue
		}
		outflow := n.valves[w].flowrate * timeleft
		open := n.MaxOutflow(Flow{w, timeleft, f.opened | b})
		m = Max(m, outflow+open)
	}
	n.flowcache[f] = m
	return m
}

func Part1(n Network) int {
	return n.MaxOutflow(Flow{"AA", 30, 0})
}

func Part2(n Network) int {
	permutations := (1 << len(n.flowable)) - 1
	m := 0
	for p := 0; p < permutations; p++ {
		mine := n.MaxOutflow(Flow{"AA", 26, p})
		elephant := n.MaxOutflow(Flow{"AA", 26, permutations ^ p})
		m = Max(m, mine+elephant)
	}
	return m
}

func main() {
	valves := map[string]*Valve{}
	flowable := map[string]int{}

	inputs := StreamInputs()
	for input := range inputs {
		v := input[0][0]
		flowrate, _ := strconv.Atoi(input[0][1])
		valve := Valve{
			v:        v,
			flowrate: flowrate,
			tunnels:  input[1],
		}
		valves[v] = &valve

		if flowrate > 0 { // index valves with flow
			flowable[v] = len(flowable)
		}
	}

	n := Network{valves, flowable, map[Flow]int{}}
	for v, valve := range n.valves {
		if valve.flowrate == 0 && v != "AA" { // skip paths to noflow valves
			continue
		}
		n.ShortestPaths(v)
	}

	fmt.Printf("Part 1: %v\n", Part1(n))
	fmt.Printf("Part 2: %v\n", Part2(n))
}
