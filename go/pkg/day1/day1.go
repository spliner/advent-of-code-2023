package day1

import (
	"strconv"
	"strings"
	"unicode"
)

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var sum int
	for _, l := range lines {
		l := strings.TrimSpace(l)
		if l == "" {
			continue
		}

		calibration, err := parseCalibration(l)
		if err != nil {
			return "", err
		}

		sum += calibration
	}

	return strconv.Itoa(sum), nil
}

func parseCalibration(input string) (int, error) {
	var firstDigit *rune
	var lastDigit *rune

	for _, r := range input {
		r := r
		if unicode.IsDigit(r) {
			if firstDigit == nil {
				firstDigit = &r
			}
			lastDigit = &r
		}
	}

	var builder strings.Builder
	builder.WriteRune(*firstDigit)
	builder.WriteRune(*lastDigit)

	calibration, err := strconv.Atoi(builder.String())
	if err != nil {
		return 0, err
	}

	return calibration, nil
}

func Part2(input string) (string, error) {
	return "", nil
}
