package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

	elf := &Elf{}

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			d.Elves = append(d.Elves, elf)
			elf = &Elf{}

			continue
		}

		food, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("error converting line to int: %w", err)
		}

		elf.Calories += food

		elf.Foods = append(elf.Foods, food)
	}

	// Add the last elf
	d.Elves = append(d.Elves, elf)

	sort.Slice(d.Elves, func(i, j int) bool {
		return d.Elves[i].Calories > d.Elves[j].Calories
	})

	return nil
}
