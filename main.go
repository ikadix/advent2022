// Package main is the entry point to the advent2022 application.
package main

import (
	"github.com/ikadix/advent2022/advent"
	"github.com/ikadix/advent2022/day1"
	"github.com/ikadix/advent2022/day2"
	"github.com/ikadix/advent2022/day3"
	"github.com/ikadix/advent2022/day4"
	"github.com/ikadix/advent2022/day5"
	"github.com/ikadix/advent2022/day6"
	"github.com/ikadix/advent2022/day7"
)

func main() {
	advent.NewAdvent().RunCalendar([]advent.Day{
		day1.New(),
		day2.New(),
		day3.New(),
		day4.New(),
		day5.New(),
		day6.New(),
		day7.New(),
	})
}
