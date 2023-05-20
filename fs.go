package main

import (
	cp "github.com/otiai10/copy"
	"io"
	"log"
	"os"
	"strings"
)

func listDotFilesAndFolders(path string) (files []string, folders []string) {
	dir, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error: Failed to open directory: %v", err)
	}
	defer dir.Close()

	dirContent, err := dir.ReadDir(0)
	if err != nil {
		log.Fatalf("Error: Failed to read directory: %v", err)
	}

	dotFiles := []string{}
	dotFolders := []string{}

	for _, file := range dirContent {
		if !strings.HasPrefix(file.Name(), ".") {
			continue
		}

		if file.IsDir() {
			dotFolders = append(dotFolders, file.Name())
		}

		dotFiles = append(dotFiles, file.Name())
	}

	return dotFiles, dotFolders
}

func compareFolders(srcDir, targetDir string) (toCopy []string) {
	srcFiles, srcFolders := listDotFilesAndFolders(srcDir)
	targetFiles, targetFolders := listDotFilesAndFolders(targetDir)

	diff := make(map[string]string)
	blacklist := map[string]bool{
		".git":      true,
		".DS_Store": true,
		".Trash":    true,
	}

	for _, targetFile := range targetFiles {
		diff[targetFile] = targetFile
	}

	for _, targetFolder := range targetFolders {
		diff[targetFolder] = targetFolder
	}

	for _, srcFile := range srcFiles {
		_, exist := diff[srcFile]
		_, isBlacklisted := blacklist[srcFile]

		if !exist && !isBlacklisted {
			toCopy = append(toCopy, srcFile)
		}
	}

	for _, srcFolder := range srcFolders {
		_, exist := diff[srcFolder]
		_, isBlacklisted := blacklist[srcFolder]

		if !exist && !isBlacklisted {
			toCopy = append(toCopy, srcFolder)
		}
	}

	return toCopy
}

func copyFolder(src, dest string) {
	err := cp.Copy(src, dest)
	if err != nil {
		log.Printf("Error copying %s to %s: %s", src, dest, err)
	}

	log.Println("Copied dir: ", src, "to: ", dest)
}

func copyFile(src, dest string) {
	fin, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	fout, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)

	if err != nil {
		log.Fatal(err)
	}
}

func move(src, dest string) {
	err := os.Rename(src, dest)
	if err != nil {
		log.Fatal(err)
	}
}

func symlink(src, dest string) {
	err := os.Symlink(src, dest)
	if err != nil {
		log.Fatal(err)
	}
}

func mkdir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatalf("Error: Failed to create directory '%s': %v", path, err)
		}
	}
}
