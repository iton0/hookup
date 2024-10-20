package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remove hook.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO implement logic
	},
}
