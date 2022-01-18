package main

import (
	"github.com/koitu/advent-of-code-2021/utils"
)

// if the image enhancement algorithm starts with a #
// make sure it end with a . (when this does occur alernate between the out of bounds returning # and .)
// make a special case for when it does occur (the test does not have this)

// afterwards just make sure the grid grows so that is there is always a boundary of 3 units from out of bounds

type image struct {
	pic     [][]bool
	x, y    int
	boundBy bool
}

func toInt(b []bool) int {

	return 0
}

func (pic *image) at(y, x int) bool {

	return true
}

func (pic *image) resize() {

}

func (pic *image) enhance(alg []bool) {

}

func mapEnhance(filepath string, iterations int) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	alg := []bool{}
	pic := image{
		pic: [][]bool{},
		// the true size (should only be updated every 2 enhances)
		x: 0,
		y: 0,
		// value returned by attempting to access out of bounds
		boundBy: false,
	}

	input.Scan()
	for _, v := range input.Text() {
		if v == '.' {
			alg = append(alg, false)
		} else if v == '#' {
			alg = append(alg, true)
		} else {
			panic("unexpected rune while scanning algorithm")
		}
	}
	input.Scan()

	for input.Scan() {
		newLn := []bool{}
		for _, v := range input.Text() {
			if v == '.' {
				newLn = append(newLn, false)
			} else if v == '#' {
				newLn = append(newLn, true)
			} else {
				panic("unexpected rune while scanning line")
			}
		}
		pic.pic = append(pic.pic, newLn)
		pic.x = len(newLn)
	}
	pic.y = len(pic.pic)

	if pic.pic[0][0] && pic.pic[pic.y-1][pic.x-1] {
		panic("infinite")
	}

	pic.resize()
	for i := 0; i < iterations; i++ {
		pic.enhance(alg)
		if alg[0] && !alg[len(alg)-1] {
			pic.boundBy = !pic.boundBy
		}
		if i%2 == 1 {
			pic.resize()
		}
	}

	return 0
}
