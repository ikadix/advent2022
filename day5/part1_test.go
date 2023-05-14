package day5_test

import (
	"testing"

	"github.com/ikadix/advent2022/day5"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	day := &day5.Day{
		InputFile: "example.txt",
	}

	day.Setup()

	answer, err := day.Part1()
	if err != nil {
		t.Fatal(err)
	}

	if answer != "CMZ" {
		t.Fail()
	}
}
