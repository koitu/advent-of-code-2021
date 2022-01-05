package main

import (
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

type coordinate struct {
	x, y, z int
}

var zero = coordinate{0, 0, 0}

type scanner struct {
	loc coordinate
	rel []coordinate
	abs []coordinate
}

func transform(c []coordinate, t func(int, int, int) (int, int, int)) (r []coordinate) {
	for _, a := range c {
		x, y, z := t(a.x, a.y, a.z)
		r = append(r, coordinate{x: x, y: y, z: z})
	}
	return
}

func (a coordinate) minus(b coordinate) coordinate {
	return coordinate{
		x: a.x - b.x,
		y: a.y - b.y,
		z: a.z - b.z,
	}
}

func (s *scanner) checkMatch(check *scanner) bool {
	absBeacons := s.abs

	// move the check scanner to (0,0,0)
	// if there are 12 pairs of beacons offset in the same way then we have our match
	// (this offset is also how the check scanner was moved to (0,0,0))
	for _, t := range transforms {
		offsets := map[coordinate]int{}
		relBeacons := transform(check.rel, t)

		for _, a := range absBeacons {
			for _, r := range relBeacons {
				offsets[r.minus(a)]++
			}
		}

		for offset, times := range offsets {
			if times >= 12 {
				for _, r := range relBeacons {
					check.abs = append(check.abs, r.minus(offset))
				}
				check.loc = zero.minus(offset)
				return true
			}
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func distance(a coordinate, b coordinate) int {
	aa := []int{a.x, a.y, a.z}
	bb := []int{b.x, b.y, b.z}

	sum := 0
	for i := 0; i < 3; i++ {
		if (aa[i] > 0 && bb[i] > 0) || (aa[i] < 0 && bb[i] < 0) {
			aVal, bVal := abs(aa[i]), abs(bb[i])
			if aVal > bVal {
				sum += aVal - bVal
			} else {
				sum += bVal - aVal
			}
		} else {
			sum += abs(aa[i]) + abs(bb[i])
		}
	}

	return sum
}

func beaconScanner(filepath string, part2 bool) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	todo := []*scanner{}
	for {
		if !input.Scan() {
			break
		}
		// ignore the first line and just make a new scanner
		newScanner := &scanner{
			loc: zero,
			rel: []coordinate{},
			abs: nil,
		}

		for input.Scan() {
			s := input.Text()
			if s == "" {
				break
			}

			coords := strings.Split(s, ",")
			newBeacon := coordinate{
				x: utils.Atoi(coords[0]),
				y: utils.Atoi(coords[1]),
				z: utils.Atoi(coords[2]),
			}
			newScanner.rel = append(newScanner.rel, newBeacon)
		}
		todo = append(todo, newScanner)
	}

	todo[0].abs = make([]coordinate, len(todo[0].rel))
	copy(todo[0].abs, todo[0].rel)

	done := []*scanner{todo[0]}
	todo = todo[1:]

	for len(todo) != 0 {
	breakpoint:
		for _, cur := range done {
			for i, s := range todo {
				if cur.checkMatch(s) {
					done = append(done, s)
					todo[i] = todo[len(todo)-1]
					todo = todo[:len(todo)-1]
					break breakpoint
				}
			}
		}
	}

	if !part2 {
		beacons := map[coordinate]bool{}
		for _, s := range done {
			for _, b := range s.abs {
				beacons[b] = true
			}
		}

		return len(beacons)
	}

	maxDist := 0
	for i := 0; i < len(done); i++ {
		for j := i + 1; j < len(done); j++ {
			dist := distance(done[i].loc, done[j].loc)
			if dist > maxDist {
				maxDist = dist
			}
		}
	}

	return maxDist
}

var transforms = []func(int, int, int) (int, int, int){
	func(x, y, z int) (int, int, int) { return x, y, z },
	func(x, y, z int) (int, int, int) { return x, -z, y },
	func(x, y, z int) (int, int, int) { return x, -y, -z },
	func(x, y, z int) (int, int, int) { return x, z, -y },
	func(x, y, z int) (int, int, int) { return -x, z, y },
	func(x, y, z int) (int, int, int) { return -x, -y, z },
	func(x, y, z int) (int, int, int) { return -x, -z, -y },
	func(x, y, z int) (int, int, int) { return -x, y, -z },

	func(x, y, z int) (int, int, int) { return z, x, y },
	func(x, y, z int) (int, int, int) { return z, -y, x },
	func(x, y, z int) (int, int, int) { return z, -x, -y },
	func(x, y, z int) (int, int, int) { return z, y, -x },
	func(x, y, z int) (int, int, int) { return -z, y, x },
	func(x, y, z int) (int, int, int) { return -z, -x, y },
	func(x, y, z int) (int, int, int) { return -z, -y, -x },
	func(x, y, z int) (int, int, int) { return -z, x, -y },

	func(x, y, z int) (int, int, int) { return y, z, x },
	func(x, y, z int) (int, int, int) { return y, -x, z },
	func(x, y, z int) (int, int, int) { return y, -z, -x },
	func(x, y, z int) (int, int, int) { return y, x, -z },
	func(x, y, z int) (int, int, int) { return -y, x, z },
	func(x, y, z int) (int, int, int) { return -y, -z, x },
	func(x, y, z int) (int, int, int) { return -y, -x, -z },
	func(x, y, z int) (int, int, int) { return -y, z, -x },
}
