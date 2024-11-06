package browser

import (
	"fmt"
	"os/exec"
	"runtime"
)

func Open(url string) error {

	if url == "" {
		return fmt.Errorf("URL is empty")
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		return fmt.Errorf("unsupported platform %s", runtime.GOOS)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to open URL %w", err)
	}

	return nil
}
