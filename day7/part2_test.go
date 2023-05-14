package day7_test

import (
	"testing"

	"github.com/ikadix/advent2022/day7"
)

func TestPart2(t *testing.T) {
	t.Parallel()

	day := &day7.Day{
		InputFile: "example.txt",
		Root: &day7.Directory{
			Name:        "/",
			Files:       map[string]int{},
			Directories: map[string]*day7.Directory{},
		},
	}

	if err := day.Setup(); err != nil {
		t.Errorf("error setting up day: %v", err)
	}

	_, err := day.Part1()
	if err != nil {
		t.Errorf("error running part 1: %v", err)
	}

	result, err := day.Part2()
	if err != nil {
		t.Errorf("error running part 2: %v", err)
	}

	if result != "24933642" {
		t.Fatal("incorrect result: ", result)
	}

}
