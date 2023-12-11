package day11

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPointDistance(t *testing.T) {
	type testCase struct {
		origin           Point
		destination      Point
		expectedDistance int
	}
	testCases := []testCase{
		{Point{1, 6}, Point{5, 11}, 9},
		{Point{4, 0}, Point{9, 10}, 15},
		{Point{0, 2}, Point{12, 7}, 17},
		{Point{0, 11}, Point{5, 11}, 5},
	}
	for _, c := range testCases {
		name := fmt.Sprintf("distance from %v to %v should be %d", c.origin, c.destination, c.expectedDistance)
		t.Run(name, func(t *testing.T) {
			distance := c.origin.Distance(c.destination)
			assert.Equal(t, c.expectedDistance, distance)
		})
	}
}

func TestExpand(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	scanner := bufio.NewScanner(strings.NewReader(input))

	universe, err := parseUniverse(scanner)
	require.Nil(t, err)

	expandedGalaxies := universe.Expand(2)
	assert.Equal(t, Point{4, 0}, expandedGalaxies[0].Position)
	assert.Equal(t, Point{9, 1}, expandedGalaxies[1].Position)
	assert.Equal(t, Point{0, 2}, expandedGalaxies[2].Position)
	assert.Equal(t, Point{8, 5}, expandedGalaxies[3].Position)
	assert.Equal(t, Point{1, 6}, expandedGalaxies[4].Position)
	assert.Equal(t, Point{12, 7}, expandedGalaxies[5].Position)
	assert.Equal(t, Point{9, 10}, expandedGalaxies[6].Position)
	assert.Equal(t, Point{0, 11}, expandedGalaxies[7].Position)
	assert.Equal(t, Point{5, 11}, expandedGalaxies[8].Position)
}

func TestPart1(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part1(scanner)

	require.Nil(t, err)
	assert.Equal(t, "374", result)
}

func TestRun(t *testing.T) {
	input := `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	type testCase struct {
		expandTimes    int
		expectedResult string
	}
	testCases := []testCase{
		{2, "374"},
		{10, "1030"},
		{100, "8410"},
	}
	for _, c := range testCases {
		name := fmt.Sprintf("expanding %d times should return %s", c.expandTimes, c.expectedResult)
		t.Run(name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(input))

			result, err := run(scanner, c.expandTimes)

			require.Nil(t, err)
			assert.Equal(t, c.expectedResult, result)
		})
	}
}
