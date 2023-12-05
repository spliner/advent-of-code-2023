package day5

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapDestination(t *testing.T) {
	type testCase struct {
		source              int
		expectedDestination int
	}
	testCases := []testCase{
		{0, 0},
		{1, 1},
		{48, 48},
		{49, 49},
		{50, 52},
		{51, 53},
		{96, 98},
		{97, 99},
		{98, 50},
		{99, 51},
	}

	lines := []*MapLine{
		NewMapLine(50, 98, 2),
		NewMapLine(52, 50, 48),
	}
	m := NewMap(lines)

	for _, c := range testCases {
		name := fmt.Sprintf("source %d should map to %d", c.source, c.expectedDestination)
		t.Run(name, func(t *testing.T) {
			destination := m.Destination(c.source)
			assert.Equal(t, c.expectedDestination, destination)
		})
	}
}

func TestPart1(t *testing.T) {
	input := `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part1(scanner)

	require.Nil(t, err)
	assert.Equal(t, "35", result)
}
