package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func extractMetadata(files []string) {
	// Example list of files
	// files := []string{
	//     "../music/Artist Name/Album Name/SongName.mp3",
	//     "../music/Artist Name/SongName2.mp3",
	//     "../music/Artist Name/Album Name 2/SongName3.mp3",
	//     "../music/Artist Name 2/Album Name/SongName4.mp3",
	//     "../music/SongName5.mp3",
	// }

	// Extract metadata from each file
	for _, file := range files {
		// artist, album, song := extractMetadataFile(file)
		extractFile(file)
		// fmt.Printf("Artist: %s, Album: %s, Song: %s\n", artist, album, song)
	}
}

// extracts the file paths
func extractFile(file string) (string, []string) {
	// Get the file name without extension
	fileName := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))

	// Split the file path into parts
	parts := strings.Split(filepath.Dir(file), string(filepath.Separator))

	// Ignore the first two parts (../music)
	if len(parts) >= 3 && parts[0] == ".." && parts[1] == "music" {
		parts = parts[2:]
	}

	fmt.Print(parts, fileName)

	return fileName, parts
}
