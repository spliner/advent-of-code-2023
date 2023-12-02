package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spliner/aoc2023/pkg/day2"
)

func init() {
	rootCmd.AddCommand(day2Cmd)
}

var day2Cmd = &cobra.Command{
	Use:   "day2 [input path] [part]",
	Short: "Day 2 solution",
	RunE: func(_ *cobra.Command, args []string) error {
		input, err := readInput(args[0])
		if err != nil {
			return err
		}

		part := args[1]
		switch part {
		case "1":
			result, err := day2.Part1(input)
			if err != nil {
				return err
			}
			fmt.Printf("Day 2 part 1: %s\n", result)
		case "2":
			result, err := day2.Part2(input)
			if err != nil {
				return err
			}
			fmt.Printf("Day 2 part 2: %s\n", result)
		default:
			return fmt.Errorf("invalid part: %s", part)
		}

		return nil
	},
}
