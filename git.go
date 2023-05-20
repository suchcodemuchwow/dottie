package main

import (
	"github.com/go-git/go-git/v5"
	"log"
	"os"
)

func isGitRepository(path string) bool {
	err := os.Chdir(path)
	if err != nil {
		return false
	}

	_, err = git.PlainOpen(path)
	if err != nil {
		return false
	}

	return true
}

func initGit(path string) {
	if isGitRepository(path) {
		return
	}

	_, err := git.PlainInit(path, false)
	if err != nil {
		log.Fatalf("Error: Failed to initialize git repository: %v", err)
	}
}
