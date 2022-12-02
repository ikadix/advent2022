package day2

import (
	"strconv"
)

// Part2 is the solution to part 2 of the day's puzzle.
func (d *Day) Part2() (string, error) {
	score := 0

	for _, match := range d.Matches {
		match.calculateNeededMove()

		score += match.NeededResult

		score += match.NeededMove
	}

	return strconv.Itoa(
		score,
	), nil
}
