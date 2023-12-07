package day7

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBids(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	scanner := bufio.NewScanner(strings.NewReader(input))

	hands, err := parseBids(scanner)

	require.Nil(t, err)
	require.Equal(t, 5, len(hands))
	require.Equal(t, uint64(765), hands[0].Value)
	require.Equal(t, uint64(684), hands[1].Value)
}

func TestHandType(t *testing.T) {
	type testCase struct {
		hand         string
		expectedType HandType
	}
	testCases := []testCase{
		{"AAAAA", FiveOfAKind},
		{"AA8AA", FourOfAKind},
		{"23332", FullHouse},
		{"TTT98", ThreeOfAKind},
		{"23432", TwoPair},
		{"A23A4", OnePair},
		{"23456", HighCard},
	}
	for _, c := range testCases {
		name := fmt.Sprintf("%s should be type %d", c.hand, c.expectedType)
		t.Run(name, func(t *testing.T) {
			hand := Hand(c.hand)
			handType := hand.Type()
			assert.Equal(t, c.expectedType, handType)
		})
	}
}

func TestBidSort(t *testing.T) {
	bids := Bids{
		{Hand("32T3K"), 765},
		{Hand("T55J5"), 684},
		{Hand("KK677"), 28},
		{Hand("KTJJT"), 220},
		{Hand("QQQJA"), 483},
	}

	want := []string{
		"32T3K",
		"KTJJT",
		"KK677",
		"T55J5",
		"QQQJA",
	}

	bids.Sort(func(h Hand) HandType { return h.Type() }, func(c Card) int { return c.JokerLabelOrder() })

	got := make([]string, len(bids))
	for i, b := range bids {
		got[i] = string(b.Hand)
	}

	assert.Equal(t, want, got)
}

func TestPart1(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part1(scanner)

	require.Nil(t, err)
	assert.Equal(t, "6440", result)
}

func TestHighestType(t *testing.T) {
	type testCase struct {
		hand         string
		expectedType HandType
	}
	testCases := []testCase{
		{"32T3K", OnePair},
		{"T55J5", FourOfAKind},
		{"KK677", TwoPair},
		{"KTJJT", FourOfAKind},
		{"QQQJA", FourOfAKind},
	}
	for _, c := range testCases {
		name := fmt.Sprintf("%s should be type %d", c.hand, c.expectedType)
		t.Run(name, func(t *testing.T) {
			hand := Hand(c.hand)
			handType := hand.HighestType()
			assert.Equal(t, c.expectedType, handType)
		})
	}
}

func TestPart2(t *testing.T) {
	input := `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`
	scanner := bufio.NewScanner(strings.NewReader(input))

	result, err := Part2(scanner)

	require.Nil(t, err)
	assert.Equal(t, "5905", result)
}
