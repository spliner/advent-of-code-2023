package day4

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

type Card struct {
	Winners map[int]struct{}
	Numbers []int
	Id      int
}

func (c *Card) Points() int {
	count := c.CountMatches()
	if count == 0 {
		return 0
	}

	return int(math.Pow(2, float64(count-1)))
}

func (c *Card) CountMatches() int {
	var count int
	for _, n := range c.Numbers {
		_, ok := c.Winners[n]
		if ok {
			count++
		}
	}
	return count
}

func Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var sum int
	for _, l := range lines {
		card, err := parseLine(strings.TrimSpace(l))
		if err != nil {
			return "", err
		}

		points := card.Points()
		sum += points
	}

	result := strconv.Itoa(sum)
	return result, nil
}

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
func parseLine(line string) (*Card, error) {
	rawCard, rawNumbers, found := strings.Cut(line, ":")
	if !found {
		return nil, errors.New("could not find ':' in input")
	}

	_, rawId, found := strings.Cut(rawCard, " ")
	if !found {
		return nil, errors.New("could not find ' ' in raw card input")
	}

	cardId, err := strconv.Atoi(strings.TrimSpace(rawId))
	if err != nil {
		return nil, err
	}

	card := Card{
		Id: cardId,
	}

	rawWinners, rawNumbers, found := strings.Cut(rawNumbers, "|")
	if !found {
		return nil, errors.New("could not find '|' in raw numbers input")
	}

	winningNumbers := strings.Split(rawWinners, " ")
	card.Winners = make(map[int]struct{}, 0)
	for _, n := range winningNumbers {
		n := strings.TrimSpace(n)
		if n == "" {
			continue
		}

		value, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}

		card.Winners[value] = struct{}{}
	}

	drawNumbers := strings.Split(rawNumbers, " ")
	card.Numbers = make([]int, 0)
	for _, n := range drawNumbers {
		n := strings.TrimSpace(n)
		if n == "" {
			continue
		}

		value, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}

		card.Numbers = append(card.Numbers, value)
	}

	return &card, nil
}

func Part2(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	quantities := make([]int, len(lines))
	for i := range quantities {
		quantities[i] = 1
	}

	for i, l := range lines {
		card, err := parseLine(strings.TrimSpace(l))
		if err != nil {
			return "", err
		}

		matches := card.CountMatches()
		for j := 0; j < matches; j++ {
			index := i + j + 1
			quantities[index] = quantities[index] + quantities[i]
		}
	}

	var sum int
	for _, q := range quantities {
		sum += q
	}

	result := strconv.Itoa(sum)
	return result, nil
}
