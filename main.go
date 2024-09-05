package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func searchFilesRecursively(dirPath string, searchString string) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		filePath := filepath.Join(dirPath, file.Name())

		if file.IsDir() {
			searchFilesRecursively(filePath, searchString)
		} else {
			readFileAndSearch(filePath, searchString)
		}
	}
}

func readFileAndSearch(filePath string, searchString string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchString) {
			fmt.Printf("Found \"%s\" in %s at line %d\n", searchString, filePath, lineNumber)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}

func main() {
	dirPath := "svg" // Replace with the actual directory path
	searchString := "<script"

	searchFilesRecursively(dirPath, searchString)
}
