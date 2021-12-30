package main

import (
	"github.com/koitu/advent-of-code-2021/utils"
)

func countIncreases(filepath string, gap int) int {
	if gap == 0 {
		return 0
	}

	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	// load all the values of input.txt into a slice as integers
	nums := []int{}
	for input.Scan() {
		nums = append(nums, utils.Atoi(input.Text()))
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
