package day6

import "strconv"

// Part2 is the solution to part 2 of the day's puzzle.
func (d *Day) Part2() (string, error) {
	for i := range d.DataStream {
		segment := d.getDataSegment(i, 14)
		if unique(segment) {
			return strconv.Itoa(i + 14), nil
		}
	}

	return "", nil
}
