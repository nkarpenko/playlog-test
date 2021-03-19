package cmd

import (
	"fmt"

	"github.com/nkarpenko/playlog-test/api/app"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display app version.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("App v" + app.Version)
		return
	},
}
