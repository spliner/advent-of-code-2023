package day3

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseSchematic(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	scanner := bufio.NewScanner(strings.NewReader(input))

	schematic, err := parseSchematic(scanner)

	require.Nil(t, err)

	assert.Equal(t, schematic.Coordinates[0][0].Value, '4')
	assert.Equal(t, schematic.Coordinates[0][1].Value, '6')
	assert.Equal(t, schematic.Coordinates[0][2].Value, '7')
	assert.Equal(t, schematic.Coordinates[1][3].Value, '*')

	assert.Equal(t, 467, schematic.Numbers[0].Value)
	assert.Equal(t, 114, schematic.Numbers[1].Value)
}

func TestPart1(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part1(scanner)

	require.Nil(t, err)
	assert.Equal(t, "4361", result)
}

func TestPart2(t *testing.T) {
	input := `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part2(scanner)

	require.Nil(t, err)
	assert.Equal(t, "467835", result)
}
