package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/koitu/advent-of-code-2021/utils"
)

func countZeros(nums *[]uint64, col uint) int {
	count := 0
	for _, num := range *nums {
		if num&(1<<col) == 0 {
			count++
		}
	}
	return count
}

func bitFilter(nums []uint64, unmask uint64, match uint64) []uint64 {
	lastEle := nums[len(nums)-1]
	var newNums []uint64

	for i := 0; i < len(nums); i++ {
		if (nums[i] & unmask) == match {
			newNums = append(newNums, nums[i])
		}
	}

	if len(newNums) == 0 {
		newNums = append(newNums, lastEle)
	}

	return newNums
}

func binaryDiagnostic(filepath string, part2 bool) int {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(file)
	scan.Scan()
	lineLen := len([]rune(scan.Text()))
	file.Close()

	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	nums := []uint64{}
	for input.Scan() {
		str := input.Text()
		val, err := strconv.ParseInt(str, 2, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, uint64(val))
	}

	if !part2 {
		lineNum := len(nums)
		g := 0
		mask := 0

		for col := lineLen - 1; col >= 0; col-- {
			g = g << 1
			mask = mask << 1
			mask++
			zeroNum := countZeros(&nums, uint(col))
			if 2*zeroNum == lineNum {
				panic("Even number of 1's and 0's")
			} else if 2*zeroNum < lineNum {
				g++
			}
		}

		return int(g * (g ^ mask))
	}

	nums_copy := make([]uint64, len(nums))
	copy(nums_copy, nums)
	var match uint64

	// could create a function to make code more DRY
	for col := lineLen - 1; col >= 0; col-- {
		zeroNum := countZeros(&nums, uint(col))
		if 2*zeroNum <= len(nums) {
			match = 1
		} else {
			match = 0
		}

		nums = bitFilter(nums, 1<<col, match<<col)

		if len(nums) == 1 {
			break
		}
	}

	for col := lineLen - 1; col >= 0; col-- {
		match = 0
		zeroNum := countZeros(&nums_copy, uint(col))
		if 2*zeroNum > len(nums_copy) {
			match = 1
		} else {
			match = 0
		}

		nums_copy = bitFilter(nums_copy, 1<<col, match<<col)

		if len(nums_copy) == 1 {
			break
		}
	}

	return int(nums[0] * nums_copy[0])
}
