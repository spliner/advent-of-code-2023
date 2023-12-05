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
	SeedToSoilMap            *Map
	SoilToFertilizerMap      *Map
	FertilizerToWaterMap     *Map
	WaterToLightMap          *Map
	LightToTemperatureMap    *Map
	TemperatureToHumidityMap *Map
	HumidityToLocationMap    *Map
	Seeds                    []int
}

func (a *Almanac) SeedLocation(seed int) int {
	soil := a.SeedToSoilMap.Destination(seed)
	fertilizer := a.SoilToFertilizerMap.Destination(soil)
	water := a.FertilizerToWaterMap.Destination(fertilizer)
	light := a.WaterToLightMap.Destination(water)
	temperature := a.LightToTemperatureMap.Destination(light)
	humidity := a.TemperatureToHumidityMap.Destination(temperature)
	location := a.HumidityToLocationMap.Destination(humidity)

	return location
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

	almanac := Almanac{
		Seeds: seeds,
	}

	// Empty line
	if !scanner.Scan() {
		return nil, errors.New("could not read second input line")
	}

	seedToSoilMap, err := parseMap(scanner)
	if err != nil {
		return nil, err
	}
	almanac.SeedToSoilMap = seedToSoilMap

	soilToFertilizerMap, err := parseMap(scanner)
	if err != nil {
		return nil, err
	}
	almanac.SoilToFertilizerMap = soilToFertilizerMap

	fertilizerToWaterMap, err := parseMap(scanner)
	if err != nil {
		return nil, err
	}
	almanac.FertilizerToWaterMap = fertilizerToWaterMap

	waterToLightMap, err := parseMap(scanner)
	if err != nil {
		return nil, err
	}
	almanac.WaterToLightMap = waterToLightMap

	lightToTemperatureMap, err := parseMap(scanner)
	if err != nil {
		return nil, err
	}
	almanac.LightToTemperatureMap = lightToTemperatureMap

	temperatureToHumidityMap, err := parseMap(scanner)
	if err != nil {
		return nil, err
	}
	almanac.TemperatureToHumidityMap = temperatureToHumidityMap

	humidityToLocationMap, err := parseMap(scanner)
	if err != nil {
		return nil, err
	}
	almanac.HumidityToLocationMap = humidityToLocationMap

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
