package day6

import (
	"aoc2024/days"
	"aoc2024/utils"
	"fmt"
	"slices"
	"strconv"
)

var input = utils.ReadInputAsBytes(6)
var startingPosition = findStartingPosition()
var positionsVisited map[[2]int]bool
var Solution = days.Day{
	Part1: part1,
	Part2: part2,
}

func part1() string {
	positionsVisited = map[[2]int]bool{}
	dir := [2]int{0, -1}
	pos := startingPosition
	for {
		positionsVisited[pos] = true
		nextPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if nextPos[1] < 0 || nextPos[1] >= len(input) || nextPos[0] < 0 || nextPos[0] >= len(input[0]) {
			break
		}
		if input[nextPos[1]][nextPos[0]] == '#' {
			dir = turnRight(dir)
			continue
		}
		pos = nextPos

	}
	return strconv.Itoa(len(positionsVisited))
}

func part2() string {
	count := 0
	//For each position visited found in part one: Try adding an obstacle there and repeating the whole navigation.
	//If we find ourselves in a position - direction combination we have already seen we know that new obstacle has caused looping
	for obstaclePositionVisited := range positionsVisited {
		positionsVisitedDirs := map[[2]int][][2]int{}
		dir := [2]int{0, -1}
		pos := startingPosition
		stuck := true
		for {
			positionsVisitedDirs[pos] = append(positionsVisitedDirs[pos], dir)
			nextPos := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
			if nextPos[1] < 0 || nextPos[1] >= len(input) || nextPos[0] < 0 || nextPos[0] >= len(input[0]) {
				stuck = false
				break
			}
			if input[nextPos[1]][nextPos[0]] == '#' || nextPos == obstaclePositionVisited {
				dir = turnRight(dir)
				continue
			}
			if slices.Contains(positionsVisitedDirs[nextPos], dir) {
				break
			}
			pos = nextPos

		}
		if stuck {
			count += 1
		}
	}
	return strconv.Itoa(count)
}

func turnRight(startingDir [2]int) [2]int {
	dirs := [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	for i, dir := range dirs {
		if dir == startingDir {
			return dirs[(i+1)%4]
		}
	}
	panic(fmt.Sprintf("Invalid dir (%d,%d)\n", startingDir[0], startingDir[1]))
}

func findStartingPosition() [2]int {
	for y, row := range input {
		for x, symbol := range row {
			if symbol == '^' {
				return [2]int{x, y}
			}
		}
	}
	panic("Starting position not found")
}
