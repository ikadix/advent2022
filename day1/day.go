// Package day1 implements the solution to Day 1 of the Advent of Code 2022.
package day1

type Day struct{
	InputFile string
}

// New returns a new instance of Day.
func New() *Day {
	return &Day{}
}

// Number returns the day number.
func (d *Day) Number() int {
	return 1
}

