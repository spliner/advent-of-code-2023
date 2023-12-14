package day14

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Tile rune

const (
	Empty       = '.'
	RoundedRock = 'O'
	CubeRock    = '#'
)

type Dish struct {
	Grid [][]Tile
}

func (d *Dish) TiltNorth() *Dish {
	newGrid := make([][]Tile, len(d.Grid))
	copy(newGrid, d.Grid)
	newDish := Dish{newGrid}

	for y, line := range newGrid {
		for x, tile := range line {
			if tile != RoundedRock {
				continue
			}

			// Move rounded rock as far north as we can
			newY := y - 1
			for newY >= 0 {
				northernTile := newGrid[newY][x]
				if northernTile == Empty {
					newGrid[newY][x] = tile
					newGrid[newY+1][x] = Empty
				} else {
					break
				}
				newY--
			}
		}
	}

	return &newDish
}

func (d *Dish) TotalLoad() int {
	var load int
	for y, line := range d.Grid {
		for _, tile := range line {
			if tile == RoundedRock {
				load += len(d.Grid) - y
			}
		}
	}
	return load
}

func (d *Dish) String() string {
	var sb strings.Builder
	for _, line := range d.Grid {
		for _, tile := range line {
			sb.WriteRune(rune(tile))
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func Part1(scanner *bufio.Scanner) (string, error) {
	dish, err := parseDish(scanner)
	if err != nil {
		return "", err
	}

	tilted := dish.TiltNorth()
	totalLoad := tilted.TotalLoad()

	result := strconv.Itoa(totalLoad)
	return result, nil
}

func parseDish(scanner *bufio.Scanner) (*Dish, error) {
	tiles := make([][]Tile, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		lineTiles := make([]Tile, len(line))
		for i, r := range line {
			var tile Tile
			if r == Empty {
				tile = Empty
			} else if r == RoundedRock {
				tile = RoundedRock
			} else if r == CubeRock {
				tile = CubeRock
			}
			if tile == 0 {
				return nil, fmt.Errorf("invalid rune: %s", string(r))
			}

			lineTiles[i] = tile
		}

		tiles = append(tiles, lineTiles)
	}

	dish := Dish{tiles}
	return &dish, nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	return "", nil
}
