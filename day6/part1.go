package day6

import (
	"strconv"
)

func (d *Day) getDataSegment(i int, length int) []string {
	var segment []string
	if len(d.DataStream[i:]) < length {
		segment = d.DataStream[i:]
	} else {
		segment = d.DataStream[i : i+length]
	}

	return segment
}

// unique returns true if all the characters in the segment are unique.
func unique(segment []string) bool {
	seen := make(map[string]bool)
	for _, c := range segment {
		if seen[c] {
			return false
		}
		seen[c] = true
	}
	return true
}

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	for i := range d.DataStream {
		segment := d.getDataSegment(i, 4)
		if unique(segment) {
			return strconv.Itoa(i + 4), nil
		}
	}

	return "", nil
}
