package main

import (
	"math"

	"github.com/koitu/advent-of-code-2021/utils"
)

func alignCrabCost(filepath string, part2 bool) int {
	crabs := utils.LoadList(filepath)
	maxCrab := 0
	for _, n := range crabs {
		if n > maxCrab {
			maxCrab = n
		}
	}

	minFuel := math.MaxInt

	for loc := 0; loc < maxCrab; loc++ {
		fuel := 0
		for _, pos := range crabs {
			cost := int(math.Abs(float64(loc - pos)))

			if !part2 {
				fuel += cost
			} else {
				fuel += (cost * (cost + 1)) / 2
			}
		}
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	return minFuel
}
