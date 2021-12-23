package main

import (
	"sort"

	"github.com/koitu/advent-of-code-2021/utils"
)

var match = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var scoreP1 = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var scoreP2 = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func opens(c rune) bool {
	return c == '(' || c == '[' || c == '{' || c == '<'
}

type queue struct {
	q []rune
}

func (q *queue) len() int {
	return len(q.q)
}

func (q *queue) push(c rune) {
	q.q = append(q.q, c)
}

func (q *queue) pop() rune {
	qlen := q.len()
	if qlen == 0 {
		panic("cannot pop, queue is of length 0")
	}
	p := q.q[qlen-1]
	q.q = q.q[:qlen-1]
	return p
}

func scoreSyntax(filepath string, part2 bool) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	score1 := 0
	score2 := []int{}
	for input.Scan() {
		q := queue{
			q: []rune{},
		}
		line := input.Text()
		corrupted := false

		for _, val := range line {
			if opens(val) {
				q.push(val)
				continue
			}

			p := q.pop()
			if match[p] != val {
				if !part2 {
					score1 += scoreP1[val]
				}
				corrupted = true
				break
			}
		}

		if part2 && !corrupted {
			cur := 0
			for q.len() != 0 {
				cur *= 5
				cur += scoreP2[match[q.pop()]]
			}
			score2 = append(score2, cur)
		}
	}

	if !part2 {
		return score1
	}
	sort.Ints(score2)
	return score2[len(score2)/2]
}
