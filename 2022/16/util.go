package main

import (
	"advent-of-go/utils"
	"regexp"
	"strings"
)

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func StreamInputs() <-chan [2][]string {
	r := regexp.MustCompile(`^Valve ([A-Z]{2}) has flow rate=(\d+); tunnels? leads? to valves? (.*)$`)
	return utils.StreamInput(
		func(line string) [2][]string {
			subs := r.FindStringSubmatch(line)
			var tunnels []string
			for _, t := range strings.Split(subs[3], ",") {
				tunnels = append(tunnels, strings.TrimSpace(t))
			}
			return [2][]string{
				subs[1:3],
				tunnels,
			}
		},
	)
}
