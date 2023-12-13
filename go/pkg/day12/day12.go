package day12

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Record struct {
	Line           string
	DamagedSprings []int
}

func (r Record) Expand(times int) Record {
	var line strings.Builder
	damagedSprings := make([]int, 0)

	for i := 0; i < times; i++ {
		line.WriteString(r.Line)
		if i < times-1 {
			line.WriteRune('?')
		}
		damagedSprings = append(damagedSprings, r.DamagedSprings...)
	}

	expanded := Record{line.String(), damagedSprings}
	return expanded
}

func Part1(scanner *bufio.Scanner) (string, error) {
	records, err := parseRecords(scanner)
	if err != nil {
		return "", err
	}

	var sum int
	for _, r := range records {
		validArrangements := validArrangements(&r)
		sum += len(validArrangements)
	}
	result := strconv.Itoa(sum)
	return result, nil
}

func parseRecords(scanner *bufio.Scanner) ([]Record, error) {
	records := make([]Record, 0)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		line := strings.TrimSpace(scanner.Text())
		rawLine, rawDamagedSprings, found := strings.Cut(line, " ")
		if !found {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		rawLine = strings.TrimSpace(rawLine)

		springSplit := strings.Split(rawDamagedSprings, ",")
		damagedSprings := make([]int, len(springSplit))
		for i, s := range springSplit {
			damagedSpring, err := strconv.Atoi(strings.TrimSpace(s))
			if err != nil {
				return nil, err
			}

			damagedSprings[i] = damagedSpring
		}

		record := Record{rawLine, damagedSprings}
		records = append(records, record)
	}

	return records, nil
}

func validArrangements(record *Record) []string {
	line := record.Line
	arrangements := possibleArrangements(line)
	validArrangements := make([]string, 0)
	for _, a := range arrangements {
		if validArrangement(a, record.DamagedSprings) {
			validArrangements = append(validArrangements, a)
		}
	}
	return validArrangements
}

type unknownGroup struct {
	startIndex int
	val        string
}

func possibleArrangements(line string) []string {
	// Maps start index -> group
	unknownGroups := make([]unknownGroup, 0)
	g := unknownGroup{}
	var currentGroup strings.Builder
	index := -1
	for i, r := range line {
		if r == '?' {
			if index == -1 {
				index = i
			}
			currentGroup.WriteRune(r)
		}
		if r != '?' || i == len(line)-1 {
			if index != -1 {
				g.startIndex = index
				g.val = currentGroup.String()
				unknownGroups = append(unknownGroups, g)
			}
			g = unknownGroup{}
			currentGroup.Reset()
			index = -1
		}
	}
	arrangements := []string{line}
	for _, group := range unknownGroups {
		arrangementsToReplace := make([]string, 0)
		start, end := -1, -1
		for i := 0; i < len(arrangements); i++ {
			a := arrangements[i]
			if a[group.startIndex] == '?' {
				if start == -1 {
					start = i
				}
				end = i
				arrangementsToReplace = append(arrangementsToReplace, a)
			}
		}
		arrangements = append(arrangements[:start], arrangements[end+1:]...)

		groupedArrangements := groupArrangements(group.val)
		for _, r := range arrangementsToReplace {
			for _, a := range groupedArrangements {
				arrangement := strings.Replace(r, group.val, a, 1)
				arrangements = append(arrangements, arrangement)
			}
		}
	}

	return arrangements
}

var cache = map[string][]string{
	"?":   {".", "#"},
	"??":  {"..", ".#", "#.", "##"},
	"???": {"...", "..#", ".#.", ".##", "#..", "#.#", "##.", "###"},
}

func groupArrangements(s string) []string {
	if val, ok := cache[s]; ok {
		return val
	}

	if len(s) == 1 {
		return []string{".", "#"}
	}

	possibilities := int(math.Pow(2, float64(len(s))))
	results := make([]string, possibilities)
	half := possibilities / 2
	for i := 0; i < possibilities; i++ {
		var val string
		if i < half {
			val = "."
		} else {
			val = "#"
		}
		results[i] = val
	}

	f := groupArrangements(s[1:])
	for i := 0; i < possibilities; i++ {
		idx := i % half
		results[i] = results[i] + f[idx]
	}
	cache[s] = results
	return results
}

var groupCountMemo map[string][]int = make(map[string][]int)

func validArrangement(arrangement string, damagedGroups []int) bool {
	groupCounts := groupCountMemo[arrangement]
	if groupCounts == nil {
		groupCounts = make([]int, 0, len(arrangement))
	}
	var currentCount int
	for _, r := range arrangement {
		if r == '#' {
			currentCount++
		} else if r == '.' && currentCount > 0 {
			groupCounts = append(groupCounts, currentCount)
			currentCount = 0
		}
	}

	groupCountMemo[arrangement] = groupCounts
	if currentCount > 0 {
		groupCounts = append(groupCounts, currentCount)
	}

	if len(groupCounts) != len(damagedGroups) {
		return false
	}

	for i := 0; i < len(groupCounts); i++ {
		count := groupCounts[i]
		expected := damagedGroups[i]

		if count != expected {
			return false
		}
	}

	return true
}

func Part2(scanner *bufio.Scanner) (string, error) {
	records, err := parseRecords(scanner)
	if err != nil {
		return "", err
	}

	var sum int
	for _, r := range records {
		expanded := r.Expand(5)
		validArrangements := validArrangements(&expanded)
		sum += len(validArrangements)
	}
	result := strconv.Itoa(sum)
	return result, nil
}
