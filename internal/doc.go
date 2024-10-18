package internal

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
)

func gitHook(key string) (string, error) {
	value, exists := gitHooks[key]
	if !exists {
		return "", errors.New("hook not found: " + key)
	}
	return value, nil
}

func OpenBrowser(key string) error {
	var cmd *exec.Cmd
	hook, err := gitHook(key)

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
