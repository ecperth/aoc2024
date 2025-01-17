package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func ReadInputAsStrings(day int) []string {
	s := getScanner(day)

	var lines []string
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func ReadInputAsBytes(day int) [][]byte {
	s := getScanner(day)

	var matrix [][]byte
	for s.Scan() {
		matrix = append(matrix, []byte(s.Text()))
	}
	return matrix
}

func getScanner(day int) *bufio.Scanner {
	file, err := os.Open(fmt.Sprintf("./inputs/day%d", day))
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

func ClearTerminal() {
	//linux
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func AtoiUnsafe(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return i
}

func SscanfUnsafe(input string, format string, args ...interface{}) {
	_, err := fmt.Sscanf(input, format, args...)
	if err != nil {
		panic(err)
	}
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign[T Number](x T) int {
	if x == 0 {
		return 0
	} else if x > 0 {
		return 1
	} else {
		return -1
	}
}

func Min[T Number](x1, x2 T) T {
	if x1 > x2 {
		return x2
	}
	return x1
}

func Max[T Number](x1, x2 T) T {
	if x1 > x2 {
		return x1
	}
	return x2
}

func Contains[kT comparable, vT any](m map[kT]vT, k kT) (contains bool) {
	_, contains = m[k]
	return
}

func SortMapKeysByValue[kT comparable, vT int](m map[kT]vT) []kT {
	keys := make([]kT, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] < m[keys[j]]
	})

	return keys
}

func StringSliceToIntSlice(input []string) []int {
	result := make([]int, 0, len(input))
	for _, s := range input {
		i, _ := strconv.Atoi(s)
		result = append(result, i)
	}

	return result
}

/*
StringSplitConsecutive
Split a string on a separate including consecutive occurrences of that separator.
StringSplitConsecutive("foobar", 'o') ==
StringSplitConsecutive("foobar", 'o') ==
["f", "bar"]
*/
func StringSplitConsecutive(input string, separator rune) []string {
	return strings.FieldsFunc(input, func(r rune) bool {
		return r == separator
	})
}
