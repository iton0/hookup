// INFO: use this command as the basic structure for other commands after finishing
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
		Run:   logic.Init,
	}
)

func init() {
	rootCmd.AddCommand(initCmd)
}
