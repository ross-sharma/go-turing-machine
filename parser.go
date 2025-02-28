package turing

import (
	"strings"
)

/*
A 0:0RB 1:0R
B 0:1LB 1:1L
...
*/

func ParseRunner(s string) *Runner {
	lines := strings.Split(s, "\n")
	lines = arrayMap(lines, strings.TrimSpace)
	lines = arrayFilter(lines, isNotEmpty)

	states := make([]*State, 0, len(lines))
	for _, line := range lines {
		states = append(states, parseState(line))
	}

	return CreateRunner(states)
}

func parseState(line string) *State {
	parts := strings.Split(line, " ")
	parts = arrayFilter(parts, isNotEmpty)
	id := parts[0]
	instructions := parseInstructions(parts[1:])
	return &State{id, instructions}
}

func parseInstructions(tokens []string) map[rune]*Instruction {
	output := make(map[rune]*Instruction)
	for _, token := range tokens {
		parts := strings.Split(token, ":")
		symbol := rune(parts[0][0])
		instruction := parseInstruction(parts[1])
		output[symbol] = instruction
	}

	return output
}

func parseInstruction(s string) *Instruction {
	write := rune(s[0])

	shift := SHIFT_RIGHT
	if shiftChar := rune(s[1]); shiftChar == 'L' {
		shift = SHIFT_LEFT
	}

	next := HALT_STATE
	if len(s) == 3 {
		next = string(s[2])
	}

	return &Instruction{write, shift, next}
}

func arrayFilter[T any](objects []T, callback func(T) bool) []T {
	output := make([]T, 0)
	for _, obj := range objects {
		if callback(obj) {
			output = append(output, obj)
		}
	}
	return output
}

func arrayMap[T any, U any](objects []T, callback func(T) U) []U {
	output := make([]U, len(objects))
	for i, obj := range objects {
		output[i] = callback(obj)
	}
	return output
}

func isNotEmpty(s string) bool { return s != "" }
