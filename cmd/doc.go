package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(docCmd)
}

var docCmd = &cobra.Command{
	Use:   "doc",
	Short: "View git hook documentation.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
