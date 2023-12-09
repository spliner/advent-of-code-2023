package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spliner/aoc2023/pkg/day1"
	"github.com/spliner/aoc2023/pkg/day2"
	"github.com/spliner/aoc2023/pkg/day3"
	"github.com/spliner/aoc2023/pkg/day4"
	"github.com/spliner/aoc2023/pkg/day5"
	"github.com/spliner/aoc2023/pkg/day6"
	"github.com/spliner/aoc2023/pkg/day7"
	"github.com/spliner/aoc2023/pkg/day8"
	"github.com/spliner/aoc2023/pkg/day9"
)

type solver func(*bufio.Scanner) (string, error)

func init() {
	addCmd(1, day1.Part1, day1.Part2)
	addCmd(2, day2.Part1, day2.Part2)
	addCmd(3, day3.Part1, day3.Part2)
	addCmd(4, day4.Part1, day4.Part2)
	addCmd(5, day5.Part1, day5.Part2)
	addCmd(6, day6.Part1, day6.Part2)
	addCmd(7, day7.Part1, day7.Part2)
	addCmd(8, day8.Part1, day8.Part2)
	addCmd(9, day9.Part1, day9.Part2)
}

var rootCmd = &cobra.Command{
	Use:     "aoc2023 [day] [input path] [part]",
	Short:   "Solutions for Advent of Code 2023",
	Version: "1.0.0",
	Example: "aoc2023 day1 ../inputs/day1.txt 1",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func addCmd(day int, part1, part2 solver) {
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("day%d [input path] [part]", day),
		Short: fmt.Sprintf("Day %d solution", day),
		Args:  cobra.ExactArgs(2),
		RunE: func(_ *cobra.Command, args []string) error {
			readFile, err := os.Open(args[0])
			if err != nil {
				return err
			}
			defer readFile.Close()

			scanner := bufio.NewScanner(readFile)
			scanner.Split(bufio.ScanLines)

			part := args[1]
			switch part {
			case "1":
				if part1 == nil {
					return fmt.Errorf("part 1 not yet implemented for day %d", day)
				}

				result, err := part1(scanner)
				if err != nil {
					return err
				}
				fmt.Printf("Day %d part 1: %s\n", day, result)
			case "2":
				if part2 == nil {
					return fmt.Errorf("part 2 not yet implemented for day %d", day)
				}

				result, err := part2(scanner)
				if err != nil {
					return err
				}
				fmt.Printf("Day %d part 2: %s\n", day, result)
			default:
				return fmt.Errorf("invalid part: %s", part)
			}

			return nil
		},
	}

	rootCmd.AddCommand(cmd)
}
