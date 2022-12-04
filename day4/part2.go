package day4

import (
	"strconv"
)

// Part2 is the solution to part 2 of the day's puzzle.
func (d *Day) Part2() (string, error) {
	contained := 0

	for _, pair := range d.Elves {
		if (pair[0][0] >= pair[1][0] &&
			pair[0][0] <= pair[1][1]) ||
			(pair[1][0] >= pair[0][0] &&
				pair[1][0] <= pair[0][1]) {
			contained++
		}
	}

	return strconv.Itoa(contained), nil
}
