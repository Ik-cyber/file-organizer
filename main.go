package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Map file extensions to target folders
var extensionMap = map[string]string{
	".jpg":  "Images",
	".jpeg": "Images",
	".png":  "Images",
	".gif":  "Images",
	".pdf":  "Documents",
	".doc":  "Documents",
	".docx": "Documents",
	".txt":  "TextFiles",
	".mp4":  "Videos",
	".mp3":  "Audio",
}

func main() {
	// Target directory containing unorganized files
	targetDir := "./testfolder"

	// Read all entries (files and folders) from the target directory
	files, err := os.ReadDir(targetDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Iterate over each file in the directory
	for _, file := range files {
		// Skip subdirectories
		if file.IsDir() {
			continue
		}

		// Get the file's extension (e.g., ".jpg", ".txt")
		ext := filepath.Ext(file.Name())

		// Look up the target folder for the extension
		folderName, ok := extensionMap[ext]
		if !ok {
			// Default to "Others" if the extension is unknown
			folderName = "Others"
		}

		// Construct the full path for the destination directory
		destDir := filepath.Join(targetDir, folderName)

		// Create the destination folder if it doesn't already exist
		if _, err := os.Stat(destDir); os.IsNotExist(err) {
			err := os.Mkdir(destDir, 0755)
			if err != nil {
				fmt.Println("Error creating folder:", err)
				continue
			}
		}

		// Define full paths for the source and destination
		oldPath := filepath.Join(targetDir, file.Name())
		newPath := filepath.Join(destDir, file.Name())

		// Move (rename) the file to the destination folder
		err = os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("Failed to move %s: %v\n", file.Name(), err)
		} else {
			fmt.Printf("Moved %s → %s/\n", file.Name(), folderName)
		}
	}
}
