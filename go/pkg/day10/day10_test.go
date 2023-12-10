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

func TestPart2(t *testing.T) {
	type testCase struct {
		input         string
		expectedTiles string
	}
	testCases := []testCase{
		{
			input: `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`,
			expectedTiles: "4",
		},
		{
			input: `..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........`,
			expectedTiles: "4",
		},
		{
			input: `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`,
			expectedTiles: "8",
		},
		{
			input: `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`,
			expectedTiles: "10",
		},
	}

	for i, c := range testCases {
		name := fmt.Sprintf("day 10 part 2 case %d should return %s", i, c.expectedTiles)
		t.Run(name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(c.input))

			result, err := Part2(scanner)

			require.Nil(t, err)
			assert.Equal(t, c.expectedTiles, result)
		})
	}
}
