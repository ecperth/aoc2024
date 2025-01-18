package day3

import (
	"aoc2024/days"
	"aoc2024/utils"
	"bytes"
	"regexp"
	"strconv"
)

var input = utils.ReadInputAsBytes(3)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {
	sum := 0
	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	for _, line := range input {
		tokens := r.FindAll(line, -1)
		var m1, m2 int
		for _, token := range tokens {
			utils.SscanfUnsafe(string(token), "mul(%d,%d)", &m1, &m2)
			sum += m1 * m2
		}
	}
	return strconv.Itoa(sum)
}

func part2() string {
	sum := 0
	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)|do\\(\\)|don't\\(\\)")
	do := true
	doToken, dontToken := []byte("do()"), []byte("don't()")
	for _, line := range input {
		tokens := r.FindAll(line, -1)
		var m1, m2 int
		for _, token := range tokens {
			if bytes.Equal(token, doToken) {
				do = true
			} else if bytes.Equal(token, dontToken) {
				do = false
			} else if do {
				utils.SscanfUnsafe(string(token), "mul(%d,%d)", &m1, &m2)
				sum += m1 * m2
			}
		}
	}
	return strconv.Itoa(sum)
}
