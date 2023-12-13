package day13

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindColumnReflection(t *testing.T) {
	input := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`
	scanner := bufio.NewScanner(strings.NewReader(input))

	m, err := parseMap(scanner)
	require.Nil(t, err)
	c := columns(m)
	start, end, found := findReflection(c)

	assert.Equal(t, true, found)
	assert.Equal(t, 4, start)
	assert.Equal(t, 5, end)
}

func TestFindRowReflection(t *testing.T) {
	input := `#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`
	scanner := bufio.NewScanner(strings.NewReader(input))

	m, err := parseMap(scanner)
	require.Nil(t, err)
	c := rows(m)
	start, end, found := findReflection(c)

	assert.Equal(t, true, found)
	assert.Equal(t, 3, start)
	assert.Equal(t, 4, end)
}

func TestPart1(t *testing.T) {
	input := `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part1(scanner)

	require.Nil(t, err)
	assert.Equal(t, "405", result)
}
