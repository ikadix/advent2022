package day8_test

import (
	"testing"

	"github.com/ikadix/advent2022/day8"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	day := day8.Day{
		InputFile: "example.txt",
	}

	if err := day.Setup(); err != nil {
		t.Fatal(err)
	}

	res, err := day.Part1()
	if err != nil {
		t.Fatal(err)
	}

	if res != "21" {
		t.Errorf("Expected 21, got %s", res)
	}
}
