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

func (d *Dish) TiltNorth() {
	for y, line := range d.Grid {
		for x, tile := range line {
			if tile != RoundedRock {
				continue
			}

			newY := y - 1
			for newY >= 0 {
				northernTile := d.Grid[newY][x]
				if northernTile != Empty {
					break
				}

				d.Grid[newY][x] = tile
				d.Grid[newY+1][x] = Empty
				newY--
			}
		}
	}
}

func (d *Dish) TiltSouth() {
	for y := len(d.Grid) - 1; y >= 0; y-- {
		for x := 0; x < len(d.Grid[y]); x++ {
			tile := d.Grid[y][x]
			if tile != RoundedRock {
				continue
			}

			newY := y + 1
			for newY < len(d.Grid) {
				southernTile := d.Grid[newY][x]
				if southernTile != Empty {
					break
				}

				d.Grid[newY][x] = tile
				d.Grid[newY-1][x] = Empty
				newY++
			}
		}
	}
}

func (d *Dish) TiltEast() {
	for x := len(d.Grid[0]) - 1; x >= 0; x-- {
		for y := 0; y < len(d.Grid); y++ {
			tile := d.Grid[y][x]
			if tile != RoundedRock {
				continue
			}

			newX := x + 1
			for newX < len(d.Grid[0]) {
				easternTile := d.Grid[y][newX]
				if easternTile != Empty {
					break
				}

				d.Grid[y][newX] = tile
				d.Grid[y][newX-1] = Empty
				newX++
			}
		}
	}
}

func (d *Dish) TiltWest() {
	for x := 0; x < len(d.Grid[0]); x++ {
		for y := 0; y < len(d.Grid); y++ {
			tile := d.Grid[y][x]
			if tile != RoundedRock {
				continue
			}

			newX := x - 1
			for newX >= 0 {
				easternTile := d.Grid[y][newX]
				if easternTile != Empty {
					break
				}

				d.Grid[y][newX] = tile
				d.Grid[y][newX+1] = Empty
				newX--
			}
		}
	}
}

func (d *Dish) Cycle() {
	d.TiltNorth()
	d.TiltWest()
	d.TiltSouth()
	d.TiltEast()
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

func (d *Dish) Hash() string {
	var sb strings.Builder
	for _, line := range d.Grid {
		for _, tile := range line {
			sb.WriteRune(rune(tile))
		}
	}
	return sb.String()
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

	dish.TiltNorth()
	totalLoad := dish.TotalLoad()

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
	dish, err := parseDish(scanner)
	if err != nil {
		return "", err
	}

	indexLookup := make(map[string]int)
	var i int
	for i = 0; i < 1_000_000_000; i++ {
		dish.Cycle()
		hash := dish.Hash()
		if _, ok := indexLookup[hash]; ok {
			break
		} else {
			indexLookup[hash] = i
		}
	}

	cycleLength := i - indexLookup[dish.Hash()]
	cycles := (1_000_000_000 - i - 1) % cycleLength
	for i = 0; i < cycles; i++ {
		dish.Cycle()
	}

	load := dish.TotalLoad()
	result := strconv.Itoa(load)
	return result, nil
}
