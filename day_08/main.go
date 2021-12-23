package main

import (
	"sort"
	"strings"

	"github.com/koitu/advent-of-code-2021/utils"
)

// this is quite awful but it works
// I might try to fix it later (lol)

// it might have been simpler if I just did brute force

func sortStr(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// if haveAll is true: return the first string in supStr to contain
// all the char in subStr
// if haveAll is false: return the first one that does not contains all
// the char in subStr
func firstStr(supStrs []string, subStr string, haveAll bool) string {
	match := []rune(subStr)

	for _, str := range supStrs {
		found := []bool{}
		for i := 0; i < len(match); i++ {
			found = append(found, false)
		}

		for _, char := range str {
			for i := range match {
				if char == match[i] {
					found[i] = true
				}
			}
		}

		foundAll := true
		for _, a := range found {
			if !a {
				foundAll = false
			}
		}

		if foundAll == haveAll {
			return str
		}
	}
	return ""
}

func removeFromSlice(s []string, str string) []string {
	for i, match := range s {
		if match == str {
			s[i] = s[len(s)-1]
			s[len(s)-1] = ""
			return s[:len(s)-1]
		}
	}
	return s
}

// return the first rune in str2 that is not in str1
func firstDiff(str1, str2 string) rune {
	str1 = sortStr(str1)
	str2 = sortStr(str2)

	for _, s2 := range str2 {
		found := false
		for _, s1 := range str1 {
			if s1 == s2 {
				found = true
				break
			}
		}

		if !found {
			return s2
		}
	}
	return ' '
}

// could make this better by using bitmap
// a -> 1
// b -> 2
// c -> 4
// etc.
func sevenSegmentMatches(filepath string, part2 bool) int {
	input, err := utils.LoadFile(filepath)
	if err != nil {
		panic(err)
	}
	unique := 0
	sum := 0

	for input.Scan() {
		found := [10]string{}
		find := []string{}
		len5 := []string{}
		len6 := []string{}
		line := strings.Split(input.Text(), "|")

		for _, str := range strings.Fields(line[1]) {
			strlen := len(str)
			if !part2 {
				if strlen == 2 || strlen == 3 || strlen == 4 || strlen == 7 {
					unique++
				}
			} else {
				find = append(find, sortStr(str))
			}
		}

		for _, str := range strings.Fields(line[0]) {
			strlen := len(str)
			if part2 {
				strSorted := sortStr(str)
				switch strlen {
				case 2:
					found[1] = strSorted
				case 3:
					found[7] = strSorted
				case 4:
					found[4] = strSorted
				case 5:
					// 2,3,5 have 5 segments
					len5 = append(len5, strSorted)
				case 6:
					// 0,6,9 have 6 segments
					len6 = append(len6, strSorted)
				case 7:
					found[8] = strSorted
				}
			}
		}

		if part2 {
			top := firstDiff(found[1], found[7])
			// 3 contains all segments that make up 1 (unlike 2 and 5)
			found[3] = firstStr(len5, found[1], true)
			len5 = removeFromSlice(len5, found[3])
			// 6 does not contain some segments that make up 1 (unlike 0 and 9)
			found[6] = firstStr(len6, found[1], false)
			len6 = removeFromSlice(len6, found[6])

			//        top
			//       +---+
			//  ltop |   | rtop
			//       +---+ <---- mid
			//  lbot |   | rbot
			//       +---+
			//        bot
			rtop := firstDiff(found[6], found[1])
			rbot := firstDiff(string(rtop), found[1])
			bot := firstDiff(found[4]+found[7], found[3])
			mid := firstDiff(string([]rune{top, rtop, rbot, bot}), found[3])

			// 0 does not have a mid seg (and 9 does)
			found[9] = firstStr(len6, string(mid), true)
			len6 = removeFromSlice(len6, found[9])
			found[0] = len6[0]

			// 5 does not have a rtop seg (and 2 does)
			found[2] = firstStr(len5, string(rtop), true)
			len5 = removeFromSlice(len5, found[2])
			found[5] = len5[0]

			re := 0
			multi := 1
			for i := 3; i >= 0; i-- {
				for j := 0; j < 10; j++ {
					if find[i] == found[j] {
						re += j * multi
					}
				}
				multi *= 10
			}
			sum += re
		}
	}
	if !part2 {
		return unique
	}
	return sum
}
