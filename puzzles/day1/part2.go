package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/AnuragProg/aoc2024/utils"
)

func Part2(){
	input, err := utils.ReadInput("./puzzles/day1/input.txt")
	if err != nil {
		panic(err)
	}

	rightCountMap := map[uint64]uint64{}
	leftCountMap := map[uint64]uint64{}
	for _, line := range input {
		split := strings.Fields(line)
		leftItem, err := strconv.ParseUint(split[0], 10, 64)
		if err != nil {
			panic(err)
		}
		rightItem, err := strconv.ParseUint(split[1], 10, 64)
		if err != nil {
			panic(err)
		}

		leftCountMap[leftItem] += 1
		rightCountMap[rightItem] += 1
	}

	similarityScore := uint64(0)
	for left, leftCount := range leftCountMap {
		if leftCountInRight, ok := rightCountMap[left]; ok {
			similarityScore += leftCount * left * leftCountInRight
		}
	}

	fmt.Printf("Day1-Part2: %v\n", similarityScore)
}
