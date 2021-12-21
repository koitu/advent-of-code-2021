package main

import (
	"strconv"
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

func lanternFish(filepath string, days int) uint64 {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}
	input.Scan()

	fishes := strings.Split(input.Text(), ",")
	pops := [10]uint64{}
	for _, fish := range fishes {
		n, err := strconv.Atoi(fish)
		if err != nil || n < 0 || n > 9 {
			panic("invalid input")
		}
		pops[n]++
	}

	for i := 0; i < days; i++ {
		new := pops[0]
		for j := 0; j < 8; j++ {
			pops[j] = pops[j+1]
		}
		pops[8] = new
		pops[6] += new
	}

	var sum uint64
	sum = 0
	for _, pop := range pops {
		sum += pop
	}
	return sum
}
