package day9

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLine(t *testing.T) {
	input := "0 3 6 9 12 15"

	numbers, err := parseLine(input)

	require.Nil(t, err)
	assert.Equal(t, []int{0, 3, 6, 9, 12, 15}, numbers)
}

func TestNextLine(t *testing.T) {
	type testCase struct {
		input    []int
		expected []int
	}
	testCases := []testCase{
		{
			[]int{0, 3, 6, 9, 12, 15},
			[]int{3, 3, 3, 3, 3},
		},
		{
			[]int{10, 13, 16, 21, 30, 45},
			[]int{3, 3, 5, 9, 15},
		},
	}
	for _, c := range testCases {
		name := fmt.Sprintf("test %v should return %v", c.input, c.expected)
		t.Run(name, func(t *testing.T) {
			nextLine := nextLine(c.input)
			assert.Equal(t, c.expected, nextLine)
		})
	}
}

func TestExpandHistory(t *testing.T) {
	type testCase struct {
		input    []int
		expected [][]int
	}
	testCases := []testCase{
		{
			[]int{0, 3, 6, 9, 12, 15},
			[][]int{
				{0, 3, 6, 9, 12, 15},
				{3, 3, 3, 3, 3},
				{0, 0, 0, 0},
			},
		},
		{
			[]int{10, 13, 16, 21, 30, 45},
			[][]int{
				{10, 13, 16, 21, 30, 45},
				{3, 3, 5, 9, 15},
				{0, 2, 4, 6},
				{2, 2, 2},
				{0, 0},
			},
		},
	}

	for _, c := range testCases {
		name := fmt.Sprintf("%v should return expected output", c.input)
		t.Run(name, func(t *testing.T) {
			result := expandHistory(c.input)
			assert.Equal(t, c.expected, result)
		})
	}
}

func TestPredictLast(t *testing.T) {
	type testCase struct {
		input    [][]int
		expected int
	}
	testCases := []testCase{
		{
			[][]int{
				{0, 3, 6, 9, 12, 15},
				{3, 3, 3, 3, 3},
				{0, 0, 0, 0},
			},
			18,
		},
		{
			[][]int{
				{1, 3, 6, 10, 15, 21},
				{2, 3, 4, 5, 6},
				{1, 1, 1, 1},
				{0, 0, 0},
			},
			28,
		},
		{
			[][]int{
				{10, 13, 16, 21, 30, 45},
				{3, 3, 5, 9, 15},
				{0, 2, 4, 6},
				{2, 2, 2},
				{0, 0},
			},
			68,
		},
	}

	for i, c := range testCases {
		name := fmt.Sprintf("example %d should return %d", i, c.expected)
		t.Run(name, func(t *testing.T) {
			prediction := predictLast(c.input)
			assert.Equal(t, c.expected, prediction)
		})
	}
}

func TestPart1(t *testing.T) {
	input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part1(scanner)

	require.Nil(t, err)
	assert.Equal(t, "114", result)
}

func TestPredictFirst(t *testing.T) {
	input := [][]int{
		{10, 13, 16, 21, 30, 45},
		{3, 3, 5, 9, 15},
		{0, 2, 4, 6},
		{2, 2, 2},
		{0, 0},
	}

	prediction := predictFirst(input)

	assert.Equal(t, 5, prediction)
}

func TestPart2(t *testing.T) {
	input := `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part2(scanner)

	require.Nil(t, err)
	assert.Equal(t, "2", result)
}
