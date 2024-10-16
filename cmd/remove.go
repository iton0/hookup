package cmd

import (
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Hookup.",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO implement logic
	},
}
