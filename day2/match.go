package day2

type (
	// Match contains all the information to work with a match result
	// from the strategy guide.
	// Including our initial (wrong) assumptions and the correct interpretation.
	Match struct {
		OpponentMove int
		OurMove      int
		Result       int
		NeededResult int
		NeededMove   int
	}
)

const (
	// potential moves.
	rock     = 1
	paper    = 2
	scissors = 3

	// potential outcomes.
	lost = 0
	draw = 3
	win  = 6
)

func mapInputToHand(input string) int {
	switch input {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors
	default:
		return 0
	}
}

func mapInputToResult(input string) int {
	switch input {
	case "X":
		return lost
	case "Y":
		return draw
	case "Z":
		return win
	default:
		return 0
	}
}

func (m *Match) calculateResult() {
	if m.OpponentMove == m.OurMove {
		m.Result = draw

		return
	}

	switch m.OpponentMove {
	case rock:
		if m.OurMove == paper {
			m.Result = win
		} else {
			m.Result = lost
		}
	case paper:
		if m.OurMove == scissors {
			m.Result = win
		} else {
			m.Result = lost
		}
	case scissors:
		if m.OurMove == rock {
			m.Result = win
		} else {
			m.Result = lost
		}
	}
}

func (m *Match) calculateNeededMove() {
	if m.NeededResult == draw {
		m.NeededMove = m.OpponentMove

		return
	}

	if m.NeededResult == win {
		switch m.OpponentMove {
		case rock:
			m.NeededMove = paper
		case paper:
			m.NeededMove = scissors
		case scissors:
			m.NeededMove = rock
		}
	} else {
		switch m.OpponentMove {
		case rock:
			m.NeededMove = scissors
		case paper:
			m.NeededMove = rock
		case scissors:
			m.NeededMove = paper
		}
	}
}
