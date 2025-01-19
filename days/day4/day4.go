package day4

import (
	"aoc2024/days"
	"aoc2024/utils"
	"bytes"
	"strconv"
)

var input = utils.ReadInputAsBytes(4)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {
	count := 0
	for y, row := range input {
		for x, value := range row {
			if value == 'X' {
				for xDir := -1; xDir <= 1; xDir++ {
					for yDir := -1; yDir <= 1; yDir++ {
						//check if we would hit a boundary in this direction before the end of the word
						if y+3*yDir < 0 || y+3*yDir >= len(input) || x+3*xDir < 0 || x+3*xDir >= len(row) {
							continue
						}
						if input[y+yDir][x+xDir] == 'M' && input[y+2*yDir][x+2*xDir] == 'A' && input[y+3*yDir][x+3*xDir] == 'S' {
							count += 1
						}
					}
				}
			}
		}
	}
	return strconv.Itoa(count)
}

func part2() string {
	count := 0
	dCompare := [2][]byte{{'M', 'S'}, {'S', 'M'}}
	for y := 1; y < len(input)-1; y++ {
		for x := 1; x < len(input[y])-1; x++ {
			if input[y][x] == 'A' {
				d1 := []byte{input[y-1][x-1], input[y+1][x+1]}
				d2 := []byte{input[y+1][x-1], input[y-1][x+1]}
				if (bytes.Equal(d1, dCompare[0]) || bytes.Equal(d1, dCompare[1])) && (bytes.Equal(d2, dCompare[0]) || bytes.Equal(d2, dCompare[1])) {
					count += 1
				}
			}
		}
	}
	return strconv.Itoa(count)
}
