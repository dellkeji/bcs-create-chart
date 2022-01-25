package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/dellkeji/bcs-create-chart/pkg/action"
)

var rootCmd = &cobra.Command{
	Use:   "bcs-create",
	Short: "bcs-create create chart content for bcs scenes",
	Long:  "bcs-create create chart content for bcs scenes",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// 获取chart名称及创建的目录
		chartName := filepath.Base(args[0])
		chartDir := filepath.Dir(chartName)
		// 创建模板
		path, err := action.Create(chartName, chartDir)
		if err != nil {
			fmt.Printf("create bcs chart error, %v", err)
			os.Exit(1)
		}
		fmt.Printf("create chart success: %s", path)
	},
}

// Execute execute unittest command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("execute bcs-create error, %v", err)
		os.Exit(1)
	}
}
