package main

import (
	"github.com/ikadix/advent2022/advent"
	"github.com/ikadix/advent2022/day1"
)

func main() {
	advent.NewAdvent().RunCalendar([]advent.Day{
		day1.New(),
	})
}
