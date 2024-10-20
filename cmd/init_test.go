package cmd

import (
	"bytes"
	"testing"
)

// Mock out the command execution
func TestInitCmd(t *testing.T) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)

	// Set up any flags or arguments as needed
	rootCmd.SetArgs([]string{"init"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}
}
