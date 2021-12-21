package main

import (
	"strconv"

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
		n, err := strconv.Atoi(input.Text())
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
