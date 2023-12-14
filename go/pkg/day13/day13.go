package day13

import (
	"bufio"
	"strconv"
	"strings"
)

type Map [][]rune

func (m Map) Rows() []string {
	rows := make([]string, len(m))
	for i, r := range m {
		rows[i] = string(r)
	}
	return rows
}

func (m Map) Columns() []string {
	columns := make([]string, len(m[0]))
	for x := 0; x < len(columns); x++ {
		var sb strings.Builder
		for y := 0; y < len(m); y++ {
			sb.WriteRune(m[y][x])
		}
		columns[x] = sb.String()
	}
	return columns
}

type Kind int

const (
	Row Kind = iota
	Column
)

func Part1(scanner *bufio.Scanner) (string, error) {
	maps, err := parseInput(scanner)
	if err != nil {
		return "", nil
	}

	var sum int
	for _, m := range maps {
		sum += solve(m, 0)
	}

	result := strconv.Itoa(sum)
	return result, nil
}

func solve(m Map, expectedDiff int) int {
	doSolve := func(patterns []string) int {
		var result int

		for i := 1; i < len(patterns); i++ {
			r := min(i, len(patterns)-i)
			previousIndex, nextIndex := i-1, i

			var count int
			for delta := 0; delta < r; delta++ {
				previous, next := patterns[previousIndex-delta], patterns[nextIndex+delta]
				for j := 0; j < len(previous); j++ {
					if previous[j] == next[j] {
						count++
					}
				}
			}

			maxCount := r * len(patterns[0])
			if (maxCount - count) == expectedDiff {
				result = i
			}
		}

		return result
	}

	rowResult := doSolve(m.Rows())
	columnResult := doSolve(m.Columns())
	return rowResult*100 + columnResult
}

func parseInput(scanner *bufio.Scanner) ([]Map, error) {
	maps := make([]Map, 0)
	currentMap := make([][]rune, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			maps = append(maps, currentMap)
			currentMap = make([][]rune, 0)
			continue
		}

		currentMap = append(currentMap, []rune(line))
	}

	if len(currentMap) > 0 {
		maps = append(maps, currentMap)
	}
	return maps, nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	maps, err := parseInput(scanner)
	if err != nil {
		return "", nil
	}

	var sum int
	for _, m := range maps {
		sum += solve(m, 1)
	}

	result := strconv.Itoa(sum)
	return result, nil
}
