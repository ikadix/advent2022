package day6_test

import (
	"strings"
	"testing"

	"github.com/ikadix/advent2022/day6"
)

func TestPart2(t *testing.T) {
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
			"19",
			"23",
			"23",
			"29",
			"26",
		}
	)

	for index := range exampleSignals {
		index := index
		t.Run(exampleSignals[index], func(t *testing.T) {
			t.Parallel()

			day := day6.Day{
				DataStream: strings.Split(exampleSignals[index], ""),
			}

			ans, err := day.Part2()
			if err != nil {
				t.Fatal(err)
			}

			if ans != exampleAnswers[index] {
				t.Fatalf("wrong answer, expected %v got %v", exampleAnswers[index], ans)
			}
		})
	}
}
