package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update hookup config.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO implement logic
	},
}
