package day12

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidArrangement(t *testing.T) {
	type testCase struct {
		line           string
		damagedSprings []int
		expected       bool
	}
	testCases := []testCase{
		{"#.#.###", []int{1, 1, 3}, true},
		{"###.#..#.", []int{1, 1, 3}, false},
		{"###.###", []int{1, 1, 3}, false},
		{"#.#.###", []int{1, 1, 2}, false},
		{".#...#....###.", []int{1, 1, 3}, true},
		{".#.#.#....###.", []int{1, 1, 3}, false},
		{".#.###.#.######", []int{1, 3, 1, 6}, true},
		{".#.###.#.######", []int{1, 3, 1, 5}, false},
		{"...###.#.######", []int{1, 3, 1, 6}, false},
		{"####.#...#...", []int{4, 1, 1}, true},
		{"#....######..#####.", []int{1, 6, 5}, true},
		{".###.##....#", []int{3, 2, 1}, true},
	}
	for _, c := range testCases {
		name := fmt.Sprintf("%s %v should return %t", c.line, c.damagedSprings, c.expected)
		t.Run(name, func(t *testing.T) {
			valid := validArrangement(c.line, c.damagedSprings)
			assert.Equal(t, c.expected, valid)
		})
	}
}

func TestGroupArrangements(t *testing.T) {
	type testCase struct {
		input    string
		expected []string
	}
	testCases := []testCase{
		{"?", []string{".", "#"}},
		{"??", []string{"..", ".#", "#.", "##"}},
		{"???", []string{
			"...",
			"..#",
			".#.",
			".##",
			"#..",
			"#.#",
			"##.",
			"###",
		}},
	}

	for _, c := range testCases {
		name := fmt.Sprintf("%s should return %d results", c.input, len(c.expected))
		t.Run(name, func(t *testing.T) {
			result := groupArrangements(c.input)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestPossibleArrangements(t *testing.T) {
	type testCase struct {
		input    string
		expected []string
	}
	testCases := []testCase{
		{"?.?.###", []string{"....###", "..#.###", "#...###", "#.#.###"}},
	}
	for _, c := range testCases {
		name := fmt.Sprintf("%s should return %d results", c.input, len(c.expected))
		t.Run(name, func(t *testing.T) {
			arrangements := possibleArrangements(c.input)
			assert.Equal(t, c.expected, arrangements)
		})
	}
}

func TestValidArrangements(t *testing.T) {
	type testCase struct {
		record   Record
		expected int
	}
	testCases := []testCase{
		{Record{"???.###", []int{1, 1, 3}}, 1},
		{Record{".??..??...?##.", []int{1, 1, 3}}, 4},
		{Record{"?#?#?#?#?#?#?#?", []int{1, 3, 1, 6}}, 1},
		{Record{"????.#...#...", []int{4, 1, 1}}, 1},
		{Record{"????.######..#####.", []int{1, 6, 5}}, 4},
		{Record{"?###????????", []int{3, 2, 1}}, 10},
	}
	for _, c := range testCases {
		name := fmt.Sprintf("%s and %v should return %d valid arrangements", c.record.Line, c.record.DamagedSprings, c.expected)
		t.Run(name, func(t *testing.T) {
			validArrangements := validArrangements(&c.record)
			assert.Equal(t, c.expected, len(validArrangements))
		})
	}
}
