package day1

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/AnuragProg/aoc2024/utils"
)


func Part1(){
	input, err := utils.ReadInput("./puzzles/day1/input.txt")
	if err != nil {
		panic(err)
	}

	n := len(input)
	left, right := make([]uint64, 0, n), make([]uint64, 0, n)

	safe := 10
	for _, line := range input {
		split := strings.Fields(line)
		if safe > 0 {
			fmt.Println(split[0], split[1])
			safe -- 
		}
		leftItem, err := strconv.ParseUint(split[0], 10, 64)
		if err != nil {
			panic(err)
		}
		rightItem, err := strconv.ParseUint(split[1], 10, 64)
		if err != nil {
			panic(err)
		}
		left = append(left, leftItem)
		right = append(right, rightItem)
	}

	slices.Sort(left)
	slices.Sort(right)
	fmt.Println(left[:5])
	fmt.Println(right[:5])

	var distance uint64 = 0
	for i := range n {
		if left[i] > right[i] {
			distance += left[i] - right[i]
		}else {
			distance += right[i] - left[i]
		}
	}
	fmt.Printf("Day1-Part1: %v\n", distance)
}
