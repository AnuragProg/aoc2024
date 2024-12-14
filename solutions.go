package main

import (
	"github.com/AnuragProg/aoc2024/puzzles/day1"
	"github.com/AnuragProg/aoc2024/puzzles/day2"
)

var solutions map[int]map[int]func() = map[int]map[int]func(){
	1: { // day 1
		1: day1.Part1,// part 1
		2: day1.Part2,
	},
	2: { // day 2
		1: day2.Part1,
		2: day2.Part2,
	},
}
