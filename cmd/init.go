package cmd

import (
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Initialize Hookup.",
		Run: func(cmd *cobra.Command, args []string) {
			// TODO implement logic
		},
	}

	// flags
	Worktree string
	Path     string
)

func init() {
	initCmd.Flags().StringVarP(&Worktree, "worktree", "wt", "", "The current working directory is a worktree")
	initCmd.Flags().StringVarP(&Path, "path", "p", "", "Specify path to initialize hookup folder")
	rootCmd.AddCommand(initCmd)
}
