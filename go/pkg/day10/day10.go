package day10

import (
	"bufio"
	"fmt"
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

func Part2(scanner *bufio.Scanner) (string, error) {
	m, err := parseMap(scanner)
	if err != nil {
		return "", err
	}

	loop := findLoop(m)
	loop[m.Start] = 0
	validPositions := make([]Position, 0)
	for y := 1; y < len(m.Tiles); y++ {
		var wallCount int
		var lastBend rune
		for x := 1; x < len(m.Tiles[y]); x++ {
			tile := m.Tiles[y][x]
			if _, ok := loop[tile.Position]; ok {
				switch tile.Value {
				case '|':
					wallCount++
					lastBend = 0
				case 'L', 'F':
					lastBend = tile.Value
				case 'J':
					if lastBend == 'F' {
						wallCount++
					}
					lastBend = 0
				case '7':
					if lastBend == 'L' {
						wallCount++
					}
					lastBend = 0
				}
			} else {
				lastBend = 0
				if wallCount%2 == 1 {
					validPositions = append(validPositions, tile.Position)
				}
			}
		}
	}

	result := strconv.Itoa(len(validPositions))
	return result, nil
}

func printMap(m *Map, validPositions []Position) {
	loop := findLoop(m)
	loop[m.Start] = 0

	var sb strings.Builder
	for y := 0; y < len(m.Tiles); y++ {
		line := m.Tiles[y]
		for x := 0; x < len(line); x++ {
			tile := m.Tiles[y][x]
			val := tile.Value
			_, ok := loop[tile.Position]
			if !ok {
				for _, p := range validPositions {
					if p == tile.Position {
						val = 'I'
						break
					} else {
						val = 'O'
					}
				}
			}
			if x == 0 || y == 0 || x == len(m.Tiles[y])-1 || y == len(m.Tiles)-1 {
				val = '.'
			}
			sb.WriteRune(val)
		}
		sb.WriteString("\n")
	}
	fmt.Println(sb.String())
}

func findLoop(m *Map) map[Position]int {
	start := m.Start

	steps, hasLoop := findSteps(m, Position{start.X + 1, start.Y}, East)
	if hasLoop {
		return steps
	}

	steps, hasLoop = findSteps(m, Position{start.X, start.Y + 1}, South)
	if hasLoop {
		return steps
	}

	steps, hasLoop = findSteps(m, Position{start.X - 1, start.Y}, West)
	if hasLoop {
		return steps
	}

	steps, hasLoop = findSteps(m, Position{start.X, start.Y - 1}, North)
	if hasLoop {
		return steps
	}

	return make(map[Position]int)
}
