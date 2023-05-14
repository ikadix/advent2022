// Package day4 implements the solution to Day 4 of the Advent of Code 2022.
package day4

type (
	// Day represents the data required for this days challenge.
	Day struct {
		Elves     [][2][2]int
		InputFile string
	}
)

// New returns a new instance of Day.
func New() *Day {
	return &Day{
		Elves:     [][2][2]int{},
		InputFile: "./day4/input.txt",
	}
}

// Number returns the day number.
//
//nolint:gomnd // This is the day number used by the advent package to print the day number.
func (d *Day) Number() int {
	return 4
}
