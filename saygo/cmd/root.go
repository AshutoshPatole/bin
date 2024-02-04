package cmd

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/TwiN/go-color"
	"github.com/common-nighthawk/go-figure"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

var platform, architecture string

var rootCmd = &cobra.Command{
	Use:   "saygo",
	Short: "Initator cli for linux machines",
	Run: func(cmd *cobra.Command, args []string) {
		platform = runtime.GOOS
		architecture = runtime.GOARCH

		myFigure := figure.NewColorFigure("SayGo", "", "green", true)
		myFigure.Print()

		totalCPU, _ := cpu.Counts(true)
		cpuUtilised, _ := cpu.Percent(time.Millisecond*100, false)

		memory, _ := mem.VirtualMemory()

		fmt.Printf("%-20s %s\n", "Operating System:", color.InBlue(platform))
		fmt.Printf("%-20s %s\n", "OS Architecture:", color.InBlue(architecture))
		fmt.Printf("%-20s %s\n", "CPU(s):", color.InBlue(totalCPU))
		fmt.Printf("%-20s %s\n", "CPU %:", color.InBlue(cpuUtilised[0]))
		fmt.Printf("%-20s %s\n", "Total Memory (GB):", color.InBlue(memory.Total/(1024*1024*1024)))
		fmt.Printf("%-20s %s\n", "Memory Used %:", color.InBlue(memory.UsedPercent))

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
