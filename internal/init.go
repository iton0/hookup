// we can set a relative path to the local .hookup folder
// ex. git config --local core.hooksPath ./.hookup

// for worktree specify which worktree to look and update the hookPath to
// something like 'hu init --worktree [worktree repo]'
// they can leave blank in which the hu will look at all the available worktrees to find a .hookup folder

package internal

import (
	"fmt"
	"os/exec"

	"github.com/iton0/hookup/util"
)

// local repo folder name to hold git hooks
const defaultFolder = ".hookup"

// Defaults to current working directory; how to change?
var path = []string{"."}

func Init() {
	var fullpath string

	// TODO need to check if user provided path and update path variable
	// This path with include the / character so how do i parse that if even
	//
	// Possible design choice: provide a convention in the README for where to place the defaultFolder. Hookup will only look in those

	// check if there is a hookup folder and create if necessary
	if !util.DoesDirectoryExist(path) {
		fullpath, err = util.CreateFolder(path, defaultFolder)
	}

	// update the local git core.hookPath
	cmd := exec.Command("git", "config", "--local", "core.hooksPath", fullpath)

	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
