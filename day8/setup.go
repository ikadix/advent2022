package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	d.TreeMap = [][]int{}

	for fileScanner.Scan() {
		line := []int{}

		for _, c := range strings.Split(fileScanner.Text(), "") {
			height, err := strconv.Atoi(c)
			if err != nil {
				return fmt.Errorf("error converting %s to int: %w", c, err)
			}

			line = append(line, height)
		}

		d.TreeMap = append(d.TreeMap, line)
	}

	return nil
}
