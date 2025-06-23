package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Ik-cyber/file-organizer/internal/organizer"
)

func main() {
	flag.Parse()
	args := flag.Args()
	// targetDir := "./testfolder"

	var targetDir string

	if len(args) >= 1 {
		targetDir = args[0]
	} else {
		cwd, err := os.Getwd()
		if err != nil {
			log.Fatalf("Could not get current working directory: %v", err)
		}
		targetDir = cwd
	}

	absPath, err := filepath.Abs(targetDir)
	if err != nil {
		log.Fatalf("Could not resolve path ath: %v", err)
	}
	fmt.Println("Organizing files at: ", absPath)

	err = organizer.OrganizeFiles(targetDir)
	if err != nil {
		log.Fatalf("Error organizing files: %v", err)
	}
}
