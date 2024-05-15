package utils

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func GetFiles(folderPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func GetRandomFile(files []string) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return files[r.Intn(len(files))]
}
