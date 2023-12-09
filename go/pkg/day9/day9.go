package day9

import (
	"bufio"
	"strconv"
	"strings"
)

func Part1(scanner *bufio.Scanner) (string, error) {
	histories, err := parseInput(scanner)
	if err != nil {
		return "", nil
	}

	var sum int
	for _, h := range histories {
		expanded := expandHistory(h)
		prediction := predictLast(expanded)
		sum += prediction
	}

	result := strconv.Itoa(sum)
	return result, nil
}

func parseInput(scanner *bufio.Scanner) ([][]int, error) {
	lines := make([][]int, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		rawLine := strings.TrimSpace(scanner.Text())
		if rawLine == "" {
			continue
		}

		line, err := parseLine(rawLine)
		if err != nil {
			return nil, err
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func parseLine(input string) ([]int, error) {
	rawNumbers := strings.Split(strings.TrimSpace(input), " ")
	numbers := make([]int, len(rawNumbers))
	for i, rawNumber := range rawNumbers {
		number, err := strconv.Atoi(strings.TrimSpace(rawNumber))
		if err != nil {
			return nil, err
		}

		numbers[i] = number
	}

	return numbers, nil
}

func expandHistory(line []int) [][]int {
	result := [][]int{line}
	currentLine := line
	for !allZeroes(currentLine) {
		nextLine := nextLine(currentLine)
		result = append(result, nextLine)
		currentLine = nextLine
	}
	return result
}

func allZeroes(line []int) bool {
	for _, n := range line {
		if n != 0 {
			return false
		}
	}
	return true
}

func nextLine(line []int) []int {
	nextLine := make([]int, len(line)-1)
	for i := 0; i < len(line)-1; i++ {
		nextLine[i] = line[i+1] - line[i]
	}
	return nextLine
}

func predictLast(input [][]int) int {
	prediction := 0
	for i := len(input) - 2; i >= 0; i-- {
		history := input[i]
		prediction = history[len(history)-1] + prediction
	}
	return prediction
}

func Part2(scanner *bufio.Scanner) (string, error) {
	histories, err := parseInput(scanner)
	if err != nil {
		return "", nil
	}

	var sum int
	for _, h := range histories {
		expanded := expandHistory(h)
		prediction := predictFirst(expanded)
		sum += prediction
	}

	result := strconv.Itoa(sum)
	return result, nil
}

func predictFirst(input [][]int) int {
	prediction := 0
	for i := len(input) - 2; i >= 0; i-- {
		history := input[i]
		prediction = history[0] - prediction
	}
	return prediction
}
