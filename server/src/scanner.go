package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func scan(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("Music directory not found.")
			return err
		}
		if !info.IsDir() && isMusicFile(path) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

// isMusicFile returns true if the given file path has a music file extension.
func isMusicFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	return ext == ".mp3" || ext == ".flac" || ext == ".wav"
}
