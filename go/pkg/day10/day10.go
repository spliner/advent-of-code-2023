package day10

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func Part1(scanner *bufio.Scanner) (string, error) {
	m, err := parseMap(scanner)
	if err != nil {
		return "", err
	}
	start := m.Start

	stepMaps := make([]map[Position]int, 0)
	steps, hasLoop := findSteps(m, Position{start.X + 1, start.Y}, East)
	if hasLoop {
		stepMaps = append(stepMaps, steps)
	}
	steps, hasLoop = findSteps(m, Position{start.X, start.Y + 1}, South)
	if hasLoop {
		stepMaps = append(stepMaps, steps)
	}
	steps, hasLoop = findSteps(m, Position{start.X - 1, start.Y}, West)
	if hasLoop {
		stepMaps = append(stepMaps, steps)
	}
	steps, hasLoop = findSteps(m, Position{start.X, start.Y - 1}, North)
	if hasLoop {
		stepMaps = append(stepMaps, steps)
	}

	var maxSteps int
	for k := range stepMaps[0] {
		minSteps := math.MaxInt
		for _, s := range stepMaps {
			val := s[k]
			if val < minSteps {
				minSteps = val
			}
		}

		if minSteps > maxSteps {
			maxSteps = minSteps
		}
	}

	result := strconv.Itoa(maxSteps)
	return result, nil
}

func parseMap(scanner *bufio.Scanner) (*Map, error) {
	tiles := [][]Tile{
		make([]Tile, 0),
	}

	var start Position
	y := 1
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())
		// Expand map to avoid comparing out of bounds
		line = "." + line + "."
		lineTiles := make([]Tile, len(line))
		for x, r := range line {
			position := Position{x, y}

			if r == 'S' {
				start = position
			}

			lineTiles[x] = NewTile(r, position)
		}

		tiles = append(tiles, lineTiles)
		y++
	}

	width := len(tiles[1])
	tiles = append(tiles, make([]Tile, width))
	for i := 0; i < width; i++ {
		tiles[0] = append(tiles[0], NewTile('.', Position{0, i}))
		tiles[len(tiles)-1][0] = NewTile('.', Position{len(tiles) - 1, i})
	}

	m := Map{
		Start: start,
		Tiles: tiles,
	}
	return &m, nil
}

func findSteps(m *Map, pos Position, d Direction) (map[Position]int, bool) {
	stepMap := make(map[Position]int)
	tile := m.Tiles[pos.Y][pos.X]
	var steps int
	dir := d
	for tile.Pipe != nil && tile.Position != m.Start {
		steps++
		stepMap[pos] = steps
		oldDir := dir
		pos, dir = tile.Pipe.NextPosition(tile.Position, oldDir)
		nextTile := m.Tiles[pos.Y][pos.X]
		if nextTile.Position != m.Start && !tile.Pipe.CanConnect(nextTile, oldDir) {
			return nil, false
		}
		tile = nextTile
	}
	return stepMap, tile.Position == m.Start
}
