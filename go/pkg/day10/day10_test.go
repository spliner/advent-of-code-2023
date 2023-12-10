package day10

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseMap(t *testing.T) {
	input := `.....
.S-7.
.|.|.
.L-J.
.....`
	scanner := bufio.NewScanner(strings.NewReader(input))

	m, err := parseMap(scanner)

	require.Nil(t, err)
	assert.Equal(t, Position{2, 2}, m.Start)
}

func TestPart1(t *testing.T) {
	type testCase struct {
		input         string
		expectedSteps string
	}
	testCases := []testCase{
		{
			input: `.....
.S-7.
.|.|.
.L-J.
.....`,
			expectedSteps: "4",
		},
		{
			input: `-L|F7
				7S-7|
				L|7||
				-L-J|
				L|-JF`,
			expectedSteps: "4",
		},
		{
			input: `7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`,
			expectedSteps: "8",
		},
	}
	for i, c := range testCases {
		name := fmt.Sprintf("day 10 part 1 case %d should return %s", i, c.expectedSteps)
		t.Run(name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(c.input))

			result, err := Part1(scanner)

			require.Nil(t, err)
			assert.Equal(t, c.expectedSteps, result)
		})
	}
}
