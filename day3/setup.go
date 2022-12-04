package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func mapItems(items []string) map[string]int {
	mappedItems := make(map[string]int)

	for _, item := range items {
		mappedItems[item]++
	}

	return mappedItems
}

func discoverRucksackFromItems(items []string) *Rucksack {
	return &Rucksack{
		AllItems: func() map[string]int {
			m := make(map[string]int)
			for _, item := range items {
				m[item]++
			}

			return m
		}(),
		Compartments: [2]Compartment{
			{
				RawItems:    items[:len(items)/2],
				MappedItems: mapItems(items[:len(items)/2]),
			},
			{
				RawItems:    items[len(items)/2:],
				MappedItems: mapItems(items[len(items)/2:]),
			},
		},
	}
}

func (d *Day) assignBadges(group []*Rucksack) []*Rucksack {
	badge := ""

	for item := range group[0].AllItems {
		_, inSecond := group[1].AllItems[item]
		_, inThird := group[2].AllItems[item]

		if inSecond && inThird {
			badge = item

			break
		}
	}

	for _, r := range group {
		r.Badge = badge
	}

	return group
}

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

	// Rucksacks are grouped into groups of 3s
	// This is to keep track of the current group
	currentGroup := []*Rucksack{}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		items := strings.Split(line, "")

		rucksack := discoverRucksackFromItems(items)

		currentGroup = append(currentGroup, rucksack)

		if len(currentGroup) == d.GroupSize {
			d.Rucksacks = append(d.Rucksacks, (d.assignBadges(currentGroup))...)
			currentGroup = []*Rucksack{}
		}
	}

	return nil
}
