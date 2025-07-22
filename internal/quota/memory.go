package quota

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/mem"
)

// MemoryUsage prints memory usage statistics with color.
func MemoryUsage() error {
	vm, err := mem.VirtualMemory()
	if err != nil {
		return err
	}
	total := float64(vm.Total) / (1 << 30)
	used := float64(vm.Used) / (1 << 30)
	free := float64(vm.Available) / (1 << 30)
	percent := vm.UsedPercent

	skyBlue := func(s string) string {
		return "\x1b[1m\x1b[38;2;16;128;255m" + s + "\x1b[0m"
	}
	faint := color.New(color.Faint).SprintFunc()
	fmt.Println(skyBlue("Memory Usage"))
	fmt.Println(faint(strings.Repeat("─", 40)))
	fmt.Printf("%-12s %-12s %-12s %-12s\n", "Total", "Used", "Free", "Used %")
	fmt.Println(faint(strings.Repeat("-", 12) + " " + strings.Repeat("-", 12) + " " + strings.Repeat("-", 12) + " " + strings.Repeat("-", 12)))
	var percentStr string
	switch {
	case percent <= 60:
		percentStr = color.New(color.FgGreen).Sprintf("%.0f%%", percent)
	case percent >= 80:
		percentStr = color.New(color.FgRed).Sprintf("%.0f%%", percent)
	default:
		percentStr = color.New(color.FgYellow).Sprintf("%.0f%%", percent)
	}
	fmt.Printf("%-12.2f %-12.2f %-12.2f %-12s\n", total, used, free, percentStr)
	return nil
}
