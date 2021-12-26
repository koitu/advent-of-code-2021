package utils

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

func LoadFile(filename string) (*bufio.Scanner, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(bytes.NewReader(file)), nil
}

func LoadList(filename string) []int {
	input, err := LoadFile(filename)
	if err != nil {
		panic(err)
	}
	input.Scan()

	args := strings.Split(input.Text(), ",")
	result := []int{}

	for _, str := range args {
		num, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		result = append(result, num)
	}

	return result
}

// func to convert string to int
// and panic if err
