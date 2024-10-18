package logic

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/iton0/hookup/pkg/git"
)

const site = "https://git-scm.com/docs/githooks"

func OpenBrowser(key string) error {
	var cmd *exec.Cmd
	hook, err := git.Hooks(key)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	url := site + hook

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	default:
		return fmt.Errorf("unsupported platform")
	}

	return cmd.Start()
}
