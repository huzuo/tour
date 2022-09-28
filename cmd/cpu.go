package cmd

import (
	"fmt"
	"time"
	"tour/internal/cpuer"

	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "打印cpu使用率",
	Long:  "打印cpu使用率",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			cpu_persent := cpuer.GetCpuInfo()
			fmt.Printf("\rCPU: %.2f%%", cpu_persent[0])
			time.Sleep(time.Second)
		}

	},
}
