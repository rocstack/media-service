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
		extractMetadataFile(file)
		// fmt.Printf("Artist: %s, Album: %s, Song: %s\n", artist, album, song)
	}
}

// extractMetadata extracts the artist name, album name, and song name from the given file path.
func extractMetadataFile(file string) (string, string, string) {
	// Get the file name without extension
	name := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file))

	// Split the file path into parts
	parts := strings.Split(filepath.Dir(file), string(filepath.Separator))

	// Ignore the first two parts (../music)
	if len(parts) >= 3 && parts[0] == ".." && parts[1] == "music" {
		parts = parts[2:]
	}

	fmt.Printf(parts[0] + "\n")

	// Extract artist, album, and song names
	var artist, album, song string
	switch len(parts) {
	case 1:
		// Song file is in the root of the directory
		artist = "Unknown Artist"
		album = "Unknown Album"
		song = name
	case 2:
		// Song file is under the artist directory
		artist = parts[0]
		album = "Unknown Album"
		song = name
	case 3:
		// Song file is under the album directory
		artist = parts[0]
		album = parts[1]
		song = name
	default:
		// Invalid file path
		artist = "Unknown Artist"
		album = "Unknown Album"
		song = name
	}
	return artist, album, song
}
