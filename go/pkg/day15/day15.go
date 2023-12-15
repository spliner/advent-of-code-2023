package day15

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Lens struct {
	Label       string
	FocalLength int
	BoxHash     int
}

func NewLens(label string, focalLength int) *Lens {
	boxHash := Hash(label)
	return &Lens{label, focalLength, boxHash}
}

type Boxes []*Box

func NewBoxes(size int) Boxes {
	boxes := make([]*Box, size)
	for i := 0; i < size; i++ {
		boxes[i] = NewBox()
	}
	return boxes
}

type Box struct {
	Lenses []*Lens
}

func NewBox() *Box {
	lenses := make([]*Lens, 0)
	return &Box{lenses}
}

func (b *Box) AddLens(lens *Lens) {
	for i := 0; i < len(b.Lenses); i++ {
		boxLens := b.Lenses[i]
		if boxLens.Label == lens.Label {
			boxLens.FocalLength = lens.FocalLength
			return
		}
	}
	b.Lenses = append(b.Lenses, lens)
}

func (b *Box) RemoveLens(label string) {
	for i := 0; i < len(b.Lenses); i++ {
		if b.Lenses[i].Label == label {
			b.Lenses = append(b.Lenses[:i], b.Lenses[i+1:]...)
			break
		}
	}
}

func Part1(scanner *bufio.Scanner) (string, error) {
	var sum int
	if !scanner.Scan() {
		return "", nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	line := strings.TrimSpace(scanner.Text())
	split := strings.Split(line, ",")
	for _, str := range split {
		str := strings.TrimSpace(str)
		if str == "" {
			continue
		}

		sum += Hash(str)
	}

	result := strconv.Itoa(sum)
	return result, nil
}

func Hash(str string) int {
	var result int
	for _, r := range str {
		result += int(r)
		result *= 17
		result %= 256
	}
	return result
}

func Part2(scanner *bufio.Scanner) (string, error) {
	if !scanner.Scan() {
		return "", nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	boxes := NewBoxes(256)

	line := strings.TrimSpace(scanner.Text())
	split := strings.Split(line, ",")
	for _, str := range split {
		str := strings.TrimSpace(str)
		if strings.Contains(str, "-") {
			label := str[:len(str)-1]
			boxHash := Hash(label)
			boxes[boxHash].RemoveLens(label)
		} else if strings.Contains(str, "=") {
			label, rawFocalLength, found := strings.Cut(str, "=")
			if !found {
				return "", fmt.Errorf("invalid input: %s", str)
			}

			focalLength, err := strconv.Atoi(rawFocalLength)
			if err != nil {
				return "", err
			}

			lens := NewLens(label, focalLength)
			boxes[lens.BoxHash].AddLens(lens)
		}
	}

	var sum int
	for i, b := range boxes {
		for j, l := range b.Lenses {
			sum += (i + 1) * (j + 1) * l.FocalLength
		}
	}

	result := strconv.Itoa(sum)
	return result, nil
}
