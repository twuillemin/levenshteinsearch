package levenshteinsearch

import (
	"fmt"
	"testing"
)

func TestVsPythonReference(t *testing.T) {

	countFoundState := 0

	words := []string{"banana", "bananas", "cabana", "foobarbazfoobarbaz", "a", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", ""}
	for n := 0; n < 5; n++ {
		for _, word := range words {
			sparse := CreateAutomaton(word, n)
			for _, query := range words {
				state := sparse.Start()
				for _, c := range query {
					if word == "banana" && query == "cabana" && n == 2 {
						countFoundState++
						fmt.Printf("state: %v\n", state)
					}
					state = sparse.Step(state, c)
				}
			}
		}
	}

	if countFoundState != 6 {
		t.Error("Expected to have 6 states like python does")
	}
}

func TestCreateDigraph(t *testing.T) {
	digraph := CreateDigraph("woof", 10)
	if len(digraph) != 1132 {
		t.Error("Expected a digraph definition with 1132 lines")
	}
}

func TestCanMatch(t *testing.T) {
	// Create the Automaton
	automaton := CreateAutomaton("bannana", 1)

	// Start it
	state0 := automaton.Start()

	// Check that "w" can match "bannana" with distance 1
	state1 := automaton.Step(state0, 'w')
	if !automaton.CanMatch(state1) {
		t.Error("Expected 'w' can match 'banana' with a distance of 1")
	}

	state2 := automaton.Step(state1, 'o')
	if automaton.CanMatch(state2) {
		t.Error("Expected 'wo' can not match 'banana' with a distance of 1")
	}

	state3 := automaton.Step(state1, 'b')
	if !automaton.CanMatch(state3) {
		t.Error("Expected 'wb' can match 'banana' with a distance of 1")
	}
}

func TestIsMatch(t *testing.T) {
	// Create the Automaton
	automaton := CreateAutomaton("woof", 1)

	// Start it
	state0 := automaton.Start()

	state1 := automaton.Step(state0, 'w')
	if !automaton.CanMatch(state1) {
		t.Error("Expected 'w' can match 'woof' with a distance of 1")
	}
	if automaton.IsMatch(state1) {
		t.Error("Expected 'w' to not match 'woof' with a distance of 1")
	}

	state2 := automaton.Step(state1, 'o')
	if !automaton.CanMatch(state2) {
		t.Error("Expected 'wo' can match 'woof' with a distance of 1")
	}
	if automaton.IsMatch(state2) {
		t.Error("Expected 'wo' to not match 'woof' with a distance of 1")
	}

	state3 := automaton.Step(state2, 'f')
	if !automaton.CanMatch(state3) {
		t.Error("Expected 'wof' can match 'woof' with a distance of 1")
	}
	if !automaton.IsMatch(state3) {
		t.Error("Expected 'wof' to match 'woof' with a distance of 1")
	}
}
