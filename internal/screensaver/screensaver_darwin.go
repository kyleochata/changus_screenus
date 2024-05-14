package main

import (
	"fmt"
	"os/exec"
)

func SetScreensaverMacOs(imagePath string) error {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`tell application "System Events" to set picture of current desktop to POSIX file "%s"`, imagePath))
	return cmd.Run()
}
