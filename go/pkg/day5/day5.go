package day5

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MapLine struct {
	DestinationStart int
	SourceStart      int
	Length           int
}

func NewMapLine(destinationStart, sourceStart, length int) *MapLine {
	return &MapLine{
		DestinationStart: destinationStart,
		SourceStart:      sourceStart,
		Length:           length,
	}
}

func (l *MapLine) Destination(source int) (int, bool) {
	if source < l.SourceStart {
		return 0, false
	}

	diff := source - l.SourceStart
	if diff > l.Length-1 {
		return 0, false
	}

	return l.DestinationStart + diff, true
}

type Map struct {
	Lines []*MapLine
}

func NewMap(lines []*MapLine) *Map {
	return &Map{lines}
}

func (m *Map) Destination(source int) int {
	for _, line := range m.Lines {
		if destination, ok := line.Destination(source); ok {
			return destination
		}
	}

	return source
}

type Almanac struct {
	Maps  []*Map
	Seeds []int
}

func (a *Almanac) SeedLocation(seed int) int {
	dest := seed
	for _, m := range a.Maps {
		dest = m.Destination(dest)
	}
	return dest
}

func Part1(scanner *bufio.Scanner) (string, error) {
	almanac, err := parseAlmanac(scanner)
	if err != nil {
		return "", err
	}

	var min *int
	for _, seed := range almanac.Seeds {
		location := almanac.SeedLocation(seed)
		if min == nil || location < *min {
			min = &location
		}
	}

	result := strconv.Itoa(*min)
	return result, nil
}

func parseAlmanac(scanner *bufio.Scanner) (*Almanac, error) {
	// seeds: 79 14 55 13
	if !scanner.Scan() {
		return nil, errors.New("could not read first input line")
	}

	line := scanner.Text()
	_, rawSeeds, found := strings.Cut(line, ":")
	if !found {
		return nil, errors.New("missing ':' in seeds line")
	}

	seeds := make([]int, 0)
	for _, s := range strings.Split(rawSeeds, " ") {
		s := strings.TrimSpace(s)
		if s == "" {
			continue
		}

		val, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}

		seeds = append(seeds, val)
	}

	// Empty line
	if !scanner.Scan() {
		return nil, errors.New("could not read second input line")
	}

	maps := make([]*Map, 7)
	for i := 0; i < 7; i++ {
		m, err := parseMap(scanner)
		if err != nil {
			return nil, err
		}
		maps[i] = m
	}

	almanac := Almanac{
		Seeds: seeds,
		Maps:  maps,
	}
	return &almanac, nil
}

func parseMap(scanner *bufio.Scanner) (*Map, error) {
	if !scanner.Scan() {
		return nil, errors.New("could not read map header")
	}

	lines := make([]*MapLine, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break
		}

		// 50 98 2
		split := strings.Split(line, " ")
		if len(split) != 3 {
			return nil, fmt.Errorf("invalid map line input: '%s'", line)
		}

		dest, err := strconv.Atoi(strings.TrimSpace(split[0]))
		if err != nil {
			return nil, err
		}

		source, err := strconv.Atoi(strings.TrimSpace(split[1]))
		if err != nil {
			return nil, err
		}

		length, err := strconv.Atoi(strings.TrimSpace(split[2]))
		if err != nil {
			return nil, err
		}

		lines = append(lines, NewMapLine(dest, source, length))
	}

	m := NewMap(lines)
	return m, nil
}

func Part2(scanner *bufio.Scanner) (string, error) {
	almanac, err := parseAlmanac(scanner)
	if err != nil {
		return "", err
	}

	pairCount := len(almanac.Seeds) / 2
	minChan := make(chan int, pairCount)
	for i := 0; i < len(almanac.Seeds); i += 2 {
		start := almanac.Seeds[i]
		end := start + almanac.Seeds[i+1]

		go func(start, end int) {
			var min *int
			for j := start; j < end; j++ {
				location := almanac.SeedLocation(j)
				if min == nil || location < *min {
					min = &location
				}
			}
			minChan <- *min
		}(start, end)
	}

	var min *int
	for i := 0; i < pairCount; i++ {
		m := <-minChan
		if min == nil || m < *min {
			min = &m
		}
	}

	close(minChan)

	result := strconv.Itoa(*min)
	return result, nil
}
