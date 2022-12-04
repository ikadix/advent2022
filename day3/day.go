// Package day3 implements the solution to Day 3 of the Advent of Code 2022.
package day3

type (
	// Compartment represents a compartment of a rucksack.
	Compartment struct {
		RawItems    []string
		MappedItems map[string]int
	}

	// Rucksack represents a rucksack.
	Rucksack struct {
		Compartments [2]Compartment
		AllItems     map[string]int
		Badge        string
	}

	// Day represents the solution to Day 3 of the Advent of Code 2022.
	Day struct {
		PriorityMap map[string]int
		Rucksacks   []*Rucksack
		GroupSize   int
		InputFile   string
	}
)

func (r *Rucksack) getOverlappingItemPriorities(priorityMap map[string]int) int {
	priority := 0

	for item := range r.Compartments[0].MappedItems {
		if _, ok := r.Compartments[1].MappedItems[item]; ok {
			priority += priorityMap[item]
		}
	}

	return priority
}

// CreatePriorityMap returns a map of item to priority.
//
//nolint:gomnd // This is a puzzle.
func CreatePriorityMap() map[string]int {
	return map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
		"A": 27,
		"B": 28,
		"C": 29,
		"D": 30,
		"E": 31,
		"F": 32,
		"G": 33,
		"H": 34,
		"I": 35,
		"J": 36,
		"K": 37,
		"L": 38,
		"M": 39,
		"N": 40,
		"O": 41,
		"P": 42,
		"Q": 43,
		"R": 44,
		"S": 45,
		"T": 46,
		"U": 47,
		"V": 48,
		"W": 49,
		"X": 50,
		"Y": 51,
		"Z": 52,
	}
}

// New returns a new instance of Day.
//
//nolint:gomnd // groupsize of 3 is defined in the puzzle.
func New() *Day {
	return &Day{
		PriorityMap: CreatePriorityMap(),
		Rucksacks:   []*Rucksack{},
		GroupSize:   3,
		InputFile:   "./day3/input.txt",
	}
}

// Number returns the day number.
//
//nolint:gomnd // This is the day number used by the advent package to print the day number.
func (d *Day) Number() int {
	return 3
}
