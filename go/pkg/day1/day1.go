package day1

import (
	"strconv"
	"strings"
	"unicode"
)

var replaces = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func Part1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	var sum int
	for _, l := range lines {
		l := strings.TrimSpace(l)
		if l == "" {
			continue
		}

		calibration, err := parsePart1Calibration(l)
		if err != nil {
			return "", err
		}

		sum += calibration
	}

	return strconv.Itoa(sum), nil
}

func parsePart1Calibration(input string) (int, error) {
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
	lines := strings.Split(input, "\n")
	var sum int
	for _, l := range lines {
		l := strings.TrimSpace(l)
		if l == "" {
			continue
		}

		calibration, err := parsePart2Calibration(l)
		if err != nil {
			return "", err
		}

		sum += calibration
	}

	return strconv.Itoa(sum), nil
}

func parsePart2Calibration(input string) (int, error) {
	runes := []rune(input)
	var firstDigit *rune
	var lastDigit *rune
	for i := 0; i < len(runes); {
		if unicode.IsDigit(runes[i]) {
			if firstDigit == nil {
				firstDigit = &runes[i]
			}
			lastDigit = &runes[i]

			i++
			continue
		}

		substr := string(runes[i:])
		for k, v := range replaces {
			if strings.HasPrefix(substr, k) {
				v := v
				if firstDigit == nil {
					firstDigit = &v
				}
				lastDigit = &v
			}
		}

		i++
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
