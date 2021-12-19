package main

import (
	"strconv"

	"github.com/koitu/advent-of-code-2021/utils"
)

// https://adventofcode.com/2021/day/1
func countIncreases(filepath string, gap int) int {
	if gap == 0 {
		return 0
	}

	reader, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	// load all the values of input.txt into a slice as integers
	nums := []int{}
	for reader.Scan() {
		n, err := strconv.Atoi(reader.Text())
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	count := 0
	cur := 0
	max := len(nums)

	for cur+gap < max {
		if nums[cur] < nums[cur+gap] {
			count++
		}
		cur++
	}

	return count
}
