package day1

import (
	"strconv"
)

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	return strconv.Itoa(d.Elves[0].Calories), nil
}
