package day7

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func (d *Day) parseCommand(currentPath []string, currentDir *Directory, line string) ([]string, *Directory) {
	command := strings.Split(strings.TrimPrefix(line, "$ "), " ")
	switch command[0] {
	case "cd":
		if command[1] == "/" {
			currentPath = []string{}
			currentDir = d.Root

			return currentPath, currentDir
		}

		if command[1] == ".." {
			currentPath = currentPath[:len(currentPath)-1]
			currentDir = d.Root

			for _, dir := range currentPath {
				currentDir = currentDir.Directories[dir]
			}

			return currentPath, currentDir
		}

		currentPath = append(currentPath, command[1])

		if _, ok := currentDir.Directories[command[1]]; !ok {
			currentDir.Directories[command[1]] = &Directory{
				Files:       make(map[string]int),
				Directories: make(map[string]*Directory),
			}
		}

		currentDir = currentDir.Directories[command[1]]
	case "ls":
		return currentPath, currentDir
	}

	return currentPath, currentDir
}

// Setup sets up any required data for the days puzzle.
func (d *Day) Setup() error {
	readFile, err := os.Open(d.InputFile)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer func() {
		if err := readFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	currentPath := []string{}

	currentDir := d.Root

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if strings.HasPrefix(line, "$") {
			currentPath, currentDir = d.parseCommand(currentPath, currentDir, line)

			continue
		}

		if strings.HasPrefix(line, "dir") {
			currentDir.Directories[strings.TrimPrefix(line, "dir ")] = &Directory{
				Files:       make(map[string]int),
				Directories: make(map[string]*Directory),
			}

			continue
		}

		file := strings.Split(line, " ")

		size, err := strconv.Atoi(file[0])
		if err != nil {
			return fmt.Errorf("error converting size to int: %w", err)
		}

		currentDir.Files[file[1]] = size
		currentDir.Size += size
	}

	d.Root.CalculateSize()

	return nil
}
