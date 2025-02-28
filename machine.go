package turing

import "fmt"

type TuringMachine struct {
	buffer   []rune
	index    int
	minIndex int
	maxIndex int
}

func CreateTuring() *TuringMachine {
	startSize := 8
	startIndex := startSize / 2

	tm := &TuringMachine{
		buffer:   make([]rune, startSize),
		index:    startIndex,
		minIndex: startIndex,
		maxIndex: startIndex,
	}

	tm.buffer[startIndex] = '0'
	return tm
}

func (tm *TuringMachine) Write(symbol rune) {
	tm.buffer[tm.index] = symbol
}

func (tm *TuringMachine) ShiftLeft() {
	tm.index += 1

	if tm.index >= len(tm.buffer) {
		tm.buffer = append(tm.buffer, '0')
	}

	if tm.index > tm.maxIndex {
		tm.maxIndex = tm.index
		tm.Write('0')
	}
}

func (tm *TuringMachine) ShiftRight() {
	tm.index -= 1

	if tm.index < 0 {
		tm.buffer = append([]rune{'0'}, tm.buffer...)
		tm.index = 0
		tm.maxIndex += 1
	}

	if tm.index < tm.minIndex {
		tm.minIndex = tm.index
		tm.Write('0')
	}
}

func (tm *TuringMachine) PrintTape() {
	for i := range tm.buffer {
		fmt.Printf("%2d ", i)
	}
	fmt.Println()
	for _, symbol := range tm.buffer {
		if symbol == 0 {
			fmt.Print(" -")
		} else {
			fmt.Printf("%2c ", symbol)
		}
	}
	fmt.Println()
}

func (tm *TuringMachine) GetTape() string {
	return string(tm.buffer[tm.minIndex : tm.maxIndex+1])
}

func (tm *TuringMachine) GetCurrentSymbol() rune {
	return tm.buffer[tm.index]
}
