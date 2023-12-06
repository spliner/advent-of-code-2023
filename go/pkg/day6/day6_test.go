package day6

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseInput(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`
	scanner := bufio.NewScanner(strings.NewReader(input))

	races, err := parseInput(scanner)

	require.Nil(t, err)
	assert.Equal(t, 3, len(races))
	assert.Equal(t, Race{7, 9}, *races[0])
	assert.Equal(t, Race{15, 40}, *races[1])
	assert.Equal(t, Race{30, 200}, *races[2])
}

func TestDistanceTravelled(t *testing.T) {
	type testCase struct {
		timeHoldingButton int
		duration          int
		expectedDistance  int
	}
	testCases := []testCase{
		{0, 7, 0},
		{1, 7, 6},
		{2, 7, 10},
		{3, 7, 12},
		{4, 7, 12},
		{5, 7, 10},
		{6, 7, 6},
		{7, 7, 0},
	}

	for _, c := range testCases {
		name := fmt.Sprintf("holding button for %d with race taking %d ms should travel %d",
			c.timeHoldingButton,
			c.duration,
			c.expectedDistance)
		t.Run(name, func(t *testing.T) {
			distance := distanceTravelled(c.duration, c.timeHoldingButton)
			assert.Equal(t, c.expectedDistance, distance)
		})
	}
}

func TestPart1(t *testing.T) {
	input := `Time:      7  15   30
Distance:  9  40  200`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part1(scanner)

	require.Nil(t, err)
	assert.Equal(t, "288", result)
}
