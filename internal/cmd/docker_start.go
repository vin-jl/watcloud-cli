package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var dockerStartCmd = &cobra.Command{
	Use:     "start [disk_size_MiB]",
	Aliases: []string{"run"},
	Short:   "Starts the rootless Docker Daemon.",
	Long:    "Starts the rootless Docker Daemon using slurm-start-dockerd.sh. Optionally specify disk size in MiB.",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cmdArgs := []string{}

		// If disk size argument is provided, add the --gres flag
		if len(args) > 0 {
			cmdArgs = append(cmdArgs, "--gres", fmt.Sprintf("tmpdisk:%s", args[0]))
		}

		// Run slurm-start-dockerd.sh
		execCmd := exec.Command("slurm-start-dockerd.sh", cmdArgs...)
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		execCmd.Stdin = os.Stdin

		if err := execCmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error starting Docker daemon: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	dockerCmd.AddCommand(dockerStartCmd)
}
