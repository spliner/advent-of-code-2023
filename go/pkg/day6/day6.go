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
	races, err := parsePart1Input(scanner)
	if err != nil {
		return "", err
	}

	product := 1
	for _, r := range races {
		possibilities := countWinPossibilities(r)
		product *= possibilities
	}

	result := strconv.Itoa(product)
	return result, nil
}

func parsePart1Input(scanner *bufio.Scanner) ([]*Race, error) {
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

func countWinPossibilities(race *Race) int {
	var possibilities int
	for i := 0; i <= race.Duration; i++ {
		distance := distanceTravelled(race.Duration, i)
		if distance > race.RecordDistance {
			possibilities++
		}
	}
	return possibilities
}

func Part2(scanner *bufio.Scanner) (string, error) {
	race, err := parsePart2Input(scanner)
	if err != nil {
		return "", err
	}

	possibilities := countWinPossibilities(race)
	result := strconv.Itoa(possibilities)
	return result, nil
}

func parsePart2Input(scanner *bufio.Scanner) (*Race, error) {
	if !scanner.Scan() {
		return nil, errors.New("could not read first line")
	}
	_, rawDuration, found := strings.Cut(strings.TrimSpace(scanner.Text()), ":")
	if !found {
		return nil, errors.New("could not find ':' in first line")
	}
	rawDuration = strings.ReplaceAll(rawDuration, " ", "")

	if !scanner.Scan() {
		return nil, errors.New("could not read second line")
	}
	_, rawRecord, found := strings.Cut(strings.TrimSpace(scanner.Text()), ":")
	if !found {
		return nil, errors.New("could not find ':' in second line")
	}
	rawRecord = strings.ReplaceAll(rawRecord, " ", "")
	duration, err := strconv.Atoi(strings.TrimSpace(rawDuration))
	if err != nil {
		return nil, err
	}
	record, err := strconv.Atoi(strings.TrimSpace(rawRecord))
	if err != nil {
		return nil, err
	}

	race := &Race{
		Duration:       duration,
		RecordDistance: record,
	}
	return race, nil
}
