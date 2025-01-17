package day1

import (
	"aoc2024/days"
	"aoc2024/utils"
	"sort"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(1)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {

	var leftList []int
	var rightList []int
	for _, line := range input {
		splitLine := strings.Fields(line)
		leftList = append(leftList, utils.AtoiUnsafe(splitLine[0]))
		rightList = append(rightList, utils.AtoiUnsafe(splitLine[1]))
	}
	sort.Ints(leftList)
	sort.Ints(rightList)

	sum := 0
	for n, leftValue := range leftList {
		sum += utils.Abs(leftValue - rightList[n])
	}
	return strconv.Itoa(sum)
}

func part2() string {

	var leftList []int
	var rightList []int
	for _, line := range input {
		splitLine := strings.Fields(line)
		leftList = append(leftList, utils.AtoiUnsafe(splitLine[0]))
		rightList = append(rightList, utils.AtoiUnsafe(splitLine[1]))
	}
	sort.Ints(leftList)
	sort.Ints(rightList)

	rightPointer := 0
	sum := 0
	for _, leftValue := range leftList {
		count := 0
		for {
			if leftValue < rightList[rightPointer] || rightPointer == len(rightList)-1 {
				break
			} else if leftValue == rightList[rightPointer] {
				count += 1
				rightPointer += 1
			} else if leftValue > rightList[rightPointer] {
				rightPointer += 1
			}
		}
		sum += count * leftValue
	}
	return strconv.Itoa(sum)
}
