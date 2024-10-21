package cmd

import (
	"github.com/iton0/hookup/internal/logic"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add git hook.",
	Run:   logic.Add,
}
