package status

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
)

// GetClusterStatus prints cluster node status and maintenance info.
func GetClusterStatus() {
	skyBlue := func(s string) string {
		// ANSI 24-bit color: \x1b[38;2;16;128;255m ... \x1b[0m
		return "\x1b[1m\x1b[38;2;16;128;255m" + s + "\x1b[0m"
	}
	sep := color.New(color.Faint).Sprint(strings.Repeat("─", 60))
	fmt.Println(skyBlue("WATcloud Infrastructure Status"))
	fmt.Println(sep)
	fmt.Println(skyBlue("Cluster Nodes"))
	printNodeStatus()
	fmt.Println(sep)
	fmt.Println(skyBlue("Maintenance Nodes"))
	printMaintenanceNodes()
}

func printNodeStatus() {
	type Check struct {
		Name     string `json:"name"`
		Status   string `json:"status"`
		LastPing string `json:"last_ping"`
		Desc     string `json:"desc"`
		Tags     string `json:"tags"`
	}
	type ChecksResp struct {
		Checks []Check `json:"checks"`
	}

	req, err := http.NewRequest("GET", "https://healthchecks.io/api/v3/checks/", nil)
	if err != nil {
		color.New(color.FgRed).Printf("Error creating request: %v\n", err)
		return
	}
	req.Header.Set("X-Api-Key", "tCsst0GSKpfvslmpmlsmivRrUCRuv6Iv")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		color.New(color.FgRed).Printf("Error fetching cluster status: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var checks ChecksResp
	if err := json.NewDecoder(resp.Body).Decode(&checks); err != nil {
		color.New(color.FgRed).Printf("Error decoding status JSON: %v\n", err)
		return
	}

	if len(checks.Checks) == 0 {
		color.New(color.FgYellow).Println("(No cluster status data found from API.)")
		return
	}

	// nameCol := color.New(color.Bold)
	statusUp := color.New(color.FgGreen)
	statusDown := color.New(color.FgRed)
	statusUnknown := color.New(color.FgYellow)

	// Print headers
	fmt.Printf("%-35s %-10s %-15s %s\n", "Node", "Status", "Last Ping", "Description")
	fmt.Printf("%-35s %-10s %-15s %s\n",
		strings.Repeat("─", 35),
		strings.Repeat("─", 10),
		strings.Repeat("─", 15),
		strings.Repeat("─", 20))

	// Only show these nodes
	allowedNodes := map[string]struct{}{
		"delta-slurm1-slurm-schedulable": {},
		"elastic-ssh":                    {},
		"thor-slurm1-slurm-schedulable":  {},
		"tr-slurm2-slurm-schedulable":    {},
		"trpro-slurm1-slurm-schedulable": {},
		"trpro-slurm2-slurm-schedulable": {},
		"wato-login1-ssh":                {},
		"wato-login2-ssh":                {},
		"wato2-slurm1-slurm-schedulable": {},
	}

	for _, check := range checks.Checks {
		if _, ok := allowedNodes[check.Name]; !ok {
			continue
		}
		ago := "?"
		if t, err := time.Parse(time.RFC3339Nano, check.LastPing); err == nil {
			dur := time.Since(t)
			switch {
			case dur < time.Minute:
				ago = fmt.Sprintf("%ds ago", int(dur.Seconds()))
			case dur < time.Hour:
				ago = fmt.Sprintf("%dm ago", int(dur.Minutes()))
			case dur < 24*time.Hour:
				ago = fmt.Sprintf("%dh ago", int(dur.Hours()))
			default:
				ago = fmt.Sprintf("%dd ago", int(dur.Hours()/24))
			}
		}

		desc := check.Desc
		if len(desc) > 40 {
			desc = desc[:37] + "..."
		}

		// Format the line first, then apply colors
		statusText := strings.ToUpper(check.Status)
		line := fmt.Sprintf("%-35s %-10s %-15s %s", check.Name, statusText, ago, desc)

		// Apply colors to specific parts
		switch check.Status {
		case "up":
			line = strings.Replace(line, statusText, statusUp.Sprint(statusText), 1)
		case "down":
			line = strings.Replace(line, statusText, statusDown.Sprint(statusText), 1)
		default:
			line = strings.Replace(line, statusText, statusUnknown.Sprint(statusText), 1)
		}

		fmt.Println(line)
	}
}

func printMaintenanceNodes() {
	faint := color.New(color.Faint).SprintFunc()
	cmd := exec.Command("scontrol", "show", "nodes")
	out, err := cmd.Output()
	if err != nil {
		if strings.Contains(err.Error(), "executable file not found") {
			color.New(color.FgYellow).Println("(scontrol not found: Maintenance node info unavailable. Install Slurm to enable this feature.)")
		} else {
			color.New(color.FgRed).Printf("Error running scontrol: %v\n", err)
		}
		return
	}
	lines := strings.Split(string(out), "\n")
	var nodeName, nodeState string
	printed := false
	fmt.Printf("%-28s  %s\n", faint("Node"), faint("State"))
	fmt.Println(faint(strings.Repeat("-", 40)))
	for _, line := range lines {
		if strings.Contains(line, "NodeName") {
			parts := strings.Split(line, " ")
			for _, part := range parts {
				if strings.HasPrefix(part, "NodeName=") {
					nodeName = strings.TrimPrefix(part, "NodeName=")
				}
				if strings.HasPrefix(part, "State=") {
					nodeState = strings.TrimPrefix(part, "State=")
				}
			}
		}
		if nodeName != "" && nodeState != "" {
			fmt.Printf("%-28s  %s\n", nodeName, nodeState)
			nodeName, nodeState = "", ""
			printed = true
		}
	}
	if !printed {
		fmt.Println(faint("(No maintenance nodes found.)"))
	}
}
