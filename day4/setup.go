package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func discoverTasksFromInput(taskRange string) [2]int {
	tasks := [2]int{}

	boundaries := strings.Split(taskRange, "-")

	var err error

	tasks[0], err = strconv.Atoi(boundaries[0])
	if err != nil {
		log.Fatal(err)
	}

	tasks[1], err = strconv.Atoi(boundaries[1])
	if err != nil {
		log.Fatal(err)
	}

	return tasks
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

	for fileScanner.Scan() {
		line := fileScanner.Text()

		ranges := strings.Split(line, ",")

		d.Elves = append(d.Elves, [2][2]int{
			discoverTasksFromInput(ranges[0]),
			discoverTasksFromInput(ranges[1]),
		})
	}

	return nil
}
