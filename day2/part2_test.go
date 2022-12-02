package day2_test

import (
	"testing"

	"github.com/ikadix/advent2022/day2"
)

func TestPart2(t *testing.T) {
	t.Parallel()

	day := &day2.Day{
		InputFile: "example.txt",
		Matches:   []*day2.Match{},
	}

	if err := day.Setup(); err != nil {
		t.Fatal(err)
	}

	got, err := day.Part2()
	if err != nil {
		t.Fatal(err)
	}

	if got != "12" {
		t.Fatalf("got %s, want 12", got)
	}
}
