package day7

import "strconv"

const (
	totalSpace     = 70000000
	requiredSpaced = 30000000
)

func (d *Day) isEnough(size int) bool {
	return (totalSpace-d.Root.Size)+size > requiredSpaced
}

func (d *Day) getDeleteble(root *Directory) []*Directory {
	deleteble := []*Directory{}

	for _, dir := range root.Directories {
		if d.isEnough(dir.Size) {
			deleteble = append(deleteble, dir)
		}
		deleteble = append(deleteble, d.getDeleteble(dir)...)
	}

	return deleteble
}

func getSmallestDir(dirs []*Directory) *Directory {
	smallest := dirs[0]

	for _, dir := range dirs {
		if dir.Size < smallest.Size {
			smallest = dir
		}
	}

	return smallest
}

// Part2 is the solution to part 2 of the day's puzzle.
func (d *Day) Part2() (string, error) {

	deleteable := d.getDeleteble(d.Root)

	if d.isEnough(d.Root.Size) {
		deleteable = append(deleteable, d.Root)
	}

	smallest := getSmallestDir(deleteable)

	return strconv.Itoa(smallest.Size), nil
}
