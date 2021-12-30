package main

import (
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

type Board struct {
	won    bool
	board  [5][5]int
	marked [5][5]bool
}

func initBoard() *Board {
	return &Board{
		won:    false,
		board:  [5][5]int{},
		marked: [5][5]bool{},
	}
}

func (b *Board) testWin() {
	for i := 0; i < 5; i++ {
		p1won := true
		p2won := true
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				p1won = false
			}
			if !b.marked[j][i] {
				p2won = false
			}
		}
		if p1won || p2won {
			b.won = true
			return
		}
	}
}

func (b *Board) update(num int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.board[i][j] == num {
				b.marked[i][j] = true
			}
		}
	}
	b.testWin()
}

func (b *Board) calcScore(num int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.marked[i][j] {
				sum += b.board[i][j]
			}
		}
	}
	return sum * num
}

func bingoSubsystem(filepath string, part2 bool) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	input.Scan()
	draws := strings.Split(input.Text(), ",")
	var nums []int

	for _, draw := range draws {
		nums = append(nums, utils.Atoi(draw))
	}

	var boards []*Board

	for input.Scan() {
		newBoard := initBoard()
		for i := 0; i < 5; i++ {
			if input.Scan() {
				line := strings.Fields(input.Text())
				for pos, val := range line {
					newBoard.board[i][pos] = utils.Atoi(val)
				}
			}
		}
		boards = append(boards, newBoard)
	}

	if !part2 {
		for _, num := range nums {
			for _, board := range boards {
				board.update(num)
				if board.won {
					return board.calcScore(num)
				}
			}
		}
	}

	boardsWon := 0
	boardsNum := len(boards)
	for _, num := range nums {
		for _, board := range boards {
			if board.won {
				continue
			}

			board.update(num)
			if board.won {
				boardsWon++
			}
			if boardsWon == boardsNum {
				return board.calcScore(num)
			}
		}
	}

	// if there are no moves left but
	return -1
}
