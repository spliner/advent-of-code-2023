package day18

import (
	"bufio"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/spliner/aoc2023/pkg/set"
)

type Instruction struct {
	Direction Direction
	Steps     int
	Color     Color
}

type Direction uint8

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type Point struct {
	X int
	Y int
}

type State struct {
	Current   Point
	Points    []Point
	Perimeter uint
}

func NewState() *State {
	return &State{
		Current: Point{0, 0},
		Points:  make([]Point, 0),
	}
}

func Part1(scanner *bufio.Scanner) (string, error) {
	instructions, err := parseInstructions(scanner)
	if err != nil {
		return "", err
	}

	grid := outline(instructions)
	printGrid(grid)
	a := area(grid)
	// printGrid(grid)

	result := strconv.FormatUint(uint64(a), 10)
	return result, nil
}

func parseInstructions(scanner *bufio.Scanner) ([]Instruction, error) {
	r, err := regexp.Compile(`^([URDL]) (\d+) \(#(\w{6})\)$`)
	if err != nil {
		return nil, err
	}

	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		matches := r.FindStringSubmatch(line)
		if len(matches) != 4 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		direction, err := parseDirection(matches[1])
		if err != nil {
			return nil, err
		}

		steps, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, err
		}

		color, err := parseColor(matches[3])
		if err != nil {
			return nil, err
		}

		instruction := Instruction{direction, steps, color}
		instructions = append(instructions, instruction)
	}

	return instructions, nil
}

func parseDirection(raw string) (Direction, error) {
	switch raw {
	case "U":
		return Up, nil
	case "L":
		return Left, nil
	case "D":
		return Down, nil
	case "R":
		return Right, nil
	}

	return 0, fmt.Errorf("invalid direction: %s", raw)
}

func parseColor(raw string) (Color, error) {
	if len(raw) != 6 {
		return Color{}, fmt.Errorf("invalid color: %s", raw)
	}

	red, err := strconv.ParseInt(raw[0:2], 16, 0)
	if err != nil {
		return Color{}, fmt.Errorf("invalid color: %s", raw)
	}

	green, err := strconv.ParseInt(raw[2:4], 16, 0)
	if err != nil {
		return Color{}, fmt.Errorf("invalid color: %s", raw)
	}

	blue, err := strconv.ParseInt(raw[4:6], 16, 0)
	if err != nil {
		return Color{}, fmt.Errorf("invalid color: %s", raw)
	}

	return Color{uint8(red), uint8(green), uint8(blue)}, nil
}

func outline(instructions []Instruction) [][]rune {
	var upCount, rightCount, downCount, leftCount int
	currentPoint := Point{0, 0}
	orderedPoints := make([]Point, 0)
	pointSet := set.New[Point]()
	var maxY, maxX int
	for _, instruction := range instructions {
		switch instruction.Direction {
		case Up:
			upCount += instruction.Steps
			for i := 1; i <= instruction.Steps; i++ {
				point := Point{currentPoint.X, currentPoint.Y - i}
				orderedPoints = append(orderedPoints, point)
				pointSet.Add(point)
			}
		case Right:
			rightCount += instruction.Steps
			for i := 1; i <= instruction.Steps; i++ {
				point := Point{currentPoint.X + i, currentPoint.Y}
				orderedPoints = append(orderedPoints, point)
				pointSet.Add(point)
			}
		case Down:
			downCount += instruction.Steps
			for i := 1; i <= instruction.Steps; i++ {
				point := Point{currentPoint.X, currentPoint.Y + i}
				orderedPoints = append(orderedPoints, point)
				pointSet.Add(point)
			}
		case Left:
			leftCount += instruction.Steps
			for i := 1; i <= instruction.Steps; i++ {
				point := Point{currentPoint.X - i, currentPoint.Y}
				orderedPoints = append(orderedPoints, point)
				pointSet.Add(point)
			}
		}

		lastPoint := orderedPoints[len(orderedPoints)-1]
		if lastPoint.Y > maxY {
			maxY = lastPoint.Y
		}
		if lastPoint.X > maxX {
			maxX = lastPoint.X
		}

		currentPoint = orderedPoints[len(orderedPoints)-1]
	}

	fmt.Println(orderedPoints)

	height := int(math.Abs(float64(upCount) - float64(downCount)))
	width := int(math.Abs(float64(leftCount) - float64(rightCount)))

	grid := make([][]rune, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]rune, width)
		for x := 0; x < width; x++ {
			p := Point{x, y}
			var val rune
			if pointSet.Contains(p) {
				val = '#'
			} else {
				val = '.'
			}
			grid[y][x] = val
		}
	}

	return grid
}

func area(grid [][]rune) uint {
	var area uint
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			tile := grid[y][x]
			if tile == '.' {
				var count int
				for j := x + 1; j < len(grid[y]); j++ {
					other := grid[y][j]
					if other == '#' {
						count++
					}
				}
				if count%2 == 1 {
					grid[y][x] = '#'
					area++
				}
			} else {
				area++
			}
		}
	}
	return area
}

func printGrid(grid [][]rune) {
	var sb strings.Builder
	for _, line := range grid {
		for _, r := range line {
			sb.WriteRune(r)
		}
		sb.WriteRune('\n')
	}

	fmt.Println(sb.String())
}

func Part2(scanner *bufio.Scanner) (string, error) {
	return "", nil
}
