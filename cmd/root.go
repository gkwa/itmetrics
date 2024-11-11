package cmd

import (
	"fmt"
	"os"

	"github.com/carlmjohnson/versioninfo/v2"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "itmetrics",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,

	Version: fmt.Sprintf("%s, commit %s", versioninfo.Short(), versioninfo.Revision),
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
