package day18

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func Part1(scanner *bufio.Scanner) (string, error) {
	instructions, err := parseInstructions(scanner)
	if err != nil {
		return "", err
	}
	fmt.Println(instructions)
	return "", nil
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

func Part2(scanner *bufio.Scanner) (string, error) {
	return "", nil
}
