package cmd

import (
	// "github.com/iton0/hookup/pkg/logic"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var (
	listCmd = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "list git hooks.",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	// flags
)
