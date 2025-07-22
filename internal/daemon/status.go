package daemon

import (
	"os"
	"strings"

	"github.com/fatih/color"
)

// ListDaemons prints the status of Docker rootless mode (docker.sock) with colorized output.
func ListDaemons() string {
	dockerSocket := "/tmp/run/docker.sock"
	skyBlue := func(s string) string {
		return "\x1b[1m\x1b[38;2;16;128;255m" + s + "\x1b[0m"
	}
	faint := color.New(color.Faint).SprintFunc()
	status := skyBlue("Daemon Status") + "\n" + faint(strings.Repeat("─", 40)) + "\n"
	if _, err := os.Stat(dockerSocket); err == nil {
		status += color.New(color.FgGreen).Sprint("Docker rootless: ") + "Found " + faint("(docker.sock present)") + "\n"
	} else {
		status += color.New(color.FgRed).Sprint("Docker rootless: ") + "Not found " + faint("(docker.sock missing)") + "\n"
	}
	return status
}
