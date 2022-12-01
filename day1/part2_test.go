package day1_test

import (
	"testing"

	"github.com/ikadix/advent2022/day1"
)

func TestPart2(t *testing.T) {
	t.Parallel()

	day := &day1.Day{
		InputFile: "./day1/example.txt",
	}

	if err := day.Setup(); err != nil {
		t.Fatalf("error setting up day: %s", err)
	}

	result, err := day.Part2()
	if err != nil {
		t.Fatalf("error running part 2 %s", err)
	}

	if result != "45000" {
		t.Fatalf("expected result to be 45000, got %s", result)
	}
}
