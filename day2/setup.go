package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Setup sets up any required data for the days puzzle.
func (d *Day) Setup() error {
	readFile, err := os.Open(d.InputFile)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}

	defer func() {
		if err := readFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	match := &Match{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		moves := strings.Split(line, " ")
		match.OpponentMove = mapInputToHand(moves[0])
		match.OurMove = mapInputToHand(moves[1])
		match.NeededResult = mapInputToResult(moves[1])

		d.Matches = append(d.Matches, match)

		match = &Match{}
	}

	return nil
}
