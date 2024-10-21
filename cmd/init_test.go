package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/iton0/hookup/internal/util"
)

func TestInitCmd(t *testing.T) {
	var buf bytes.Buffer
	rootCmd.SetOut(&buf)
	testCreatedDir := "../cmd/.hookup"

	tests := []struct {
		args []string
		want string
	}{
		{
			args: []string{"init"},
			// INFO: this is the last line of expected output
			want: "hooksPath successfully set to ./.hookup",
		},
		// Add more test cases here if necessary, e.g., for error conditions
	}

	for _, tt := range tests {
		buf.Reset() // Reset the buffer before each command execution
		rootCmd.SetArgs(tt.args)

		// Execute the command
		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("Command failed for args %v: %v", tt.args, err)
		}

		// Check the output
		got := buf.String()
		lines := strings.Split(got, "\n")
		if len(lines) > 0 {
			gotLastLine := lines[len(lines)-2] // Get the last non-empty line
			if gotLastLine != tt.want {
				t.Errorf("got last line %q, want %q for args %v", gotLastLine, tt.want, tt.args)
			}
		} else {
			t.Errorf("got empty output for args %v", tt.args)
		}

		// Clean up if needed, e.g., removing directory created during the test
		if util.DoesDirectoryExist(testCreatedDir) {
			os.Remove(testCreatedDir) // Ensure the cleanup is performed after each test case
		}
	}
}
