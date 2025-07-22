package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "watcloud",
	Short: "WATcloud CLI: Inspect resource usage and daemon status",
	Long:  `A CLI for inspecting user-specific resource usage and daemon status on WATcloud systems.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
