package main

import (
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

type lineSegment struct {
	x1, y1, x2, y2 int
}

func (ls *lineSegment) swap() {
	ls.y1, ls.y2 = ls.y2, ls.y1
	ls.x1, ls.x2 = ls.x2, ls.x1
}

func initLineSeg(args []int) lineSegment {
	if len(args) != 4 {
		panic("unexpected input")
	}
	lS := lineSegment{
		x1: args[0],
		y1: args[1],
		x2: args[2],
		y2: args[3],
	}

	if lS.y1 > lS.y2 {
		lS.swap()
	}
	if lS.x1 > lS.x2 {
		lS.swap()
	}
	return lS
}

func countOverlaps(filepath string, part2 bool) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	lineSegs := []lineSegment{}
	for input.Scan() {
		fullLine := strings.Split(input.Text(), " -> ")
		newNums := []int{}

		for _, linePart := range fullLine {
			vals := strings.Split(linePart, ",")
			for _, val := range vals {
				newNums = append(newNums, utils.Atoi(val))
			}
		}

		lineSegs = append(lineSegs, initLineSeg(newNums))
	}

	grid := [1005][1005]int{}
	for _, ls := range lineSegs {
		if ls.x1 == ls.x2 {
			for i := ls.y1; i <= ls.y2; i++ {
				grid[ls.x1][i]++
			}
		} else if ls.y1 == ls.y2 {
			for i := ls.x1; i <= ls.x2; i++ {
				grid[i][ls.y1]++
			}
		} else if part2 {
			if (ls.x2 - ls.x1) == (ls.y2 - ls.y1) {
				// slope is 1 (or 45 degrees)
				for i := 0; i <= ls.x2-ls.x1; i++ {
					grid[ls.x1+i][ls.y1+i]++
				}
			} else if (ls.x2 - ls.x1) == -(ls.y2 - ls.y1) {
				// slope is -1 (or -45 degrees)
				for i := 0; i <= ls.x2-ls.x1; i++ {
					grid[ls.x1+i][ls.y1-i]++
				}
			}
		}
	}

	overlaps := 0
	for _, gridLine := range grid {
		for _, num := range gridLine {
			if num >= 2 {
				overlaps++
			}
		}
	}
	return overlaps
}
