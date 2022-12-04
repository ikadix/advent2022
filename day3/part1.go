package day3

import (
	"strconv"
)

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	priority := 0

	for _, r := range d.Rucksacks {
		priority += r.getOverlappingItemPriorities(d.PriorityMap)
	}

	return strconv.Itoa(priority), nil
}
