package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bcs-create",
	Short: "bcs-create create chart content for bcs scenes",
	Long:  "bcs-create create chart content for bcs scenes",
}

// Execute execute unittest command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("execute bcs-create error, %v", err)
		os.Exit(1)
	}
}
