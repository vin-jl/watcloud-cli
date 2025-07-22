package cmd

import (
	"github.com/spf13/cobra"
)

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Manage and inspect user daemons",
}

func init() {
	rootCmd.AddCommand(daemonCmd)
}
