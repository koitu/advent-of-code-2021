package main

import (
	"fmt"
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

type node struct {
	value               int
	left, right, parent *node
}

//lint:ignore U1000 method is useful for debugging (converts nodes back into original string)
func (n *node) display() {
	if n.left == nil && n.right == nil {
		fmt.Printf("%d", n.value)
		return
	}

	if n.left != nil {
		fmt.Printf("[")
		n.left.display()
	}

	if n.right != nil {
		fmt.Printf(",")
		n.right.display()
		fmt.Printf("]")
	}
}

// explode => replace a pair with 0
// pair's left value is added to the first number to the left of the pair
// pair's right value is added to the first number to the right of the pair
// exploding pairs will always have two numbers
func (n *node) explode() {
	p := n.parent
	last := n
	for p != nil {
		if p.left != last {
			p = p.left
			for p.value == -1 {
				p = p.right
			}
			p.value += n.left.value
			break
		}
		last = p
		p = p.parent
	}

	p = n.parent
	last = n
	for p != nil {
		if p.right != last {
			p = p.right
			for p.value == -1 {
				p = p.left
			}
			p.value += n.right.value
			break
		}
		last = p
		p = p.parent
	}

	n.value = 0
	n.left = nil
	n.right = nil
}

// split => replace a number with a pair
// pair's left should be number divided by two and rounded down
// pair's right should be number divided by two and rounded up
// sum of new pair's left and right should equal number
func (n *node) split() {
	n.left = newNode(n, n.value/2)
	n.right = newNode(n, (n.value+1)/2)

	n.value = -1
}

func (n *node) magnitude() int {
	if n.value != -1 {
		return n.value
	}

	return 3*n.left.magnitude() + 2*n.right.magnitude()
}

func (n *node) reduce(depth int, split bool) bool {
	if depth >= 5 && n.value == -1 {
		n.explode()
		return true
	}

	if split && n.value >= 10 {
		n.split()
		return true
	}

	if n.left != nil {
		if n.left.reduce(depth+1, split) {
			return true
		}
	}

	if n.right != nil {
		if n.right.reduce(depth+1, split) {
			return true
		}
	}
	return false
}

func (n *node) copy(parent *node) *node {
	new := &node{
		value:  n.value,
		left:   nil,
		right:  nil,
		parent: parent,
	}
	if n.left != nil {
		new.left = n.left.copy(new)
	}
	if n.right != nil {
		new.right = n.right.copy(new)
	}

	return new
}

func newNode(parent *node, value int) *node {
	return &node{
		value:  value,
		left:   nil,
		right:  nil,
		parent: parent,
	}
}

func parseSnailFish(s string) *node {
	sf := &node{
		value:  -1,
		left:   nil,
		right:  nil,
		parent: nil,
	}
	ss := strings.Split(s, "")
	ss = ss[1:]

	num := ""
	cur := sf

	for _, r := range ss {
		switch r {
		case "[":
			if cur.left == nil {
				cur.left = newNode(cur, -1)
				cur = cur.left
			} else {
				cur.right = newNode(cur, -1)
				cur = cur.right
			}

		case "]":
			if num != "" {
				cur.right = newNode(cur, utils.Atoi(num))
				num = ""
			}
			cur = cur.parent

		case ",":
			if num != "" {
				cur.left = newNode(cur, utils.Atoi(num))
				num = ""
			}

		default:
			num += r
		}
	}
	return sf
}

func snailFishSum(nums []*node, order []int) int {
	sum := nums[order[0]].copy(nil)
	for i := 1; i < len(order); i++ {
		sum = &node{
			value:  -1,
			left:   sum,
			right:  nums[order[i]].copy(nil),
			parent: nil,
		}
		if sum.left != nil {
			sum.left.parent = sum
		}
		if sum.right != nil {
			sum.right.parent = sum
		}
		for {
			for {
				// exhaust the explodes before moving on the splits
				if !sum.reduce(1, false) {
					break
				}
			}
			if !sum.reduce(1, true) {
				break
			}
		}
	}

	return sum.magnitude()
}

// might be useful for later

// func heapPerm(l []int) (perms [][]int) {
// 	var run func(a []int, k int)
// 	run = func(a []int, k int) {
// 		if k == len(a) {
// 			perms = append(perms, a)
// 		} else {
// 			for i := k; i < len(l); i++ {
// 				a[k], a[i] = a[i], a[k]
// 				run(a, k+1)
// 				a[k], a[i] = a[i], a[k]
// 			}
// 		}
// 	}
// 	run(l, 0)

// 	return perms
// }

func snailFish(filepath string, part2 bool) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	nums := []*node{}
	for input.Scan() {
		nums = append(nums, parseSnailFish(input.Text()))
	}

	if !part2 {
		order := []int{}
		for i := range nums {
			order = append(order, i)
		}
		return snailFishSum(nums, order)

	}

	max := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			new := snailFishSum(nums, []int{i, j})
			if new > max {
				max = new
			}
		}
	}

	return max
}
