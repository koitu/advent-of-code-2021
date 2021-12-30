package utils

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
)

// load the given file and return a scanner to it
func LoadFile(filename string) (*bufio.Scanner, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(bytes.NewReader(file)), nil
}

// convert a list of comma separated values into a int slice
func LoadList(filename string) []int {
	input, err := LoadFile(filename)
	if err != nil {
		panic(err)
	}
	input.Scan()

	args := strings.Split(input.Text(), ",")
	result := []int{}

	for _, str := range args {
		result = append(result, Atoi(str))
	}

	return result
}

// convert string to int and panic on error
func Atoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return val
}
