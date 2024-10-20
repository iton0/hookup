package cmd

import (
	"bytes"
	"fmt"
	"testing"
)

// Mock out the command execution
func TestRootCmd(t *testing.T) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)

	// Set up any flags or arguments as needed
	rootCmd.SetArgs([]string{"--version"})

	// Execute the command
	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	// Check the output
	got := buf.String()
	want := fmt.Sprintf("hu version %s\n", rootCmd.Version)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
