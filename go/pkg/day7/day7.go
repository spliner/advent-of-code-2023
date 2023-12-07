package day7

import (
	"bufio"
	"errors"
	"sort"
	"strconv"
	"strings"
)

type HandType int64

const (
	HighCard HandType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Card rune

var labelOrders map[Card]int = map[Card]int{
	'A': 0,
	'K': 1,
	'Q': 2,
	'J': 3,
	'T': 4,
	'9': 5,
	'8': 6,
	'7': 7,
	'6': 8,
	'5': 9,
	'4': 10,
	'3': 11,
	'2': 12,
}

var jokerLabelOrders map[Card]int = map[Card]int{
	'A': 0,
	'K': 1,
	'Q': 2,
	'T': 3,
	'9': 4,
	'8': 5,
	'7': 6,
	'6': 7,
	'5': 8,
	'4': 9,
	'3': 10,
	'2': 11,
	'J': 12,
}

func (c Card) LabelOrder() int {
	return labelOrders[c]
}

func (c Card) JokerLabelOrder() int {
	return jokerLabelOrders[c]
}

type Hand []Card

func (h Hand) Type() HandType {
	frequencies := make(map[Card]int)
	var highestCount int
	for _, c := range h {
		frequencies[c] = frequencies[c] + 1
		if frequencies[c] > highestCount {
			highestCount = frequencies[c]
		}
	}

	switch len(frequencies) {
	case 1:
		return FiveOfAKind
	case 2:
		if highestCount == 4 {
			return FourOfAKind
		}
		return FullHouse
	case 3:
		if highestCount == 3 {
			return ThreeOfAKind
		}
		return TwoPair
	case 4:
		return OnePair
	default:
		return HighCard
	}
}

func (h Hand) HighestType() HandType {
	frequencyMap := make(map[Card]int)
	for _, c := range h {
		frequencyMap[c] = frequencyMap[c] + 1
	}

	sortedFrequencies := make([]int, 0)
	for k, v := range frequencyMap {
		if k == 'J' {
			continue
		}

		sortedFrequencies = append(sortedFrequencies, v)
	}

	sort.Slice(sortedFrequencies, func(i, j int) bool {
		x := sortedFrequencies[i]
		y := sortedFrequencies[j]
		return y < x
	})

	jokerFrequency := frequencyMap['J']

	if len(frequencyMap) == 1 || sortedFrequencies[0]+jokerFrequency == 5 {
		return FiveOfAKind
	}
	if sortedFrequencies[0] == 4 || sortedFrequencies[0]+jokerFrequency == 4 {
		return FourOfAKind
	}
	if len(frequencyMap) == 2 || sortedFrequencies[0]+sortedFrequencies[1]+jokerFrequency == 5 {
		return FullHouse
	}
	if sortedFrequencies[0] == 3 || sortedFrequencies[0]+jokerFrequency == 3 {
		return ThreeOfAKind
	}
	if sortedFrequencies[0] == 2 && sortedFrequencies[1] == 2 {
		return TwoPair
	}
	if sortedFrequencies[0] == 2 || jokerFrequency > 0 {
		return OnePair
	}

	return HighCard
}

type Bid struct {
	Hand  Hand
	Value uint64
}

type Bids []Bid

func (b Bids) Sort(handTypeFunc func(Hand) HandType, labelOrderFunc func(c Card) int) {
	sort.Slice(b, func(i, j int) bool {
		x := b[i]
		y := b[j]

		xType := handTypeFunc(x.Hand)
		yType := handTypeFunc(y.Hand)

		if xType == yType {
			for i := 0; i < len(x.Hand); i++ {
				xOrder := labelOrderFunc(x.Hand[i])
				yOrder := labelOrderFunc(y.Hand[i])

				if xOrder == yOrder {
					continue
				}

				return xOrder > yOrder
			}
		}

		return xType < yType
	})
}

func Part1(scanner *bufio.Scanner) (string, error) {
	bids, err := parseBids(scanner)
	if err != nil {
		return "", err
	}

	bids.Sort(func(h Hand) HandType { return h.Type() }, func(c Card) int { return c.LabelOrder() })
	var sum uint64
	for i, b := range bids {
		sum += uint64(i+1) * b.Value
	}

	result := strconv.FormatUint(sum, 10)

	return result, nil
}

func parseBids(scanner *bufio.Scanner) (Bids, error) {
	bids := make([]Bid, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := scanner.Text()
		rawCards, rawBid, found := strings.Cut(line, " ")
		if !found {
			return nil, errors.New("could not find ' ' in line")
		}

		cards := make([]Card, len(rawCards))
		for i, c := range rawCards {
			cards[i] = Card(c)
		}

		value, err := strconv.Atoi(rawBid)
		if err != nil {
			return nil, err
		}

		bid := Bid{
			Hand:  cards,
			Value: uint64(value),
		}
		bids = append(bids, bid)
	}

	return Bids(bids), nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	bids, err := parseBids(scanner)
	if err != nil {
		return "", err
	}

	bids.Sort(func(h Hand) HandType { return h.HighestType() }, func(c Card) int { return c.JokerLabelOrder() })

	var sum uint64
	for i, b := range bids {
		sum += uint64(i+1) * b.Value
	}

	result := strconv.FormatUint(sum, 10)

	return result, nil
}
