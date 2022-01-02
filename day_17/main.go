package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

// y will always be accelerating down (y target is negative)
func yHits(vel, step, maxStep, pos, minPos, maxPos int) (hits []int) {
	if pos < minPos || step > maxStep {
		return
	}

	if pos >= minPos && pos <= maxPos {
		hits = append(hits, step)
	}

	hits = append(hits, yHits(vel-1, step+1, maxStep, pos+vel, minPos, maxPos)...)

	return
}

// x will slow down until vel is 0 (x target is positive)
func xHits(vel, step, maxStep, pos, minPos, maxPos int) (hits []int) {
	if pos > maxPos || step > maxStep {
		return
	}

	if pos >= minPos && pos <= maxPos {
		hits = append(hits, step)
	}

	newVel := vel - 1
	if vel < 0 {
		vel = 0
	}
	hits = append(hits, xHits(newVel, step+1, maxStep, pos+vel, minPos, maxPos)...)

	return
}

type pair struct {
	x, y int
}

func trickShot(filepath string, part2 bool) int {
	// target area: x=20..30, y=-10..-5
	f, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	// 20..30..-10..-5
	preParse := strings.Replace(string(f[15:]), ", y=", "..", 1)
	parse := strings.Split(preParse, "..")
	vals := []int{}
	for _, s := range parse {
		vals = append(vals, utils.Atoi(s))
	}

	xMin, xMax, yMin, yMax := vals[0], vals[1], vals[2], vals[3]

	if yMin >= 0 || yMax >= 0 || xMin > xMax || yMin > yMax {
		panic(fmt.Errorf("unexpected input: %v", vals))
	}

	yMaxVel := -yMin - 1
	if !part2 {
		return (yMaxVel * (yMaxVel + 1)) / 2
	}

	if xMin < 0 && xMax < 0 {
		xMin, xMax = -xMax, -xMin
	}

	// the possible ways to get within the target area in one step
	oneStep := (xMax - xMin + 1) * (yMax - yMin + 1)
	maxSteps := (yMaxVel + 1) * 2
	xVelMax := (xMax + 1) / 2

	yReqSteps := map[int][]int{}
	for vel := yMax + 1; vel <= yMaxVel; vel++ {
		for _, step := range yHits(vel, 0, maxSteps, 0, yMin, yMax) {
			yReqSteps[step] = append(yReqSteps[step], vel)
		}
	}

	xReqSteps := map[int][]int{}
	for vel := 0; vel <= xVelMax; vel++ {
		for _, step := range xHits(vel, 0, maxSteps, 0, xMin, xMax) {
			xReqSteps[step] = append(xReqSteps[step], vel)
		}
	}

	// I'm sure there is a better way to find the num of distinct pairs
	// but this is what I got for now
	pairs := map[pair]bool{}
	for step, xVels := range xReqSteps {
		for _, xVel := range xVels {
			for _, yVel := range yReqSteps[step] {
				pairs[pair{x: xVel, y: yVel}] = true
			}
		}
	}

	// find every sequence of numbers that add to within xMin to xMax
	return oneStep + len(pairs)
}
