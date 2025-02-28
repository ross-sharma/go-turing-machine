package turing

const (
	HALT_STATE  = ""
	SHIFT_LEFT  = "L"
	SHIFT_RIGHT = "R"
)

type Instruction struct {
	write rune
	shift string
	next  string
}

type State struct {
	id           string
	instructions map[rune]*Instruction
}

type Runner struct {
	machine      *TuringMachine
	currentState *State
	states       map[string]*State
}

func CreateRunner(states []*State) *Runner {
	runner := &Runner{
		machine: CreateTuring(),
		states:  make(map[string]*State),
	}

	for _, state := range states {
		runner.states[state.id] = state
	}

	if len(states) > 0 {
		runner.currentState = states[0]
	}

	return runner
}

func (r *Runner) Step() {
	symbol := r.machine.GetCurrentSymbol()
	instruction := r.currentState.instructions[symbol]

	r.machine.Write(instruction.write)

	if instruction.shift == SHIFT_LEFT {
		r.machine.ShiftLeft()
	} else {
		r.machine.ShiftRight()
	}

	if instruction.next == HALT_STATE {
		r.currentState = nil
	} else {
		r.currentState = r.states[instruction.next]
	}
}

func (r *Runner) RunUntilHalt() {
	for !r.IsHalted() {
		r.Step()
	}
}

func (r *Runner) IsHalted() bool {
	return r.currentState == nil
}

func (r *Runner) GetStateId() string {
	if r.currentState == nil {
		return HALT_STATE
	} else {
		return r.currentState.id
	}
}

func (r *Runner) GetTape() string {
	return r.machine.GetTape()
}
