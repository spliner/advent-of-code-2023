package day3

import (
	"strconv"
	"strings"
	"unicode"
)

type Point struct {
	X int
	Y int
}

type Schematic struct {
	Coordinates    [][]Coordinate
	NumberMap      map[Point]Number
	Numbers        []Number
	GearCandidates []Point
}

func (s *Schematic) IsPartNumber(number Number) bool {
	for _, coordinate := range number.Coordinates {
		x := coordinate.Point.X
		y := coordinate.Point.Y
		if s.IsSymbol(x-1, y-1) ||
			s.IsSymbol(x-1, y) ||
			s.IsSymbol(x-1, y+1) ||
			s.IsSymbol(x, y) ||
			s.IsSymbol(x, y-1) ||
			s.IsSymbol(x, y+1) ||
			s.IsSymbol(x+1, y-1) ||
			s.IsSymbol(x+1, y) ||
			s.IsSymbol(x+1, y+1) {
			return true
		}
	}

	return false
}

func (s *Schematic) IsSymbol(x, y int) bool {
	coordinates := s.Coordinates
	return x >= 0 && x < len(coordinates) && y >= 0 && y < len(coordinates) && coordinates[y][x].IsSymbol()
}

func (s *Schematic) FindAdjacentPartNumbers(gear Point) []int {
	partNumberMap := make(map[int]struct{})
	points := []Point{
		{gear.X - 1, gear.Y - 1},
		{gear.X - 1, gear.Y},
		{gear.X - 1, gear.Y + 1},
		{gear.X, gear.Y},
		{gear.X, gear.Y - 1},
		{gear.X, gear.Y + 1},
		{gear.X + 1, gear.Y - 1},
		{gear.X + 1, gear.Y},
		{gear.X + 1, gear.Y + 1},
	}
	for _, p := range points {
		n, ok := s.NumberMap[p]
		if ok {
			partNumberMap[n.Value] = struct{}{}
		}
	}
	partNumbers := make([]int, 0, len(partNumberMap))
	for k := range partNumberMap {
		partNumbers = append(partNumbers, k)
	}
	return partNumbers
}

func (s *Schematic) IsDigit(x, y int) bool {
	coordinates := s.Coordinates
	return x >= 0 && x < len(coordinates) && y >= 0 && y < len(coordinates) && unicode.IsDigit(coordinates[y][x].Value)
}

type Coordinate struct {
	Value rune
	Point Point
}

func (c Coordinate) IsSymbol() bool {
	return c.Value != '.' && !unicode.IsDigit(c.Value)
}

type Number struct {
	Coordinates []Coordinate
	Value       int
}

func Part1(input string) (string, error) {
	schematic, err := parseSchematic(input)
	if err != nil {
		return "", err
	}

	var sum int
	for _, number := range schematic.Numbers {
		if schematic.IsPartNumber(number) {
			sum += number.Value
		}
	}
	result := strconv.Itoa(sum)
	return result, nil
}

func parseSchematic(input string) (*Schematic, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	coordinates := make([][]Coordinate, len(lines))
	numbers := make([]Number, 0)
	numberMap := make(map[Point]Number)
	gearCandidates := make([]Point, 0)
	for y, line := range lines {
		line := strings.TrimSpace(line)
		coordinates[y] = make([]Coordinate, len(line))

		var currentNumber strings.Builder
		currentNumberCoordinates := make([]Coordinate, 0)
		for x, r := range line {
			coordinate := Coordinate{
				Value: r,
				Point: Point{
					X: x,
					Y: y,
				},
			}

			if unicode.IsDigit(r) {
				currentNumberCoordinates = append(currentNumberCoordinates, coordinate)
				currentNumber.WriteRune(r)
			} else if currentNumber.Len() > 0 {
				numStr := currentNumber.String()
				num, err := strconv.Atoi(numStr)
				if err != nil {
					return nil, err
				}

				number := Number{Value: num, Coordinates: currentNumberCoordinates}
				numbers = append(numbers, number)
				for _, c := range currentNumberCoordinates {
					numberMap[c.Point] = number
				}
				currentNumber.Reset()
				currentNumberCoordinates = make([]Coordinate, 0)
			}

			if r == '*' {
				gearCandidates = append(gearCandidates, coordinate.Point)
			}

			coordinates[y][x] = coordinate
		}
		if currentNumber.Len() > 0 {
			numStr := currentNumber.String()
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, err
			}
			number := Number{Value: num, Coordinates: currentNumberCoordinates}
			numbers = append(numbers, number)
			for _, c := range number.Coordinates {
				numberMap[c.Point] = number
			}
		}
	}

	schematic := Schematic{
		Coordinates:    coordinates,
		Numbers:        numbers,
		NumberMap:      numberMap,
		GearCandidates: gearCandidates,
	}

	return &schematic, nil
}

func Part2(input string) (string, error) {
	schematic, err := parseSchematic(input)
	if err != nil {
		return "", err
	}

	var sum int
	for _, g := range schematic.GearCandidates {
		adjacentPartNumbers := schematic.FindAdjacentPartNumbers(g)
		if len(adjacentPartNumbers) == 2 {
			sum += adjacentPartNumbers[0] * adjacentPartNumbers[1]
		}
	}

	result := strconv.Itoa(sum)
	return result, nil
}
