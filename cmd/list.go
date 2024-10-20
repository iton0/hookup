package cmd

import (
	// "github.com/iton0/hookup/pkg/logic"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list git hooks.",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// flags
)

func init() {
	rootCmd.AddCommand(listCmd)
}
