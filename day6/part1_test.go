package day6_test

import (
	"strings"
	"testing"

	"github.com/ikadix/advent2022/day6"
)

func TestPart1(t *testing.T) {
	t.Parallel()

	var (
		exampleSignals = []string{
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			"nppdvjthqldpwncqszvftbrmjlhg",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		}
		exampleAnswers = []string{
			"7",
			"5",
			"6",
			"10",
			"11",
		}
	)

	for i := range exampleSignals {
		t.Run(exampleSignals[i], func(t *testing.T) {
			day := day6.Day{
				DataStream: strings.Split(exampleSignals[i], ""),
			}

			ans, err := day.Part1()
			if err != nil {
				t.Fatal(err)
			}

			if ans != exampleAnswers[i] {
				t.Fatalf("wrong answer, expected %v got %v", exampleAnswers[i], ans)
			}
		})
	}

}
