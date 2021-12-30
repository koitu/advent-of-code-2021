package main

import (
	"math"
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

func polymerization(filepath string, turns int) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	input.Scan()
	template := strings.Split(input.Text(), "")
	input.Scan()

	count := map[string]int{}
	state := map[string]int{}
	last := ""
	for _, cur := range template {
		count[cur]++
		if last != "" {
			state[last+cur]++
		}
		last = cur
	}

	rules := map[string]string{}
	for input.Scan() {
		line := strings.Split(input.Text(), " -> ")
		rules[line[0]] = line[1]
	}

	for i := 0; i < turns; i++ {
		newState := map[string]int{}

		// rules["HC"] = "B"
		// count: every "HC" will create a "B"
		// state: every "HC" becomes "HB" and "BC"
		for str, val := range state {
			s := strings.Split(str, "")
			rule := rules[str]
			newState[s[0]+rule] += val
			newState[rule+s[1]] += val
			count[rule] += val
		}

		state = newState
	}

	high := 0
	low := math.MaxInt

	for _, val := range count {
		if val > high {
			high = val
		}
		if val < low {
			low = val
		}
	}

	return high - low
}
