// Package day7 implements the solution to Day 7 of the Advent of Code 2022.
package day7

import (
	"fmt"
)

type (
	// Directory represents a directory in the file system.
	Directory struct {
		Files       map[string]int
		Directories map[string]*Directory
		Size        int
	}

	// Day contains the data for the solution.
	Day struct {
		Root      *Directory
		InputFile string
	}
)

// Print the directory tree.
//
//nolint:forbidigo // This is a debug function.
func (d *Directory) Print(indent string) {
	for name, dir := range d.Directories {
		fmt.Println(indent, name, dir.Size)
		dir.Print(indent + "  ")
	}

	for file, size := range d.Files {
		fmt.Println(indent, file, size)
	}
}

// CalculateSize calculates the size of the directory and all subdirectories.
func (d *Directory) CalculateSize() int {
	size := d.Size

	for _, dir := range d.Directories {
		size += dir.CalculateSize()
	}

	d.Size = size

	return size
}

// New returns a new instance of Day.
func New() *Day {
	return &Day{
		Root: &Directory{
			Files:       make(map[string]int),
			Directories: make(map[string]*Directory),
		},
		InputFile: "./day7/input.txt",
	}
}

// Number returns the day number.
//
//nolint:gomnd // This is the day number used by the advent package to print the day number.
func (d *Day) Number() int {
	return 7
}
