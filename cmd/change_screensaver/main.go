package main

import (
	"changus_screenus/internal/config"
	"changus_screenus/internal/pictureBackground"
	"changus_screenus/internal/utils"
	"fmt"
	"log"
	"runtime"
)

func main() {
	config := config.LoadConfig()
	files, err := utils.GetFiles(config.FolderPath)
	if err != nil {
		log.Fatal(err)
	}

	if len(files) == 0 {
		log.Fatal("No images in the folder")
	}

	randomFile := utils.GetRandomFile(files)
	fmt.Println("Selected file:", randomFile)
	switch os := runtime.GOOS; os {
	case "windows":
		err = pictureBackground.ChangeWindows(randomFile)
		// case "darwin":
		// 	err = pictureBackground.ChangeMac(randomFile)
		// case "linux":
		// 	err = pictureBackground.ChangeLinux(randomFile)
	default:
		log.Fatalf("Unsupported OS: %s", os)
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Changus_Screenus!!!")
}
