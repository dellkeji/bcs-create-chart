package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/dellkeji/bcs-create-chart/pkg/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "version of bcs-create",
	Run: func(cmd *cobra.Command, args []string) {
		info := []string{
			"Version  :  " + version.Version,
			"GoVersion:  " + version.GoVersion,
		}
		fmt.Println(strings.Join(info, "\n"))
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
