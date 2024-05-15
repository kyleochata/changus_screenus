//go:build windows

package pictureBackground

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func ChangeWindows(imagePath string) error {
	absImagePath, err := filepath.Abs(imagePath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}

	// Prepare PowerShell command
	psCommand := fmt.Sprintf(`$definition = @"
using System;
using System.Runtime.InteropServices;

public class ScreenSaver {
    [DllImport("user32.dll", CharSet = CharSet.Auto)]
    public static extern void SystemParametersInfo(uint uiAction, uint uiParam, string pvParam, uint fWinIni);
}
"@
Add-Type -TypeDefinition $definition
[ScreenSaver]::SystemParametersInfo(0x0014, 0, "%s", 0x0001)`, absImagePath)

	// Log the PowerShell command for debugging
	log.Printf("Running PowerShell command: %s\n", psCommand)

	// Execute the PowerShell command
	cmd := exec.Command("powershell", "-Command", psCommand)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("PowerShell command failed: %v\nOutput: %s\n", err, string(output))
		return fmt.Errorf("PowerShell command failed: %v\nOutput: %s", err, string(output))
	}

	log.Printf("PowerShell command output: %s\n", string(output))
	return nil
}
