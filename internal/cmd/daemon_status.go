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
		result := daemon.ListDaemons()
		fmt.Println(result)
	},
}

func init() {
	daemonCmd.AddCommand(daemonStatusCmd)
}
