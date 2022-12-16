package main

import (
	"fmt"
	"sort"
)

type Sensor struct {
	x, y     int
	distance int
}

func (s Sensor) Covered(p [2]int) bool {
	d := Distance([2]int{s.x, s.y}, p)
	return d <= s.distance
}

func (s Sensor) Range(y int) (int, int) {
	x0 := s.x - s.distance + Abs(s.y-y)
	x1 := s.distance - Abs(s.y-y) + s.x
	return x0, x1
}

func Distance(a, b [2]int) int {
	return Abs(a[0]-b[0]) + Abs(a[1]-b[1])
}

func Part1() int {
	inputs := StreamInputs()
	sensors := []Sensor{}
	beacons := make(map[[2]int]struct{})
	for input := range inputs {
		s := [2]int{input[0], input[1]}
		b := [2]int{input[2], input[3]}
		d := Distance(s, b)
		sensor := Sensor{
			x:        s[0],
			y:        s[1],
			distance: d,
		}
		sensors = append(sensors, sensor)
		beacons[b] = struct{}{}
	}

	row := 2000000
	intervals := make([][2]int, 0, len(sensors))
	for _, s := range sensors {
		if Abs(row-s.y) <= s.distance {
			a, b := s.Range(row)
			intervals = append(intervals, [2]int{a, b})
		}
	}

	sort.SliceStable(
		intervals,
		func(i, j int) bool {
			return intervals[i][0] < intervals[j][0]
		},
	)

	cover := make([][2]int, 0, len(sensors))
	cover = append(cover, intervals[0])
	i := 0
	for _, interval := range intervals {
		if interval[0] <= cover[i][1] {
			cover[i][1] = Max(cover[i][1], interval[1])
		} else {
			cover = append(cover, interval)
			i += 1
		}
	}

	count := 0
	for _, c := range cover {
		count += c[1] - c[0] + 1
	}
	for b := range beacons {
		if b[1] == row {
			count -= 1
		}
	}
	return count
}

func Part2() [2]int {
	inputs := StreamInputs()
	sensors := []Sensor{}
	for input := range inputs {
		s := [2]int{input[0], input[1]}
		b := [2]int{input[2], input[3]}
		d := Distance(s, b)
		sensor := Sensor{
			x:        s[0],
			y:        s[1],
			distance: d,
		}
		sensors = append(sensors, sensor)
	}

	bound := 4000000
	for y := 0; y <= bound; y++ {
	xLoop:
		for x := 0; x <= bound; x++ {
			for _, s := range sensors {
				if s.Covered([2]int{x, y}) {
					_, b := s.Range(y)
					x = b
					continue xLoop
				}
			}
			return [2]int{x, y}
		}
	}
	return [2]int{}
}

func main() {
	fmt.Printf("Part 1:%v\n", Part1())
	fmt.Printf("Part 2:%v\n", Part2())
}
