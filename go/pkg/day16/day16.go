package day16

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/spliner/aoc2023/pkg/set"
)

type Direction uint8

const (
	North Direction = iota
	East
	South
	West
)

type Point struct {
	X int
	Y int
}

type Beam struct {
	Position  Point
	Direction Direction
}

type Tile rune

type Grid struct {
	Tiles [][]Tile
}

func (g *Grid) IsInBounds(p Point) bool {
	return p.X >= 0 && p.Y >= 0 && p.Y < len(g.Tiles) && p.X < len(g.Tiles[p.Y])
}

func Part1(scanner *bufio.Scanner) (string, error) {
	grid, err := parseGrid(scanner)
	if err != nil {
		return "", err
	}

	positions := walk(grid, Point{0, 0}, East)
	result := strconv.Itoa(positions)
	return result, nil
}

func walk(grid *Grid, startPosition Point, startDirection Direction) int {
	energizedPositions := set.New[Point]()
	beams := []*Beam{
		{startPosition, startDirection},
	}
	beamDirectionHistory := make(map[Point]*set.Set[Direction])
	for len(beams) > 0 {
		beam := beams[0]
		if _, ok := beamDirectionHistory[beam.Position]; !ok {
			beamDirectionHistory[beam.Position] = set.New[Direction]()
		}
		if !grid.IsInBounds(beam.Position) || beamDirectionHistory[beam.Position].Contains(beam.Direction) {
			if len(beams) > 1 {
				beams = beams[1:]
			} else {
				beams = make([]*Beam, 0)
			}
			continue
		}

		energizedPositions.Add(beam.Position)

		beamDirectionHistory[beam.Position].Add(beam.Direction)

		tile := grid.Tiles[beam.Position.Y][beam.Position.X]
		switch tile {
		case '|':
			if beam.Direction == East || beam.Direction == West {
				beam.Direction = North
				splitBeam := &Beam{Point{beam.Position.X, beam.Position.Y + 1}, South}
				beams = append(beams, splitBeam)
			}
		case '-':
			if beam.Direction == North || beam.Direction == South {
				beam.Direction = East
				splitBeam := Beam{Point{beam.Position.X - 1, beam.Position.Y}, West}
				beams = append(beams, &splitBeam)
			}
		case '/':
			switch beam.Direction {
			case North:
				beam.Direction = East
			case East:
				beam.Direction = North
			case South:
				beam.Direction = West
			case West:
				beam.Direction = South
			}
		case '\\':
			switch beam.Direction {
			case North:
				beam.Direction = West
			case East:
				beam.Direction = South
			case South:
				beam.Direction = East
			case West:
				beam.Direction = North
			}
		}

		beam.Position = nextPoint(beam.Position, beam.Direction)
	}

	return energizedPositions.Length()
}

func print(grid *Grid, beams []*Beam) {
	var sb strings.Builder
	for y, line := range grid.Tiles {
		for x, tile := range line {
			val := rune(tile)
			for _, b := range beams {
				if b.Position.X == x && b.Position.Y == y {
					switch b.Direction {
					case North:
						val = '^'
					case East:
						val = '>'
					case South:
						val = 'v'
					case West:
						val = '<'
					}
					break
				}
			}
			sb.WriteRune(val)
		}
		sb.WriteString("\n")
	}

	fmt.Println(sb.String())
}

func nextPoint(p Point, d Direction) Point {
	switch d {
	case North:
		return Point{p.X, p.Y - 1}
	case East:
		return Point{p.X + 1, p.Y}
	case South:
		return Point{p.X, p.Y + 1}
	default:
		return Point{p.X - 1, p.Y}
	}
}

func parseGrid(scanner *bufio.Scanner) (*Grid, error) {
	grid := make([][]Tile, 0)
	var y int
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		tiles := make([]Tile, len(line))
		for x, r := range line {
			tiles[x] = Tile(r)
		}

		grid = append(grid, tiles)
		y++
	}

	return &Grid{grid}, nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	grid, err := parseGrid(scanner)
	if err != nil {
		return "", err
	}

	height := len(grid.Tiles)
	width := len(grid.Tiles[0])

	var maxCount int
	for x := 0; x < width; x++ {
		startPoint := Point{x, 0}
		count := walk(grid, startPoint, South)
		if count > maxCount {
			maxCount = count
		}

		startPoint = Point{x, height - 1}
		count = walk(grid, startPoint, North)
		if count > maxCount {
			maxCount = count
		}
	}

	for y := 0; y < height; y++ {
		startPoint := Point{0, y}
		count := walk(grid, startPoint, East)
		if count > maxCount {
			maxCount = count
		}

		startPoint = Point{width - 1, y}

		count = walk(grid, startPoint, West)
		if count > maxCount {
			maxCount = count
		}
	}

	result := strconv.Itoa(maxCount)
	return result, nil
}
