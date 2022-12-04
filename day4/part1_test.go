package day4_test

import (
	"testing"

	"github.com/ikadix/advent2022/day4"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	day := &day4.Day{
		Elves:     [][2][2]int{},
		InputFile: "example.txt",
	}

	if err := day.Setup(); err != nil {
		t.Fatalf("error setting up: %s", err)
	}

	result, err := day.Part1()
	if err != nil {
		t.Fatalf("error running part 1: %s", err)
	}

	if result != "2" {
		t.Fatalf("expected 2 got %s", result)
	}
}
