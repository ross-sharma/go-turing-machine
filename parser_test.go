package turing

import (
	"testing"
)

func TestParse(t *testing.T) {
	runner := ParseRunner("A 0:0R 1:0R")
	assertRunner(t, runner, false, "A", "0")
	runner.Step()
	assertRunner(t, runner, true, HALT_STATE, "00")
}

func TestParse3StateBusyBeaver(t *testing.T) {
	code := `
	A 0:1RB 1:1R
	B 0:0RC 1:1RB
	C 0:1LC 1:1LA
	`

	runner := ParseRunner(code)
	runner.RunUntilHalt()
	assertSymbolCount(t, runner.GetTape(), '1', 6)
}

func TestParse4StateBusyBeaver(t *testing.T) {
	code := `
	A 0:1RB 1:1LB
	B 0:1LA 1:0LC
	C 0:1R  1:1LD
	D 0:1RD 1:0RA		
	`
	runner := ParseRunner(code)
	runner.RunUntilHalt()
	assertSymbolCount(t, runner.GetTape(), '1', 13)
}

func TestParse5StateBusyBeaver(t *testing.T) {
	code := `
	A 0:1RB 1:1LC
	B 0:1RC 1:1RB
	C 0:1RD 1:0LE
	D 0:1LA 1:1LD
	E 0:1R  1:0LA
	`
	runner := ParseRunner(code)
	runner.RunUntilHalt()
	assertSymbolCount(t, runner.GetTape(), '1', 4098)
}