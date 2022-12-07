package day7

import (
	"strconv"
)

func getAllDirectorySizes(root *Directory) []int {
	sizes := []int{}
	for _, dir := range root.Directories {
		sizes = append(sizes, dir.Size)
		sizes = append(sizes, getAllDirectorySizes(dir)...)
	}

	return sizes
}

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	totalSize := 0
	for _, size := range getAllDirectorySizes(d.Root) {
		if size <= 100000 {
			totalSize += size
		}
	}

	return strconv.Itoa(totalSize), nil
}
