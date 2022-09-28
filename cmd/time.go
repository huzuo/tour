package cmd

import (
	"log"
	"strconv"
	"strings"
	"time"
	"tour/internal/timer"

	"github.com/spf13/cobra"
)

var caclulateTime string
var duration string

var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间格式处理",
	Long:  "时间格式处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s, %d", nowTime.Format("2006-01-02 15:04:05"), nowTime.Unix())
	},
}

var caclulateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTimer time.Time
		var layout = "2006-01-02 15:04:05"
		if caclulateTime == "" {
			currentTimer = timer.GetNowTime()
		} else {
			var err error
			space := strings.Count(caclulateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			if space == 1 {
				layout = "2006-01-02 15:04"
			}
			currentTimer, err = time.Parse(layout, caclulateTime)
			if err != nil {
				t, _ := strconv.Atoi(caclulateTime)
				currentTimer = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCaclulateTime(currentTimer, duration)
		if err != nil {
			log.Printf("timer.GetCaclulateTime err: %v", err)
		}
		log.Printf("输出结果：%s, %d", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(caclulateTimeCmd)

	caclulateTimeCmd.Flags().StringVarP(&caclulateTime, "calculate", "c", "", "需要计算的时间，有效单位为时间戳或已格式化后的时间")
	caclulateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "", `持续时间，有效单位为："ns","us","ms","s","m","h"`)
}
