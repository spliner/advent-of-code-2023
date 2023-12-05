package day2

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

type Game struct {
	Subsets []Subset
	Id      int
}

type Subset map[string]int

func NewGame(id int, subsets []Subset) *Game {
	return &Game{
		Id:      id,
		Subsets: subsets,
	}
}

func (g *Game) MinSubset() Subset {
	minSubset := make(map[string]int)
	for _, s := range g.Subsets {
		for color, count := range s {
			current, ok := minSubset[color]
			if !ok || count > current {
				minSubset[color] = count
			}
		}
	}
	return minSubset
}

func Part1(scanner *bufio.Scanner) (string, error) {
	var sum int

	subsetLimits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

gameLoop:
	for scanner.Scan() {
		game, err := parseGame(scanner.Text())
		if err != nil {
			return "", err
		}

		for _, s := range game.Subsets {
			if !subsetWithinLimits(s, subsetLimits) {
				continue gameLoop
			}
		}

		sum += game.Id
	}

	return strconv.Itoa(sum), nil
}

func subsetWithinLimits(subset Subset, limits Subset) bool {
	for color, limit := range limits {
		val, ok := subset[color]
		if !ok {
			continue
		}

		if val > limit {
			return false
		}
	}

	return true
}

// Game 1: 7 green, 4 blue, 3 red; 4 blue, 10 red, 1 green; 1 blue, 9 red
func parseGame(input string) (*Game, error) {
	gameSplit := strings.Split(input, ": ")
	if len(gameSplit) != 2 {
		return nil, errors.New("invalid input")
	}

	gameIdSplit := strings.Split(gameSplit[0], " ")
	if len(gameIdSplit) != 2 {
		return nil, errors.New("invalid input")
	}

	gameId, err := strconv.Atoi(gameIdSplit[1])
	if err != nil {
		return nil, err
	}

	subsetSplit := strings.Split(gameSplit[1], "; ")
	subsets := make([]Subset, len(subsetSplit))
	for i, s := range subsetSplit {
		m := make(map[string]int)
		subsets[i] = m

		cubes := strings.Split(s, ", ")
		for _, c := range cubes {
			cubeSplit := strings.Split(c, " ")
			if len(cubeSplit) != 2 {
				return nil, errors.New("invalid input")
			}

			count, err := strconv.Atoi(cubeSplit[0])
			if err != nil {
				return nil, err
			}

			color := cubeSplit[1]
			m[color] = count
		}
	}

	return NewGame(gameId, subsets), nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	var sum int
	for scanner.Scan() {
		game, err := parseGame(strings.TrimSpace(scanner.Text()))
		if err != nil {
			return "", err
		}

		minSubset := game.MinSubset()
		power := 1
		count, ok := minSubset["red"]
		if ok {
			power *= count
		}

		count, ok = minSubset["green"]
		if ok {
			power *= count
		}

		count, ok = minSubset["blue"]
		if ok {
			power *= count
		}

		sum += power
	}

	return strconv.Itoa(sum), nil
}
