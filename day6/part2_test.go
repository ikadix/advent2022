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

	for i := range exampleSignals {
		t.Run(exampleSignals[i], func(t *testing.T) {
			day := day6.Day{
				DataStream: strings.Split(exampleSignals[i], ""),
			}

			ans, err := day.Part2()
			if err != nil {
				t.Fatal(err)
			}

			if ans != exampleAnswers[i] {
				t.Fatalf("wrong answer, expected %v got %v", exampleAnswers[i], ans)
			}
		})
	}

}
