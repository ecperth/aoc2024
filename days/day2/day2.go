package day2

import (
	"aoc2024/days"
	"aoc2024/utils"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(2)
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {
	safeCount := 0
	for _, line := range input {
		report := utils.StringSliceToIntSlice(strings.Fields(line))

		safe, _ := isSafeReport(report)
		if safe {
			safeCount += 1
		}
	}
	return strconv.Itoa(safeCount)
}

func part2() string {
	safeCount := 0
	for _, line := range input {
		report := utils.StringSliceToIntSlice(strings.Fields(line))

		safe, unsafeIndex := isSafeReport(report)
		if safe {
			safeCount += 1
			continue
		}
		//try removing left in failing pair
		amendedReport := append([]int(nil), report[:unsafeIndex]...)
		amendedReport = append(amendedReport, report[unsafeIndex+1:]...)
		safe, _ = isSafeReport(amendedReport)
		if safe {
			safeCount += 1
			continue
		}
		//try removing right in failing pair
		amendedReport = append([]int(nil), report[:unsafeIndex+1]...)
		amendedReport = append(amendedReport, report[unsafeIndex+2:]...)
		safe, _ = isSafeReport(amendedReport)
		if safe {
			safeCount += 1
			continue
		}
		//special case. If second pair fails try removing first number
		if unsafeIndex == 1 {
			amendedReport = append([]int(nil), report[:0]...)
			amendedReport = append(amendedReport, report[1:]...)
			safe, _ = isSafeReport(amendedReport)
			if safe {
				safeCount += 1
				continue
			}
		}
	}
	return strconv.Itoa(safeCount)
}

func isSafeReport(report []int) (bool, int) {
	safe := true
	dir := 0
	for {
		for i := range len(report) - 1 {
			diff := report[i] - report[i+1]
			if diff >= 1 && diff <= 3 {
				if dir == 0 {
					dir = 1
				} else if dir == -1 {
					safe = false
				}
			} else if diff >= -3 && diff <= -1 {
				if dir == 0 {
					dir = -1
				} else if dir == 1 {
					safe = false
				}
			} else {
				safe = false
			}
			if !safe {
				return false, i
			}
		}
		return true, -1
	}
}
