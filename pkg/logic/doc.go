package logic

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/iton0/hookup/pkg/git"
)

func OpenBrowser(key string) error {
	const Site = "https://git-scm.com/docs/githooks"
	var cmd *exec.Cmd
	hook, err := git.Hooks(key)

	if err != nil {
		return err
	}

	url := Site + hook

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
