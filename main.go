package main

import (
	"fmt"
	"os"
	"path/filepath"
)

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
	targetDir := "./testfolder"

	files, err := os.ReadDir(targetDir)

	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		ext := filepath.Ext(file.Name()) // Gets Extention
		folderName, ok := extensionMap[ext]
		if !ok {
			folderName = "Others"
		}
		// Create the folder if it doesn't exist
		destDir := filepath.Join(targetDir, folderName)
		if _, err := os.Stat(destDir); os.IsNotExist(err) {
			err := os.Mkdir(destDir, 0755)
			if err != nil {
				fmt.Println("Error creating folder:", err)
				continue
			}
		}
		// Move the file
		oldPath := filepath.Join(targetDir, file.Name())
		newPath := filepath.Join(destDir, file.Name())

		err = os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("Failed to move %s: %v\n", file.Name(), err)
		} else {
			fmt.Printf("Moved %s → %s/\n", file.Name(), folderName)
		}
	}
}
