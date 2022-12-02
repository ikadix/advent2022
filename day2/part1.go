package day2

import (
	"strconv"
)

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	score := 0

	for _, match := range d.Matches {
		match.calculateResult()

		score += match.Result

		score += match.OurMove
	}

	return strconv.Itoa(
		score,
	), nil
}
