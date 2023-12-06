package day6

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Race struct {
	Duration       int
	RecordDistance int
}

func Part1(scanner *bufio.Scanner) (string, error) {
	races, err := parseInput(scanner)
	if err != nil {
		return "", err
	}

	product := 1
	for _, r := range races {
		possibilities := foo(r)
		product *= len(possibilities)
	}

	result := strconv.Itoa(product)
	return result, nil
}

func parseInput(scanner *bufio.Scanner) ([]*Race, error) {
	if !scanner.Scan() {
		return nil, errors.New("could not read first line")
	}
	_, rawDurations, found := strings.Cut(strings.TrimSpace(scanner.Text()), ":")
	if !found {
		return nil, errors.New("could not find ':' in first line")
	}

	if !scanner.Scan() {
		return nil, errors.New("could not read second line")
	}
	_, rawRecords, found := strings.Cut(strings.TrimSpace(scanner.Text()), ":")
	if !found {
		return nil, errors.New("could not find ':' in second line")
	}

	splitRawDurations := strings.Fields(rawDurations)
	splitRawRecords := strings.Fields(rawRecords)
	if len(splitRawDurations) != len(splitRawRecords) {
		return nil, fmt.Errorf("got different times / records length (%d / %d)", len(splitRawDurations), len(splitRawRecords))
	}

	races := make([]*Race, len(splitRawDurations))
	for i := 0; i < len(splitRawDurations); i++ {
		rawDuration := splitRawDurations[i]
		rawRecord := splitRawRecords[i]

		duration, err := strconv.Atoi(strings.TrimSpace(rawDuration))
		if err != nil {
			return nil, err
		}
		record, err := strconv.Atoi(strings.TrimSpace(rawRecord))
		if err != nil {
			return nil, err
		}

		races[i] = &Race{
			Duration:       duration,
			RecordDistance: record,
		}
	}

	return races, nil
}

func distanceTravelled(duration, timeHoldingButton int) int {
	// speed = timeHoldingButton
	//
	// distance = speed * time
	// distance = speed * (duration - timeHoldingButton)
	// distance = timeHoldingButton * (duration - timeHoldingButton)
	// distance = thb * duration - thb^2
	return timeHoldingButton*duration - int(math.Pow(float64(timeHoldingButton), 2))
}

func foo(race *Race) []int {
	possibilities := make([]int, 0)
	for i := 0; i <= race.Duration; i++ {
		distance := distanceTravelled(race.Duration, i)
		if distance > race.RecordDistance {
			possibilities = append(possibilities, i)
		}
	}
	return possibilities
}

func Part2(scanner *bufio.Scanner) (string, error) {
	return "", nil
}
