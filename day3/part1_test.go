package day3_test

import (
	"testing"

	"github.com/ikadix/advent2022/day3"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	day := &day3.Day{
		InputFile:   "example.txt",
		PriorityMap: day3.CreatePriorityMap(),
		GroupSize:   3,
	}

	if err := day.Setup(); err != nil {
		t.Fatalf("error setting up day: %v", err)
	}

	result, err := day.Part1()
	if err != nil {
		t.Fatalf("error running part 1: %v", err)
	}

	if result != "157" {
		t.Fatalf("expected result to be 157, got %s", result)
	}
}
