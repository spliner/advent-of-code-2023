package day2

import (
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

func Part1(input string) (string, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var sum int

	subsetLimits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

gameLoop:
	for _, l := range lines {
		game, err := parseGame(l)
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

func Part2(input string) (string, error) {
	return "", nil
}
