package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("Starting scan...")
	root := filepath.Join("..", "music")
	fmt.Println(root)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".mp3") {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path %v: %v\n", root, err)
		return
	}
}
