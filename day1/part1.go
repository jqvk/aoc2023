package day1

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/jqvk/aoc2023/common"
)

func readInput() []string {
	content, err := ioutil.ReadFile("day1/input")
	common.ErrPanic(err)
	scanner := bufio.NewScanner(bytes.NewReader(content))
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func getLineSum(line string) int {
	first := -1
	last := -1
	index := 0

	for true {
		if first >= 0 && last >= 0 {
			break
		}
		if first == -1 && isDigit(line[index]) {
			first = int(line[index] - '0')
		}
		if last == -1 && isDigit(line[len(line)-1-index]) {
			last = int(line[len(line)-1-index] - '0')
		}
    index++
	}

	return first * 10 + last
}

func Part1() {
	content := readInput()
	total := 0
	for _, line := range content {
		sum := getLineSum(line)
		total += sum
	}
  fmt.Println("D1P1:", total)
}
