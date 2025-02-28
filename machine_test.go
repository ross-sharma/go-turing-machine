package turing

import (
	"testing"
)

func assertMachine(t *testing.T, tm *TuringMachine, expectedTape string) {
	tape := tm.GetTape()
	if tape != expectedTape {
		t.Errorf("Tape was: '%s' (expected '%s')", tape, expectedTape)
	}
}

func TestCreation(t *testing.T) {
	tm := CreateTuring()
	assertMachine(t, tm, "0")
}

func TestWrite(t *testing.T) {
	tm := CreateTuring()
	tm.Write('1')
	assertMachine(t, tm, "1")
	tm.Write('0')
	assertMachine(t, tm, "0")
}

func TestShiftLeft(t *testing.T) {
	tm := CreateTuring()
	assertMachine(t, tm, "0")
	tm.ShiftLeft()
	assertMachine(t, tm, "00")
	tm.ShiftLeft()
	assertMachine(t, tm, "000")
	tm.ShiftLeft()
	assertMachine(t, tm, "0000")
}

func TestShiftRight(t *testing.T) {
	tm := CreateTuring()
	tm.ShiftRight()
	assertMachine(t, tm, "00")
	tm.ShiftRight()
	assertMachine(t, tm, "000")
	tm.ShiftRight()
	assertMachine(t, tm, "0000")
}

func TestShiftAndWrite(t *testing.T) {
	tm := CreateTuring()

	tm.Write('1')
	assertMachine(t, tm, "1")

	tm.ShiftRight()
	assertMachine(t, tm, "01")

	tm.Write('1')
	assertMachine(t, tm, "11")

	tm.ShiftLeft()
	assertMachine(t, tm, "11")

	tm.ShiftLeft()
	assertMachine(t, tm, "110")
}
