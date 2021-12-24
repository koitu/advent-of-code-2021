package main

import (
	"strconv"
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

type cavern struct {
	grid                [][]int
	ylen, xlen, flashes int
	part2               bool
}

// checks if y,x are coordinates within the cavern
func (b *cavern) check(y, x int) bool {
	return y < b.ylen && y > -1 && x < b.xlen && x > -1
}

func (b *cavern) increase(y, x int) {
	if !b.check(y, x) {
		return
	}

	b.grid[y][x]++
	if b.grid[y][x] == 10 {
		b.increase(y-1, x)
		b.increase(y+1, x)
		b.increase(y, x-1)
		b.increase(y, x+1)
		b.increase(y-1, x-1)
		b.increase(y-1, x+1)
		b.increase(y+1, x-1)
		b.increase(y+1, x+1)
	}
}

func (b *cavern) update() bool {
	for y, row := range b.grid {
		for x := range row {
			b.increase(y, x)
		}
	}
	allFlash := true

	for y, row := range b.grid {
		for x := range row {
			if b.grid[y][x] >= 10 {
				b.flashes++
				b.grid[y][x] = 0
			} else {
				allFlash = false
			}
		}
	}
	return allFlash
}

func dumboOctoSim(filepath string, part2 bool) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	nums := [][]int{}
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
	}

	c := cavern{
		grid:    nums,
		ylen:    len(nums),
		xlen:    len(nums[0]),
		flashes: 0,
		part2:   part2,
	}

	if !part2 {
		for i := 0; i < 100; i++ {
			c.update()
		}

		return c.flashes
	}

	i := 1
	for !c.update() {
		i++
	}
	return i
}
