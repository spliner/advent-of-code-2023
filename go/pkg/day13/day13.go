package day13

import (
	"bufio"
	"strconv"
	"strings"
)

type Map [][]rune

func Part1(scanner *bufio.Scanner) (string, error) {
	maps, err := parseInput(scanner)
	if err != nil {
		return "", nil
	}

	var sum int
	for _, m := range maps {
		rows := rows(m)
		start, _, found := findReflection(rows)
		if found {
			sum += 100 * (start + 1)
		} else {
			columns := columns(m)
			start, _, found = findReflection(columns)
			sum += start + 1
		}
	}

	result := strconv.Itoa(sum)
	return result, nil
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

func parseMap(scanner *bufio.Scanner) (Map, error) {
	m := make([][]rune, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		line := strings.TrimSpace(scanner.Text())
		m = append(m, []rune(line))
	}
	return m, nil
}

func rows(m Map) []string {
	rows := make([]string, len(m))
	for i, r := range m {
		rows[i] = string(r)
	}
	return rows
}

func findReflection(patterns []string) (start, end int, found bool) {
	for i := 0; i < len(patterns)-1; i++ {
		if patterns[i] == patterns[i+1] && isReflection(patterns, i, i+1) {
			start = i
			end = i + 1
			found = true
			return
		}
	}
	return
}

func isReflection(patterns []string, startIndex, endIndex int) bool {
	for startIndex >= 0 && endIndex < len(patterns) {
		if patterns[startIndex] != patterns[endIndex] {
			return false
		}
		startIndex--
		endIndex++
	}
	return true
}

func columns(m Map) []string {
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

func Part2(scanner *bufio.Scanner) (string, error) {
	return "", nil
}
