package cmd

import (
	"github.com/spf13/cobra"
)

var quotaCmd = &cobra.Command{
	Use:   "quota",
	Short: "Show quota usage (disk, memory, CPU)",
}

func init() {
	rootCmd.AddCommand(quotaCmd)
}
