package day7

import "bufio"

var labelOrders map[rune]int

func init() {
	labelOrders = map[rune]int{
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
	}
}

type Card struct {
	Label rune
}

func (c *Card) LabelOrder() int {
	return labelOrders[c.Label]
}

type Hand struct {
	Cards []*Card
	Bid   int
}

func Part1(scanner *bufio.Scanner) (string, error) {
	return "", nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	return "", nil
}
