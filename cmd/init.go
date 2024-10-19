package cmd

import (
	"github.com/iton0/hookup/pkg/logic"
	"github.com/spf13/cobra"
)

var (
	initCmd = &cobra.Command{
		// TODO check docs about argument restrictions and all dat
		Use:   "init",
		Short: "Initialize Hookup.",
		Run: func(cmd *cobra.Command, args []string) {
			logic.Init(Path)

		},
	}

	// flags
	Path string
)

func init() {
	initCmd.Flags().StringVarP(&Path, "path", "p", "", "Specify path to initialize hookup folder")
	rootCmd.AddCommand(initCmd)
}
