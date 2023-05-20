package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("🚀 Welcome to dottie dotfiles manager")

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error: Failed to get user home directory: %v", err)
	}

	//initialSetup(homeDir)

	//moveFilesIfNotExist(homeDir, "/Users/dev/dottie")

	toCopy := compareFolders(homeDir, "/Users/dev/dottie")

}
