package main

import (
	"container/heap"
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

type Node struct {
	y, x, weight int
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].weight < pq[j].weight }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*Node)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil // activate gc
	*pq = old[0 : n-1]
	return node
}

func bestPath(risk [][]int, endy, endx int) int {
	visited := [][]bool{}
	for i := range risk {
		visited = append(visited, make([]bool, len(risk[i])))
	}

	pq := make(PriorityQueue, 1)
	pq[0] = &Node{
		y:      0,
		x:      0,
		weight: 0,
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		n := heap.Pop(&pq).(*Node)
		y := n.y
		x := n.x
		weight := n.weight

		if visited[y][x] {
			continue
		}

		if y == endy && x == endx {
			return weight
		}

		if y-1 >= 0 {
			heap.Push(&pq, &Node{y: y - 1, x: x, weight: weight + risk[y-1][x]})
		}
		if y+1 <= endy {
			heap.Push(&pq, &Node{y: y + 1, x: x, weight: weight + risk[y+1][x]})
		}
		if x-1 >= 0 {
			heap.Push(&pq, &Node{y: y, x: x - 1, weight: weight + risk[y][x-1]})
		}
		if x+1 <= endx {
			heap.Push(&pq, &Node{y: y, x: x + 1, weight: weight + risk[y][x+1]})
		}

		visited[y][x] = true
	}
	return -1
}

func addWithWrap(i, j int) (result int) {
	result = i + j
	if result > 9 {
		result -= 9
	}
	return
}

func traverseCave(filepath string, part2 bool) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	cave := [][]int{}
	for input.Scan() {
		newRow := []int{}
		line := strings.Split(input.Text(), "")
		for _, val := range line {
			newRow = append(newRow, utils.Atoi(val))
		}

		cave = append(cave, newRow)
	}

	if part2 {
		// 25 times larger (~0.25s for input.txt)
		bigCave := [][]int{}
		for _, row := range cave {
			newRow := []int{}
			for i := 0; i < 5; i++ {
				for _, v := range row {
					newRow = append(newRow, addWithWrap(i, v))
				}
			}
			bigCave = append(bigCave, newRow)
		}

		for i := 1; i < 5; i++ {
			for j := 0; j < len(cave); j++ {
				newRow := []int{}
				for _, v := range bigCave[j] {
					newRow = append(newRow, addWithWrap(i, v))
				}
				bigCave = append(bigCave, newRow)
			}
		}

		return bestPath(bigCave, len(bigCave)-1, len(bigCave[0])-1)
	}

	return bestPath(cave, len(cave)-1, len(cave[0])-1)
}
