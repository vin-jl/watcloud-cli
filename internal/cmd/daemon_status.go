package cmd

import (
	"fmt"

	"github.com/WATonomous/watcloud-cli/internal/daemon"
	"github.com/spf13/cobra"
)

var daemonStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Lists all non-interactive background user processes (daemons).",
	Run: func(cmd *cobra.Command, args []string) {
		result, _ := daemon.ListDaemons()
		fmt.Print(result)
	},
}

func init() {
	daemonCmd.AddCommand(daemonStatusCmd)
}
