package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Starting scan...")

	// Scan the music directory
	files, err := scan("../music")
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(files)

	fmt.Println("Found files:", len(files))

	fmt.Println("Extracting metadata...")

	extractMetadata(files)

	// For each found file, extract artist name, album name, and song name

	// Index found music files to the database

	// Create a new artist if it doesn't exist

	// Add metadata to the database
}
