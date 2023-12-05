package day2

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseGame(t *testing.T) {
	input := "Game 1: 7 green, 4 blue, 3 red; 4 blue, 10 red, 1 green; 1 blue, 9 red"

	game, err := parseGame(input)

	require.Nil(t, err)
	assert.Equal(t, 1, game.Id)
	require.Equal(t, 3, len(game.Subsets))

	subset := game.Subsets[0]
	require.Equal(t, 3, len(subset))
	assert.Equal(t, 7, subset["green"])
	assert.Equal(t, 4, subset["blue"])
	assert.Equal(t, 3, subset["red"])

	subset = game.Subsets[1]
	require.Equal(t, 3, len(subset))
	assert.Equal(t, 4, subset["blue"])
	assert.Equal(t, 10, subset["red"])
	assert.Equal(t, 1, subset["green"])

	subset = game.Subsets[2]
	require.Equal(t, 2, len(subset))
	assert.Equal(t, 1, subset["blue"])
	assert.Equal(t, 9, subset["red"])
}

func TestPart1(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part1(scanner)

	require.Nil(t, err)
	assert.Equal(t, "8", result)
}

func TestMinSubset(t *testing.T) {
	subsets := []Subset{
		map[string]int{
			"blue": 3,
			"red":  4,
		},
		map[string]int{
			"red":   1,
			"green": 2,
			"blue":  6,
		},
		map[string]int{
			"green": 2,
		},
	}
	game := NewGame(1, subsets)

	minSubset := game.MinSubset()

	assert.Equal(t, 4, minSubset["red"])
	assert.Equal(t, 2, minSubset["green"])
	assert.Equal(t, 6, minSubset["blue"])
}

func TestPart2(t *testing.T) {
	input := `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part2(scanner)

	require.Nil(t, err)
	assert.Equal(t, "2286", result)
}
