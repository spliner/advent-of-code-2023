package day1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseCalibration(t *testing.T) {
	type testCase struct {
		input    string
		expected int
	}
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
			calibration, err := parseCalibration(c.input)
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

	result, err := Part1(input)

	require.Nil(t, err)
	assert.Equal(t, "142", result)
}
