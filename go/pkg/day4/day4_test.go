package day4

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseLine(t *testing.T) {
	line := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"

	card, err := parseLine(line)

	require.Nil(t, err)
	assert.Equal(t, 1, card.Id)
	assert.Equal(t, 5, len(card.Winners))
	assert.Equal(t, 8, len(card.Numbers))
}

func TestPoints(t *testing.T) {
	type testCase struct {
		line           string
		expectedPoints int
	}

	testCases := []testCase{
		{
			line:           "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
			expectedPoints: 8,
		},
		{
			line:           "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
			expectedPoints: 2,
		},
		{
			line:           "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
			expectedPoints: 2,
		},
		{
			line:           "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
			expectedPoints: 1,
		},
		{
			line:           "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
			expectedPoints: 0,
		},
		{
			line:           "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
			expectedPoints: 0,
		},
	}
	for i, c := range testCases {
		name := fmt.Sprintf("card %d should be worth %d points", i+1, c.expectedPoints)
		t.Run(name, func(t *testing.T) {
			card, err := parseLine(c.line)
			require.Nil(t, err)

			points := card.Points()
			assert.Equal(t, c.expectedPoints, points)
		})
	}
}

func TestPart1(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	result, err := Part1(input)

	require.Nil(t, err)
	assert.Equal(t, "13", result)
}

func TestPart2(t *testing.T) {
	input := `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

	result, err := Part2(input)

	require.Nil(t, err)
	assert.Equal(t, "30", result)
}
