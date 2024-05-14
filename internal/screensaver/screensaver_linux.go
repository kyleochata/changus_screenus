package screensaver

import "os/exec"

func SetScreensaverLinux(imagePath string) error {
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.screensaver", "picture-uri", "file://"+imagePath)
	return cmd.Run()
}
