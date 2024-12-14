package day2

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/AnuragProg/aoc2024/utils"
)


/*
 1 2 7 8 9

*/


func Part2() {
	input := utils.Must(utils.ReadInput("./puzzles/day2/input.txt"))

	safeCount := 0

	for _, line := range input {
		numStrs := strings.Fields(line)
		nums := make([]int, len(numStrs))
		for idx, numStr := range numStrs {
			nums[idx] = utils.Must(strconv.Atoi(numStr))
		}

		inFavorOfInc, inFavorOfDec, equals := 0, 0, 0

		isSafe := true

		for idx:=0; idx<len(nums)-1; idx++ {
			if nums[idx+1] - nums[idx] > 0 && nums[idx+1] - nums[idx] <= 3{
				inFavorOfInc++
			}else if nums[idx] - nums[idx+1] > 0 && nums[idx] - nums[idx+1] <= 3{
				inFavorOfDec++
			}else if nums[idx] == nums[idx+1]{
				equals++
				if equals > 1 {
					isSafe = false
					break
				}
			}
		}

		if isSafe {
			safeCount += 1
		}
	}
	fmt.Printf("Day2-Part1: %v\n", safeCount)
}
