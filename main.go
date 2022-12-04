package main

import (
	"github.com/ikadix/advent2022/advent"
	"github.com/ikadix/advent2022/day1"
	"github.com/ikadix/advent2022/day2"
	"github.com/ikadix/advent2022/day3"
)

func main() {
	advent.NewAdvent().RunCalendar([]advent.Day{
		day1.New(),
		day2.New(),
		day3.New(),
	})
}
