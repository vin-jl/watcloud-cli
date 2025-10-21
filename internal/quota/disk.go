package quota

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/disk"
)

// DiskUsage prints disk usage statistics for the root filesystem with color.
func DiskUsage() error {
	usage, err := disk.Usage("/")
	if err != nil {
		return err
	}
	total := float64(usage.Total) / (1 << 30)
	used := float64(usage.Used) / (1 << 30)
	free := float64(usage.Free) / (1 << 30)
	percent := usage.UsedPercent

	skyBlue := func(s string) string {
		return "\x1b[1m\x1b[38;2;16;128;255m" + s + "\x1b[0m"
	}
	faint := color.New(color.Faint).SprintFunc()
	fmt.Println(skyBlue("Disk Usage"))
	fmt.Println(faint(strings.Repeat("─", 40)))
	fmt.Printf("%-12s %-12s %-12s %-12s\n", "Total", "Used", "Free", "Used %")
	fmt.Println(faint(strings.Repeat("-", 12) + " " + strings.Repeat("-", 12) + " " + strings.Repeat("-", 12) + " " + strings.Repeat("-", 12)))
	var percentStr string
	switch {
	case percent <= 70:
		percentStr = color.New(color.FgGreen).Sprintf("%.0f%%", percent)
	case percent >= 90:
		percentStr = color.New(color.FgRed).Sprintf("%.0f%%", percent)
	default:
		percentStr = color.New(color.FgYellow).Sprintf("%.0f%%", percent)
	}
	fmt.Printf("%-12s %-12s %-12s %-12s\n",
		fmt.Sprintf("%.2f GiB", total),
		fmt.Sprintf("%.2f GiB", used),
		fmt.Sprintf("%.2f GiB", free),
		percentStr)
	fmt.Println()
	return nil
}
