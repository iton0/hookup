package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// https://github.com/spf13/cobra/blob/main/site/content/user_guide.md
var rootCmd = &cobra.Command{
	// TODO check docs about this
	Use:     "hu",
	Aliases: []string{"hookup"},
	Short:   "Hookup CLI",
	Long:    `Hookup is a tool that allows custom git hooks management.`,
	Version: "0.0.1",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := cmd.Help(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	// add commands here
	rootCmd.AddCommand(initCmd, docCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
