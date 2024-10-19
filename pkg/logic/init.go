package logic

import (
	"fmt"
	"os/exec"

	"github.com/iton0/hookup/cmd"
	"github.com/iton0/hookup/pkg/util"
)

// local repo folder name to hold git hooks
const defaultFolder = ".hookup"

// Defaults to current working directory; how to change?
var path = []string{"."}

func Init() {
	var fullPath string

	// Check for the flag arguements

	// TODO: how do i properly do the worktree flag because setup may differ.
	// The portable way is the use relative path but i need to make sure that I properly use relative
	if cmd.Worktree != "" {
		// Check if it is a worktree via does .git directory exist in cwd
	}
	// User provided custom path; update path var
	if cmd.Path != "" {
		path = util.SplitFilePath(cmd.Path)
	}

	// Possible design choice: provide a convention in the README for where to place the defaultFolder. Hookup will only look in those

	// check if there is a hookup folder and create if necessary
	if !util.DoesDirectoryExist(path) {
		var err error // is there a better way to do this?
		fullPath, err = util.CreateFolder(path, defaultFolder)
		fmt.Println("Error:", err)
	}

	// update the local git core.hookPath
	cmd := exec.Command("git", "config", "--local", "core.hooksPath", fullPath)

	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
