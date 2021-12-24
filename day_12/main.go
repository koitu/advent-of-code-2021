package main

import (
	"sort"
	"strings"
	"unicode"

	"github.com/koitu/advent-of-code-2021/utils"
)

func IsLower(s string) bool {
	for _, c := range s {
		if !unicode.IsLower(c) {
			return false
		}
	}
	return true
}

func countPaths(start string, caves map[string][]string, small map[string]bool, visited map[string]bool) int {
	if start == "end" {
		return 1
	}

	// go is not great for dfs
	// although maps are passed by value that value holds a pointer
	visitSave := map[string]bool{}
	for k, v := range visited {
		visitSave[k] = v
	}

	// only mark small caves as visited
	if small[start] {
		visitSave[start] = true
	}

	count := 0
	for _, next := range caves[start] {
		if !visitSave[next] {
			count += countPaths(next, caves, small, visitSave)
		}
	}
	return count
}

func countPathsP2(start string, caves map[string][]string, small map[string]bool, visited map[string]bool, extra bool) int {
	if start == "end" {
		return 1
	}

	visitSave := map[string]bool{}
	for k, v := range visited {
		visitSave[k] = v
	}

	// only mark small caves as visited
	if small[start] {
		visitSave[start] = true
	}

	count := 0
	for _, next := range caves[start] {
		if next == "start" {
			continue
		} else if !visitSave[next] {
			count += countPathsP2(next, caves, small, visitSave, extra)
		} else if extra {
			count += countPathsP2(next, caves, small, visitSave, false)
		}
	}
	return count
}

func passagePaths(filepath string, part2 bool) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}

	caves := map[string][]string{}
	small := map[string]bool{}
	for input.Scan() {
		line := strings.Split(input.Text(), "-")
		from, to := line[0], line[1]

		caves[from] = append(caves[from], to)
		caves[to] = append(caves[to], from)

		small[from] = IsLower(from)
		small[to] = IsLower(to)
	}

	visited := map[string]bool{}
	for i, _ := range caves {
		sort.Strings(caves[i])
	}

	if !part2 {
		return countPaths("start", caves, small, visited)
	}
	return countPathsP2("start", caves, small, visited, true)
}
