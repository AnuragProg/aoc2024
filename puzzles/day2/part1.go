package day2

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/AnuragProg/aoc2024/utils"
)

func Part1() {
	input := utils.Must(utils.ReadInput("./puzzles/day2/input.txt"))

	safeCount := 0

	outer:
	for _, line := range input {
		numStrs := strings.Fields(line)
		nums := make([]int, len(numStrs))
		for idx, numStr := range numStrs {
			nums[idx] = utils.Must(strconv.Atoi(numStr))
		}

		if nums[0] == nums[1] {
			continue
		}
		increasing := nums[0] < nums[1]
		for i := 1; i < len(nums); i++ {
			if math.Abs(float64(nums[i-1]-nums[i])) > 3 || math.Abs(float64(nums[i-1]-nums[i])) < 1 {
				continue outer
			}
			correctSequence := (increasing && nums[i-1] < nums[i]) || (!increasing && nums[i-1] > nums[i])
			if !correctSequence {
				continue outer
			}
		}
		safeCount += 1
	}
	fmt.Printf("Day2-Part1: %v\n", safeCount)
}
