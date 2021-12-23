package main

import (
	"sort"
	"strconv"
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

type Bound struct {
	n, e, s, w int
	visited    [][]bool
}

// part solution
// the visited 2d array
// - iterate over the nums 2d array attempting to start a new basin at any point
// - if can start then recursively expand the basin keeping track of how much expanded
// every time you do one of these update the visited 2d array
//
// just only update the visited 2d array when part2 is true
// overwise just make it full of false

//   n
// w + e
//   s

// return true if y, x is within not out of arr and has not been visited
func (b *Bound) check(y, x int) bool {
	return y > b.n && x < b.e && y < b.s && x > b.w && !b.visited[y][x]
}

// will return the number if less than every surrounding number
// else will return 0
func (b *Bound) lowest(nums [][]int, y, x int) int {
	if !b.check(y, x) {
		return 0
	}
	isLowest := true

	val := nums[y][x]
	if b.check(y-1, x) && nums[y-1][x] <= val {
		isLowest = false
	}
	if b.check(y, x-1) && nums[y][x-1] <= val {
		isLowest = false
	}
	if b.check(y+1, x) && nums[y+1][x] <= val {
		isLowest = false
	}
	if b.check(y, x+1) && nums[y][x+1] <= val {
		isLowest = false
	}

	if isLowest {
		return val + 1
	}
	return 0
}

func (b *Bound) findBasin(nums [][]int, y, x int) int {
	if !b.check(y, x) || nums[y][x] == 9 {
		return 0
	}

	// consider [y][x] as visited and check the n,e,s,w if they are part of the basin
	b.visited[y][x] = true
	return 1 + b.findBasin(nums, y-1, x) + b.findBasin(nums, y+1, x) + b.findBasin(nums, y, x-1) + b.findBasin(nums, y, x+1)
}

func sumOfLowPoints(filepath string, part2 bool) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	nums := [][]int{}
	visit := [][]bool{}
	for input.Scan() {
		newNums := []int{}
		line := strings.Split(input.Text(), "")
		for _, val := range line {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			newNums = append(newNums, num)
		}
		nums = append(nums, newNums)
		visit = append(visit, make([]bool, len(newNums)))
	}
	if len(nums) == 0 {
		return 0
	}

	// the code from beyond here assumes that input is a rectangle
	b := Bound{
		n:       -1,
		e:       len(nums[0]),
		s:       len(nums),
		w:       -1,
		visited: visit,
	}

	if !part2 {
		sum := 0

		for y, row := range nums {
			for x := range row {
				sum += b.lowest(nums, y, x)
			}
		}
		return sum
	}

	basins := []int{}
	for y, row := range nums {
		for x := range row {
			basins = append(basins, b.findBasin(nums, y, x))
		}
	}

	sort.Ints(basins)
	if len(basins) < 3 {
		return 0
	}
	blen := len(basins)

	return basins[blen-1] * basins[blen-2] * basins[blen-3]
}
