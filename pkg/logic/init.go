package logic

import (
	"fmt"
	"os/exec"

	"github.com/iton0/hookup/pkg/util"
)

// INFO: main logic function for command will have to take all flags as arguments
func Init(userPath string) {
	// local repo folder name to hold git hooks
	const DefaultFolder = ".hookup"

	var path []string
	var fullPath string

	// Check for the flag arguments
	// User provided custom path; update path var
	if userPath != "" {
		path = util.SplitFilePath(userPath)
	} else {
		// Defaults to current working directory via relative path
		path = []string{"."}
	}

	// check if there is a hookup folder and create if necessary
	if !util.DoesDirectoryExist(path) {
		var err error
		fullPath, err = util.CreateFolder(path, DefaultFolder)
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
