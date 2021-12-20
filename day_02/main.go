package main

import (
	"strconv"
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

type Position struct {
	hori, depth, aim int
}

func (pos *Position) update(command string, units int) {
	switch command {
	case "forward":
		pos.hori += units
	case "down":
		pos.depth += units
	case "up":
		pos.depth -= units
	}

}

func (pos *Position) updateWithAim(command string, units int) {
	switch command {
	case "down":
		pos.aim += units
	case "up":
		pos.aim -= units
	case "forward":
		pos.hori += units
		pos.depth += pos.aim * units
	}
}

// https://adventofcode.com/2021/day/2
func getSubPosition(filepath string, withAim bool) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	pos := Position{0, 0, 0}

	for input.Scan() {
		s := strings.Fields(input.Text())
		command := s[0]
		units, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		if !withAim {
			pos.update(command, units)
		} else {
			pos.updateWithAim(command, units)
		}

	}

	return pos.hori * pos.depth
}
