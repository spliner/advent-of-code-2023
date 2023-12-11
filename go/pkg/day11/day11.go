package day11

import (
	"bufio"
	"math"
	"strconv"
	"strings"

	"github.com/spliner/aoc2023/pkg/set"
)

type Point struct {
	X int
	Y int
}

func (p Point) Distance(other Point) int {
	distance := math.Abs(float64(other.Y-p.Y)) + math.Abs(float64(other.X-p.X))
	return int(distance)
}

type Galaxy struct {
	Position Point
}

type Universe struct {
	Map      [][]rune
	Galaxies []Galaxy
}

func (u *Universe) Expand(expandTimes int) []Galaxy {
	expandTimes = expandTimes - 1

	expandedGalaxies := make([]Galaxy, len(u.Galaxies))
	emptyColumnSet, emptyRowSet := u.emptyIndexes()
	emptyColumns := emptyColumnSet.All()
	emptyRows := emptyRowSet.All()
	for i, g := range u.Galaxies {
		x := g.Position.X
		var columnCount int
		for _, c := range emptyColumns {
			if c < x {
				columnCount++
			}
		}
		x += columnCount * expandTimes

		y := g.Position.Y
		var rowCount int
		for _, c := range emptyRows {
			if c < y {
				rowCount++
			}
		}

		y += rowCount * expandTimes

		position := Point{x, y}
		expandedGalaxies[i] = Galaxy{position}
	}

	return expandedGalaxies
}

func (u *Universe) emptyIndexes() (emtptyColumns, emptyRows *set.Set[int]) {
	emtptyColumns = set.New[int]()
	emptyRows = set.New[int]()

	height := len(u.Map)
	width := len(u.Map[0])

	for y := 0; y < height; y++ {
		emptyRow := true
		for x := 0; x < width; x++ {
			if u.Map[y][x] != '.' {
				emptyRow = false
				break
			}
		}
		if emptyRow {
			emptyRows.Add(y)
		}
	}

	for x := 0; x < width; x++ {
		emptyColumn := true
		for y := 0; y < height; y++ {
			if u.Map[y][x] != '.' {
				emptyColumn = false
				break
			}
		}

		if emptyColumn {
			emtptyColumns.Add(x)
		}
	}

	return
}

func Part1(scanner *bufio.Scanner) (string, error) {
	return run(scanner, 2)
}

func run(scanner *bufio.Scanner, expandTimes int) (string, error) {
	universe, err := parseUniverse(scanner)
	if err != nil {
		return "", nil
	}

	expandedGalaxies := universe.Expand(expandTimes)
	var sum int
	for i := 0; i < len(expandedGalaxies); i++ {
		for j := i + 1; j < len(expandedGalaxies); j++ {
			gX := expandedGalaxies[i]
			gY := expandedGalaxies[j]

			distance := gX.Position.Distance(gY.Position)

			sum += distance
		}
	}

	result := strconv.Itoa(sum)
	return result, nil
}

func parseUniverse(scanner *bufio.Scanner) (*Universe, error) {
	lines := make([][]rune, 0)
	galaxies := make([]Galaxy, 0)
	var y int
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		rawLine := strings.TrimSpace(scanner.Text())
		line := make([]rune, len(rawLine))
		for x, r := range rawLine {
			line[x] = r

			if r == '#' {
				p := Point{x, y}
				galaxy := Galaxy{p}
				galaxies = append(galaxies, galaxy)
			}
		}

		lines = append(lines, line)
		y++
	}

	universe := Universe{lines, galaxies}
	return &universe, nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	return run(scanner, 1_000_000)
}
