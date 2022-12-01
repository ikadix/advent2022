package advent

import (
	"fmt"
	"sync"
)

type Day interface {
	Setup() error
	Number() int
	Part1() (string, error)
	Part2() (string, error)
}

type Advent struct {
	wg      *sync.WaitGroup
	mut     *sync.RWMutex
	results [25][2]string
}

func (a *Advent) runDay(d Day) {
	defer a.wg.Done()

	if err := d.Setup(); err != nil {
		fmt.Println(err)
		return
	}

	firstResult, err := d.Part1()
	if err != nil {
		fmt.Println(err)
		return
	}

	a.mut.Lock()
	a.results[d.Number()-1][0] = firstResult
	a.mut.Unlock()

	secondResult, err := d.Part2()
	if err != nil {
		fmt.Println(err)
		return
	}

	a.mut.Lock()
	a.results[d.Number()-1][1] = secondResult
	a.mut.Unlock()
}

func (a *Advent) printCalendar() {
	for dayNum, day := range a.results {
		for partNum, part := range day {
			if part != "" {
				fmt.Printf("Day %d | Part %d = %s\n",
					dayNum+1,
					partNum+1,
					part,
				)
			}
		}
	}
}

func (a *Advent) RunCalendar(days []Day) {
	for _, day := range days {
		a.wg.Add(1)
		go a.runDay(day)
	}
	a.wg.Wait()

	a.printCalendar()
}

func NewAdvent() *Advent {
	return &Advent{
		wg:      &sync.WaitGroup{},
		mut:     &sync.RWMutex{},
		results: [25][2]string{},
	}
}
