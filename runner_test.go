package turing

import (
	"testing"
)

func assertRunner(
	t *testing.T,
	runner *Runner,
	isHalted bool,
	stateId string,
	expectedTape string,
) {
	assertMachine(t, runner.machine, expectedTape)

	if isHalted != runner.IsHalted() {
		t.Errorf("isHalted was %t", runner.IsHalted())
	}

	if stateId != runner.GetStateId() {
		t.Errorf(
			"stateId is \"%s\", expected \"%s\"",
			runner.currentState.id, stateId,
		)
	}
}

func assertSymbolCount(t *testing.T, tape string, symbol rune, expectedCount int) {
	actualCount := 0
	for _, s := range tape {
		if s == symbol {
			actualCount += 1
		}
	}
	if actualCount != expectedCount {
		t.Errorf(
			"Found %d of symbol %c, expected %d. Tape: %s",
			actualCount, symbol, expectedCount, tape,
		)
	}
}

func TestCreateRunner(t *testing.T) {
	states := []*State{
		{
			id: "A",
			instructions: map[rune]*Instruction{
				'0': {
					write: '0',
					shift: SHIFT_RIGHT,
					next:  "A",
				},
				'1': {
					write: '0',
					shift: SHIFT_RIGHT,
					next:  "A",
				},
			},
		},
	}

	runner := CreateRunner(states)
	assertRunner(t, runner, false, "A", "0")
}

func TestHalt(t *testing.T) {
	states := []*State{
		{
			id: "A",
			instructions: map[rune]*Instruction{
				'0': {
					write: '0',
					shift: SHIFT_RIGHT,
					next:  HALT_STATE,
				},
				'1': {
					write: '0',
					shift: SHIFT_RIGHT,
					next:  HALT_STATE,
				},
			},
		},
	}
	runner := CreateRunner(states)
	runner.Step()
	assertRunner(t, runner, true, HALT_STATE, "00")
}

func TestLoop(t *testing.T) {
	states := []*State{
		{
			id: "A",
			instructions: map[rune]*Instruction{
				'0': {
					write: '0',
					shift: SHIFT_LEFT,
					next:  "B",
				},
				'1': {
					write: '0',
					shift: SHIFT_LEFT,
					next:  "B",
				},
			},
		},
		{
			id: "B",
			instructions: map[rune]*Instruction{
				'0': {
					write: '1',
					shift: SHIFT_RIGHT,
					next:  "A",
				},
				'1': {
					write: '1',
					shift: SHIFT_RIGHT,
					next:  "A",
				},
			},
		},
	}
	runner := CreateRunner(states)

	check := func(stateId string, tape string) {
		assertRunner(t, runner, false, stateId, tape)
	}

	check("A", "0")
	runner.Step()
	check("B", "00")
	runner.Step()
	check("A", "01")
	runner.Step()
	check("B", "01")
}

func Test3StateBusyBeaver(t *testing.T) {
	states := []*State{
		{
			id: "A",
			instructions: map[rune]*Instruction{
				'0': {
					write: '1',
					shift: SHIFT_RIGHT,
					next:  "B",
				},
				'1': {
					write: '1',
					shift: SHIFT_RIGHT,
					next:  HALT_STATE,
				},
			},
		},
		{
			id: "B",
			instructions: map[rune]*Instruction{
				'0': {
					write: '0',
					shift: SHIFT_RIGHT,
					next:  "C",
				},
				'1': {
					write: '1',
					shift: SHIFT_RIGHT,
					next:  "B",
				},
			},
		},
		{
			id: "C",
			instructions: map[rune]*Instruction{
				'0': {
					write: '1',
					shift: SHIFT_LEFT,
					next:  "C",
				},
				'1': {
					write: '1',
					shift: SHIFT_LEFT,
					next:  "A",
				},
			},
		},
	}
	runner := CreateRunner(states)
	runner.RunUntilHalt()
	assertSymbolCount(t, runner.GetTape(), '1', 6)
}
