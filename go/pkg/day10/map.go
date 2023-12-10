package day10

import "strings"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) IsVertical() bool {
	return d == North || d == South
}

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	}
	return "Unknown"
}

type Position struct {
	X int
	Y int
}
type Tile struct {
	Value    rune
	Position Position
	Pipe     Pipe
}

func NewTile(value rune, pos Position) Tile {
	tile := Tile{
		Value:    value,
		Position: pos,
	}

	if value == '|' {
		tile.Pipe = &VerticalPipe{}
	} else if value == '-' {
		tile.Pipe = &HorizontalPipe{}
	} else if value == 'L' {
		tile.Pipe = &NorthEastBendPipe{}
	} else if value == 'J' {
		tile.Pipe = &NorthWestBendPipe{}
	} else if value == '7' {
		tile.Pipe = &SouthWestBendPipe{}
	} else if value == 'F' {
		tile.Pipe = &SouthEastBendPipe{}
	}

	return tile
}

func (t Tile) String() string {
	return string(t.Value)
}

type Map struct {
	Start Position
	Tiles [][]Tile
}

func (m *Map) String() string {
	var sb strings.Builder
	for _, line := range m.Tiles {
		for _, tile := range line {
			sb.WriteRune(tile.Value)
		}
		sb.WriteString("\n")
	}
	return sb.String()
}
