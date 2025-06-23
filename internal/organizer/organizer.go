package organizer

import (
	"fmt"
	"os"
	"path/filepath"
)

func OrganizeFiles(targetDir string) error {
	files, err := os.ReadDir(targetDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := filepath.Ext(file.Name())
		folderName, ok := extensionMap[ext]
		if !ok {
			folderName = "Others"
		}

		destDir := filepath.Join(targetDir, folderName)
		if _, err := os.Stat(destDir); os.IsNotExist(err) {
			err := os.Mkdir(destDir, 0755)
			if err != nil {
				fmt.Println("Error creating folder:", err)
				continue
			}
		}

		oldPath := filepath.Join(targetDir, file.Name())
		newPath := filepath.Join(destDir, file.Name())

		err = os.Rename(oldPath, newPath)
		if err != nil {
			fmt.Printf("Failed to move %s: %v\n", file.Name(), err)
		} else {
			fmt.Printf("Moved %s → %s/\n", file.Name(), folderName)
		}
	}

	return nil
}
