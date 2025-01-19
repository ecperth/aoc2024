package day5

import (
	"aoc2024/days"
	"aoc2024/utils"
	"slices"
	"strconv"
	"strings"
)

var input = utils.ReadInputAsStrings(5)
var rules, updates = readRulesAndUpdates()

var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {
	sum := 0
	for _, update := range updates {
		correct := true
		for i := 0; i < len(update)-1; i++ {
			for j := i + 1; j < len(update); j++ {
				if slices.Contains(rules[update[j]], update[i]) {
					correct = false
					break
				}
			}
			if !correct {
				break
			}
		}
		if correct {
			sum += update[(len(update)-1)/2]
		}
	}
	return strconv.Itoa(sum)
}

func part2() string {
	sum := 0
	for _, update := range updates {
		updated := false
		updateCopy := make([]int, len(update))
		copy(updateCopy, update)
		for i := 0; i < len(update)-1; i++ {
			for j := i + 1; j < len(update); j++ {
				if slices.Contains(rules[updateCopy[j]], updateCopy[i]) {
					updated = true
					uj := updateCopy[j]
					updateCopy[j] = updateCopy[i]
					updateCopy[i] = uj
				}
			}
		}
		if updated {
			sum += updateCopy[(len(update)-1)/2]
		}
	}
	return strconv.Itoa(sum)
}

func readRulesAndUpdates() (rules map[int][]int, updates [][]int) {
	rules = make(map[int][]int)
	readingRules := true
	for _, line := range input {
		if len(line) == 0 {
			readingRules = false
			continue
		}
		if readingRules {
			var r1, r2 int
			utils.SscanfUnsafe(line, "%d|%d", &r1, &r2)
			rules[r1] = append(rules[r1], r2)
		} else {
			updates = append(updates, utils.StringSliceToIntSlice(strings.Split(line, ",")))
		}
	}
	return rules, updates
}
