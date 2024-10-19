package logic

import (
	"fmt"
	"os/exec"

	"github.com/iton0/hookup/pkg/util"
)

// local repo folder name to hold git hooks
const defaultFolder = ".hookup"

// Defaults to current working directory; how to change?
var path = []string{"."}

// INFO: main logic function for command will have to take all flags as arguments
func Init(userPath string) {
	var fullPath string

	// Check for the flag arguments
	// User provided custom path; update path var
	if userPath != "" {
		path = util.SplitFilePath(userPath)
	}

	// check if there is a hookup folder and create if necessary
	if !util.DoesDirectoryExist(path) {
		var err error
		fullPath, err = util.CreateFolder(path, defaultFolder)
		fmt.Println("Error:", err)
	}

	// update the local git core.hookPath
	cmd := exec.Command("git", "config", "--local", "core.hooksPath", fullPath)

	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error updating hooksPath:", err)
		return
	}
}
