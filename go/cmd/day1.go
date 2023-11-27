package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spliner/aoc2023/pkg/day1"
)

func init() {
	rootCmd.AddCommand(day1Cmd)
}

var day1Cmd = &cobra.Command{
	Use:   "day1 [input path] [part]",
	Short: "Day 1 solution",
	RunE: func(_ *cobra.Command, args []string) error {
		input, err := readInput(args[0])
		if err != nil {
			return err
		}

		part := args[1]
		switch part {
		case "1":
			result, err := day1.Part1(input)
			if err != nil {
				return err
			}
			fmt.Printf("Day 1 part 1: %s\n", result)
		case "2":
			result, err := day1.Part2(input)
			if err != nil {
				return err
			}
			fmt.Printf("Day 1 part 2: %s\n", result)
		default:
			return fmt.Errorf("invalid part: %s", part)
		}

		return nil
	},
}
