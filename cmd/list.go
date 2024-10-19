package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Hookup.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO implement logic
	},
}
