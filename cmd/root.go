package cmd

import (
	"fmt"
	"os"

	"github.com/iton0/hookup/pkg/logic"
	"github.com/spf13/cobra"
)

// https://github.com/spf13/cobra/blob/main/site/content/user_guide.md
var (
	rootCmd = &cobra.Command{
		// TODO check docs about this
		Use:     "hu",
		Aliases: []string{"hookup"},
		Short:   "Hookup CLI",
		Long:    `Hookup is a management tool for git hooks.`,
		Version: "0.0.1",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Help(); err != nil {
				return err
			}

			// TODO: create logic for flags

			return nil
		},
	}

	// flags
	Doc string
)

func init() {
	rootCmd.Flags().StringVar(&Doc, "doc", "", "Documentation for specified git hook")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
