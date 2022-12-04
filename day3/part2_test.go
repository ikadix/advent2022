package day3_test

import (
	"testing"

	"github.com/ikadix/advent2022/day3"
)

func TestPart2(t *testing.T) {
	t.Parallel()

	day := &day3.Day{
		InputFile:   "example.txt",
		GroupSize:   3,
		PriorityMap: day3.CreatePriorityMap(),
	}

	if err := day.Setup(); err != nil {
		t.Fatalf("error setting up day: %v", err)
	}

	result, err := day.Part2()
	if err != nil {
		t.Fatalf("error running part 1: %v", err)
	}

	if result != "70" {
		t.Fatalf("expected result to be 70, got %s", result)
	}
}
