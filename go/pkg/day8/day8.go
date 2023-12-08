package day8

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spliner/aoc2023/pkg/mathutils"
)

var emptyLineErr = errors.New("empty line")

type Direction rune

type Node struct {
	Value string
	Left  string
	Right string
}

type Map struct {
	Directions []Direction
	NodeMap    map[string]Node
}

func Part1(scanner *bufio.Scanner) (string, error) {
	m, err := parseMap(scanner)
	if err != nil {
		return "", err
	}

	node := m.NodeMap["AAA"]
	steps := walk(m, node, func(n Node) bool { return n.Value != "ZZZ" })

	result := strconv.Itoa(steps)
	return result, nil
}

func parseMap(scanner *bufio.Scanner) (*Map, error) {
	if !scanner.Scan() {
		return nil, errors.New("could not read first line")
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	directions := []Direction(strings.TrimSpace(scanner.Text()))

	nodes := make([]Node, 0)
	nodeMap := make(map[string]Node)

	scanner.Scan() // Empty line

	if !scanner.Scan() {
		return nil, errors.New("could not root node line")
	}

	rootNode, err := readLine(scanner)
	if err != nil {
		return nil, err
	}

	nodes = append(nodes, rootNode)
	nodeMap[rootNode.Value] = rootNode

	for scanner.Scan() {
		node, err := readLine(scanner)
		if err != nil {
			if errors.Is(err, emptyLineErr) {
				continue
			}

			return nil, err
		}

		nodes = append(nodes, node)
		nodeMap[node.Value] = node
	}

	m := Map{
		Directions: directions,
		NodeMap:    nodeMap,
	}

	return &m, nil
}

func readLine(scanner *bufio.Scanner) (Node, error) {
	if err := scanner.Err(); err != nil {
		return Node{}, err
	}

	line := strings.TrimSpace(scanner.Text())
	if line == "" {
		return Node{}, emptyLineErr
	}

	nodeValue, rawDestinations, found := strings.Cut(line, " = ")
	if !found {
		return Node{}, fmt.Errorf("invalid line: %s", line)
	}

	nodeValue = strings.TrimSpace(nodeValue)

	left, right, found := strings.Cut(rawDestinations, ", ")
	if !found {
		return Node{}, fmt.Errorf("invalid line: %s", line)
	}

	left = left[1:]
	right = right[:len(right)-1]

	node := Node{
		Value: nodeValue,
		Left:  left,
		Right: right,
	}

	return node, nil
}

func walk(m *Map, node Node, condition func(n Node) bool) int {
	var steps, directionIndex int
	for condition(node) {
		steps++
		direction := m.Directions[directionIndex]
		if direction == 'L' {
			node = m.NodeMap[node.Left]
		} else {
			node = m.NodeMap[node.Right]
		}

		directionIndex++
		if directionIndex >= len(m.Directions) {
			directionIndex = 0
		}
	}

	return steps
}

func Part2(scanner *bufio.Scanner) (string, error) {
	m, err := parseMap(scanner)
	if err != nil {
		return "", err
	}

	nodes := make([]Node, 0)
	for _, node := range m.NodeMap {
		if strings.HasSuffix(node.Value, "A") {
			nodes = append(nodes, node)
		}
	}

	nodeSteps := make([]int, len(nodes))
	for i, node := range nodes {
		steps := walk(m, node, func(n Node) bool { return !strings.HasSuffix(n.Value, "Z") })
		nodeSteps[i] = steps
	}

	lcm := mathutils.LeastCommonMultiple(nodeSteps)

	result := strconv.Itoa(lcm)
	return result, nil
}
