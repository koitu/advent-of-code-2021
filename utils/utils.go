package utils

import (
	"bufio"
	"bytes"
	"os"
)

func LoadFile(filename string) (*bufio.Scanner, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(bytes.NewReader(file)), nil
}
