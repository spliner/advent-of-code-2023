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

func (u *Universe) Expand() *Universe {
	emptyColumns, emptyRows := emptyIndexes(u)

	oldHeight := len(u.Map)
	oldWidth := len(u.Map[0])
	newHeight := oldHeight + emptyRows.Length()
	newWidth := oldWidth + emptyColumns.Length()

	galaxies := make([]Galaxy, len(u.Galaxies))
	var galaxyIndex int

	newMap := make([][]rune, newHeight)
	var oldY, y int
	for y < newHeight {
		if emptyRows.Contains(oldY) {
			newMap[y] = make([]rune, newWidth)
			newMap[y+1] = make([]rune, newWidth)
			for x := 0; x < newWidth; x++ {
				val := '.'
				newMap[y][x] = val
				newMap[y+1][x] = val
			}

			y++
		} else {
			var oldX, x int
			row := make([]rune, newWidth)
			for x < newWidth {
				if emptyColumns.Contains(oldX) {
					row[x] = '.'
					x++
					row[x] = '.'
				} else {
					val := u.Map[oldY][oldX]
					row[x] = val
					if val == '#' {
						position := Point{x, y}
						galaxies[galaxyIndex] = Galaxy{position}
						galaxyIndex++
					}
				}

				x++
				oldX++
			}

			newMap[y] = row
		}

		oldY++
		y++
	}

	return &Universe{newMap, galaxies}
}

func emptyIndexes(u *Universe) (emtptyColumns, emptyRows *set.Set[int]) {
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
	universe, err := parseUniverse(scanner)
	if err != nil {
		return "", nil
	}

	expandedUniverse := universe.Expand()
	var sum int
	for i := 0; i < len(expandedUniverse.Galaxies); i++ {
		for j := i + 1; j < len(expandedUniverse.Galaxies); j++ {
			gX := expandedUniverse.Galaxies[i]
			gY := expandedUniverse.Galaxies[j]

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
	return "", nil
}
