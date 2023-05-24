package day8

import (
	"log"
	"strconv"
	"sync"
	"sync/atomic"
)

func (d *Day) checkLeft(x int, y int, height int) bool {
	for i := 0; i < x; i++ {
		if d.TreeMap[y][i] >= height {
			return false
		}
	}

	return true
}

func (d *Day) checkRight(x int, y int, height int) bool {
	if x == len(d.TreeMap[y])-1 {
		return true
	}

	for i := x + 1; i < len(d.TreeMap[y]); i++ {
		if d.TreeMap[y][i] >= height {
			return false
		}
	}

	return true
}

func (d *Day) checkUp(x int, y int, height int) bool {
	for i := 0; i < y; i++ {
		if d.TreeMap[i][x] >= height {
			return false
		}
	}

	return true
}

func (d *Day) checkDown(x int, y int, height int) bool {
	if y == len(d.TreeMap)-1 {
		return true
	}

	for i := y + 1; i < len(d.TreeMap); i++ {
		if d.TreeMap[i][x] >= height {
			return false
		}
	}

	return true
}

func (d *Day) isVisible(xCoord int, yCoord int) bool {
	height := d.TreeMap[xCoord][yCoord]

	visibility := [4]bool{}

	mux := &sync.Mutex{}

	group := &sync.WaitGroup{}

	type checkFunc = func(int, int, int) bool

	checks := []checkFunc{
		d.checkLeft,
		d.checkRight,
		d.checkUp,
		d.checkDown,
	}

	for index, check := range checks {
		group.Add(1)

		go func(index int, check checkFunc) {
			defer group.Done()

			mux.Lock()
			defer mux.Unlock()

			visibility[index] = check(xCoord, yCoord, height)
		}(index, check)
	}

	group.Wait()

	for _, visible := range visibility {
		if visible {
			return true
		}
	}

	return false
}

// Part1 is the solution to part 1 of the day's puzzle.
func (d *Day) Part1() (string, error) {
	var visible int64

	var checked int64

	group := &sync.WaitGroup{}

	for yCoord := 0; yCoord < len(d.TreeMap); yCoord++ {
		for xCoord := 0; xCoord < len(d.TreeMap[yCoord]); xCoord++ {
			group.Add(1)

			go func(x int, y int) {
				defer group.Done()

				if d.isVisible(x, y) {
					atomic.AddInt64(&visible, 1)
				}

				atomic.AddInt64(&checked, 1)
			}(xCoord, yCoord)
		}
	}

	group.Wait()

	log.Println("Checked", checked, "squares out of ", len(d.TreeMap)*len(d.TreeMap[0]))

	return strconv.Itoa(int(visible)), nil
}
