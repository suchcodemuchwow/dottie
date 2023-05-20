package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
)

type Config struct {
	DotFile   string `json:"dot_file"`
	GitRemote string `json:"git_remote"`
	Interval  uint16 `json:"interval"`
}

func (cfg *Config) assignDefaults(homeDir string) {
	if cfg.GitRemote == "" {
		panic("Git remote is required")
	}

	if cfg.DotFile == "" {
		cfg.DotFile = filepath.Join(homeDir, "dotfiles")
	}

	if cfg.Interval == 0 {
		cfg.Interval = 60
	}
}

func ask(question string) string {
	var input string

	fmt.Printf("%s ", question)

	fmt.Scanln(&input)
	return input
}

func initialSetup(homeDir string) {
	config := Config{}

	config.DotFile = ask("📄 Enter folder name to save your dotfiles: (~/dotfiles)")
	config.GitRemote = ask("📦 Enter git remote url:")
	intervalStr := ask("♻️ How frequently (seconds) we should check: (60)")

	if intervalStr == "" {
		intervalStr = "0"
	}

	interval, err := strconv.Atoi(intervalStr)
	if err != nil {
		log.Fatalf("Error: Invalid interval. '%s' is not a valid number", intervalStr)
	}

	config.Interval = uint16(interval)

	config.assignDefaults(homeDir)

	arr, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		log.Fatalf("Error: Failed to marshal config to JSON: %v", err)
	}

	serialized := string(arr)
	println(serialized)
}
