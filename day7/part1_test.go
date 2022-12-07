package day7_test

import (
	"testing"

	"github.com/ikadix/advent2022/day7"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	day := &day7.Day{
		InputFile: "example.txt",
		Root: &day7.Directory{
			Files:       map[string]int{},
			Directories: map[string]*day7.Directory{},
		},
	}

	if err := day.Setup(); err != nil {
		t.Errorf("error setting up day: %v", err)
	}

	result, err := day.Part1()
	if err != nil {
		t.Errorf("error running part 1: %v", err)
	}

	if result != "95437" {
		t.Errorf("expected result to be 95437, got %s", result)
	}
}
