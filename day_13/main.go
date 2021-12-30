package main

import (
	"fmt"
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

type coordinate struct {
	x, y int
}

// dots will never appear on the fold line

//   1     4     7
// . # . . | . . # .
//  x=4-3       x=4+3
//   =4-(7-4)
//   =2*4-7

func foldPaper(filepath string, part2 bool) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	// consume the coordinates into a slice
	points := []coordinate{}
	for input.Scan() {
		in := input.Text()
		if in == "" {
			break
		}

		line := strings.Split(in, ",")
		points = append(points, coordinate{x: utils.Atoi(line[0]), y: utils.Atoi(line[1])})
	}

	// process the instructions
	for input.Scan() {
		instruct := strings.Split(input.Text(), "=")
		fold := utils.Atoi(instruct[1])

		if instruct[0] == "fold along y" {
			for i, coor := range points {
				if coor.y > fold {
					points[i].y = fold*2 - coor.y
				}
			}
		} else if instruct[0] == "fold along x" {
			for i, coor := range points {
				if coor.x > fold {
					points[i].x = fold*2 - coor.x
				}
			}
		}

		if !part2 {
			// if part1 then process the first instruction
			break
		}
	}

	// check points and remove duplicates
	maxy := 0
	maxx := 0
	for _, point := range points {
		if point.y < 0 || point.x < 0 {
			panic(fmt.Sprintf("point (%d,%d) is folded outside the paper", point.x, point.y))
		}
		if point.y > maxy {
			maxy = point.y
		}
		if point.x > maxx {
			maxx = point.x
		}
	}

	// [x][y]
	maxy++
	maxx++
	visable := make([][]bool, maxx)
	for i := 0; i < maxx; i++ {
		visable[i] = make([]bool, maxy)
	}

	for _, point := range points {
		visable[point.x][point.y] = true
	}

	count := 0
	if !part2 {
		// count the visable parts
		for _, col := range visable {
			for _, vis := range col {
				if vis {
					count++
				}
			}
		}
	} else {
		// print the visable parts
		for i := 0; i < maxy; i++ {
			for j := 0; j < maxx; j++ {
				if visable[j][i] {
					fmt.Printf("# ")
				} else {
					fmt.Printf(". ")
				}
			}
			fmt.Printf("\n")
		}
	}

	return count
}
