package cmd

import (
	"github.com/taylormonacelli/itmetrics/run"

	"github.com/spf13/cobra"
)

var (
	outdir       string
	manifestPath string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		run.Run(manifestPath, outdir)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&manifestPath, "manifest", "m", "https://raw.githubusercontent.com/taylormonacelli/wildcalifornia/master/all.k", "Path or URL to manifest.k")
	runCmd.Flags().StringVarP(&outdir, "outdir", "o", ".", "Output directory")
}
