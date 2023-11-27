package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "aoc2023 [day]",
	Short:   "Solutions for Advent of Code 2023",
	Version: "1.0.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func readInput(path string) (string, error) {
	input, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(input), nil
}
