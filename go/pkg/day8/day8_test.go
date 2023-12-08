package day8

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPart1(t *testing.T) {
	type testCase struct {
		input         string
		expectedSteps string
	}
	testCases := []testCase{
		{
			input: `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`,
			expectedSteps: "2",
		},
		{
			input: `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`,
			expectedSteps: "6",
		},
	}
	for i, c := range testCases {
		name := fmt.Sprintf("input %d should take %s steps", (i + 1), c.expectedSteps)
		t.Run(name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(c.input))

			result, err := Part1(scanner)

			require.Nil(t, err)
			assert.Equal(t, c.expectedSteps, result)
		})
	}
}

func TestPart2(t *testing.T) {
	input := `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part2(scanner)

	require.Nil(t, err)
	assert.Equal(t, "6", result)
}
