package day3

import (
	"strconv"
)

// Part2 is the solution to part 2 of the day's puzzle.
func (d *Day) Part2() (string, error) {
	totalBadgePriority := 0

	for i := 0; i < len(d.Rucksacks); i += d.GroupSize {
		totalBadgePriority += d.PriorityMap[d.Rucksacks[i].Badge]
	}

	return strconv.Itoa(totalBadgePriority), nil
}
