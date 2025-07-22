package quota

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/cpu"
)

// CPUUsage prints the current CPU usage percentage with color.
func CPUUsage() error {
	percent, err := cpu.Percent(0, true)
	if err != nil || len(percent) == 0 {
		return err
	}
	info, _ := cpu.Info()
	logical, _ := cpu.Counts(true)
	physical, _ := cpu.Counts(false)
	skyBlue := func(s string) string {
		return "\x1b[1m\x1b[38;2;16;128;255m" + s + "\x1b[0m"
	}
	faint := color.New(color.Faint).SprintFunc()
	fmt.Println(skyBlue("CPU Usage"))
	fmt.Println(faint(strings.Repeat("─", 40)))
	if len(info) > 0 {
		fmt.Printf("Model: %s\n", info[0].ModelName)
	}
	fmt.Printf("Cores: %d logical / %d physical\n", logical, physical)
	fmt.Println(faint(strings.Repeat("-", 40)))
	// Print per-core usage
	fmt.Printf("%-8s %-8s\n", "Core", "Usage %")
	fmt.Println(faint(strings.Repeat("-", 8) + " " + strings.Repeat("-", 8)))
	for i, p := range percent {
		var percentStr string
		switch {
		case p < 50:
			percentStr = color.New(color.FgGreen).Sprintf("%.0f%%", p)
		case p <= 70:
			percentStr = color.New(color.FgYellow).Sprintf("%.0f%%", p)
		default:
			percentStr = color.New(color.FgRed).Sprintf("%.0f%%", p)
		}
		fmt.Printf("%-8d %-8s\n", i, percentStr)
	}
	return nil
}
