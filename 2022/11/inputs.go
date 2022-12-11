package main

var TestMonkeys []Monkey = []Monkey{
	{
		Items:   []uint64{79, 98},
		Operate: func(old uint64) uint64 { return old * 19 },
		Test: func(i uint64) int {
			if i%23 == 0 {
				return 2
			}
			return 3
		},
		Mod: 23,
	},
	{
		Items:   []uint64{54, 65, 75, 74},
		Operate: func(old uint64) uint64 { return old + 6 },
		Test: func(i uint64) int {
			if i%19 == 0 {
				return 2
			}
			return 0
		},
		Mod: 19,
	},
	{
		Items:   []uint64{79, 60, 97},
		Operate: func(old uint64) uint64 { return old * old },
		Test: func(i uint64) int {
			if i%13 == 0 {
				return 1
			}
			return 3
		},
		Mod: 13,
	},
	{
		Items:   []uint64{74},
		Operate: func(old uint64) uint64 { return old + 3 },
		Test: func(i uint64) int {
			if i%17 == 0 {
				return 0
			}
			return 1
		},
		Mod: 17,
	},
}

var InputMonkeys []Monkey = []Monkey{
	{
		Items:   []uint64{91, 54, 70, 61, 64, 64, 60, 85},
		Operate: func(old uint64) uint64 { return old * 13 },
		Test: func(i uint64) int {
			if i%2 == 0 {
				return 5
			}
			return 2
		},
		Mod: 2,
	},
	{
		Items:   []uint64{82},
		Operate: func(old uint64) uint64 { return old + 7 },
		Test: func(i uint64) int {
			if i%13 == 0 {
				return 4
			}
			return 3
		},
		Mod: 13,
	},
	{
		Items:   []uint64{84, 93, 70},
		Operate: func(old uint64) uint64 { return old + 2 },
		Test: func(i uint64) int {
			if i%5 == 0 {
				return 5
			}
			return 1
		},
		Mod: 5,
	},
	{
		Items:   []uint64{78, 56, 85, 93},
		Operate: func(old uint64) uint64 { return (old * 2) },
		Test: func(i uint64) int {
			if i%3 == 0 {
				return 6
			}
			return 7
		},
		Mod: 3,
	},
	{
		Items:   []uint64{64, 57, 81, 95, 52, 71, 58},
		Operate: func(old uint64) uint64 { return old * old },
		Test: func(i uint64) int {
			if i%11 == 0 {
				return 7
			}
			return 3
		},
		Mod: 11,
	},
	{
		Items:   []uint64{58, 71, 96, 58, 68, 90},
		Operate: func(old uint64) uint64 { return old + 6 },
		Test: func(i uint64) int {
			if i%17 == 0 {
				return 4
			}
			return 1
		},
		Mod: 17,
	},
	{
		Items:   []uint64{56, 99, 89, 97, 81},
		Operate: func(old uint64) uint64 { return old + 1 },
		Test: func(i uint64) int {
			if i%7 == 0 {
				return 0
			}
			return 2
		},
		Mod: 7,
	},
	{
		Items:   []uint64{68, 72},
		Operate: func(old uint64) uint64 { return old + 8 },
		Test: func(i uint64) int {
			if i%19 == 0 {
				return 6
			}
			return 0
		},
		Mod: 19,
	},
}
