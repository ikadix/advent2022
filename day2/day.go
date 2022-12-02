// Package day2 implements the solution to Day 2 of the Advent of Code 2022.
package day2

// Day contains the path to our input file
// and the matches we've parsed from it.
type Day struct {
	InputFile string
	Matches   []*Match
}

// New returns a new instance of Day.
func New() *Day {
	return &Day{
		InputFile: "./day2/input.txt",
		Matches:   []*Match{},
	}
}

// Number returns the day number.
func (d *Day) Number() int {
	return 2
}
