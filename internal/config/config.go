package config

import (
	"flag"
)

type Config struct {
	FolderPath string
}

func LoadConfig() Config {
	folderPath := flag.String("folder", ".", "Path to the folder with images")
	flag.Parse()
	return Config{
		FolderPath: *folderPath,
	}
}
