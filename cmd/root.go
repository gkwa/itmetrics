package cmd

import (
	"fmt"

	"github.com/carlmjohnson/versioninfo"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "itmetrics",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,

	Version: fmt.Sprintf("%s, commit %s", versioninfo.Short(), versioninfo.Revision),
}

func Execute() error {
	return rootCmd.Execute()
}
