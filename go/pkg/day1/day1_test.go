package day1

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCase struct {
	input    string
	expected int
}

func TestParsePart1Calibration(t *testing.T) {
	testCases := []testCase{
		{
			input:    "1abc2",
			expected: 12,
		},
		{
			input:    "pqr3stu8vwx",
			expected: 38,
		},
		{
			input:    "a1b2c3d4e5f",
			expected: 15,
		},
		{
			input:    "treb7uchet",
			expected: 77,
		},
	}
	for _, c := range testCases {
		name := fmt.Sprintf("%s should return %d", c.input, c.expected)
		t.Run(name, func(t *testing.T) {
			calibration, err := parsePart1Calibration(c.input)
			require.Nil(t, err)
			assert.Equal(t, c.expected, calibration)
		})
	}
}

func TestPart1(t *testing.T) {
	input := `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part1(scanner)

	require.Nil(t, err)
	assert.Equal(t, "142", result)
}

func TestPart2Calibration(t *testing.T) {
	testCases := []testCase{
		{
			input:    "two1nine",
			expected: 29,
		},
		{
			input:    "eightwothree",
			expected: 83,
		},
		{
			input:    "abcone2threexyz",
			expected: 13,
		},
		{
			input:    "xtwone3four",
			expected: 24,
		},
		{
			input:    "4nineeightseven2",
			expected: 42,
		},
		{
			input:    "zoneight234",
			expected: 14,
		},
		{
			input:    "7pqrstsixteen",
			expected: 76,
		},
	}
	for _, c := range testCases {
		name := fmt.Sprintf("%s should return %d", c.input, c.expected)
		t.Run(name, func(t *testing.T) {
			calibration, err := parsePart2Calibration(c.input)
			require.Nil(t, err)
			assert.Equal(t, c.expected, calibration)
		})
	}
}

func TestPart2(t *testing.T) {
	input := `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part2(scanner)

	require.Nil(t, err)
	assert.Equal(t, "281", result)
}
