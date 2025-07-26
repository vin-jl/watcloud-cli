package daemon

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
)

// probe represents one daemon we test.
type probe struct {
	name   string
	socket func() string // returns discovered socket or ""
}

// only Docker and Slurm remain.
var allProbes = []probe{
	{
		name: "Docker",
		socket: func() string {
			if p := os.Getenv("DOCKER_HOST"); strings.HasPrefix(p, "unix://") {
				return strings.TrimPrefix(p, "unix://")
			}
			paths := []string{
				os.Getenv("XDG_RUNTIME_DIR") + "/docker.sock", // rootless
				"/var/run/docker.sock",                        // root-ful
			}
			for _, p := range paths {
				if p == "" {
					continue
				}
				if _, err := os.Stat(p); err == nil {
					return p
				}
			}
			return ""
		},
	},
	{
		name: "Slurm",
		socket: func() string {
			paths := []string{
				"/var/run/slurmd.pid",
				"/run/slurmd.pid",
			}
			for _, p := range paths {
				if _, err := os.Stat(p); err == nil {
					return p
				}
			}
			return ""
		},
	},
}

// ListDaemons returns (formattedReport, exitCode).
func ListDaemons() (string, int) {
	var b strings.Builder
	header := color.New(color.Bold, color.FgHiCyan).Sprint("Daemon Status")
	divider := color.New(color.Faint).Sprint(strings.Repeat("─", 40))
	fmt.Fprintln(&b, header)
	fmt.Fprintln(&b, divider)

	w := tabwriter.NewWriter(&b, 0, 0, 2, ' ', 0)
	ok := color.New(color.FgGreen).SprintFunc()
	bad := color.New(color.FgRed).SprintFunc()

	exitCode := 0
	for _, p := range allProbes {
		if sock := p.socket(); sock != "" {
			fmt.Fprintf(w, "%s:\t%s\t(%s)\n", p.name, ok("running"), sock)
		} else {
			fmt.Fprintf(w, "%s:\t%s\t(not detected)\n", p.name, bad("stopped"))
			exitCode = 1
		}
	}
	w.Flush()
	return b.String(), exitCode
}
