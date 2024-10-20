package logic

import (
	"os/exec"

	"github.com/iton0/hookup/pkg/util"
	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) {
	cmd.Println("Initializing hookup...")

	// local repo folder name to hold git hooks via relative path
	const FullPath = "./.hookup"

	// check if cwd is git repo
	if err := exec.Command("git", "-C", ".", "rev-parse", "--is-inside-work-tree").Run(); err != nil {
		cmd.Println("Fatal: Must be run in a git repository")
		return
	}

	// check if there is a hookup folder and create if necessary
	if !util.DoesDirectoryExist(FullPath) {
		cmd.Println("No .hookup folder found.")
		cmd.Println("Creating .hookup folder now...")
		if err := util.CreateFolder(FullPath); err != nil {
			cmd.Println("Error creating folder: ", err)
			return
		}
		cmd.Println("Folder created successfully!")
	}

	// update the local git core.hookPath
	cmd.Println("Updating local hooksPath...")
	if err := exec.Command("git", "config", "--local", "core.hooksPath", FullPath).Run(); err != nil {
		cmd.Println("Error updating hooksPath: ", err)
		return
	}
	cmd.Println("hooksPath successfully set to " + FullPath)
}
