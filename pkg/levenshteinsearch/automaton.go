package levenshteinsearch

import (
	"fmt"
	"sort"
)

// LevenshteinAutomaton is simply the definition of the automaton.
type LevenshteinAutomaton struct {
	distanceMax       int
	searchedTermRunes []rune
}

// GetDistanceMax returns the maximum distance defined for this automaton
func (automaton *LevenshteinAutomaton) GetDistanceMax() int {
	return automaton.distanceMax
}

// GetSearchedTerm returns the searched term defined for this automaton
func (automaton *LevenshteinAutomaton) GetSearchedTerm() (out string) {

	// Not really efficient, but should be short. And this function is
	// probably not useful
	for _, v := range automaton.searchedTermRunes {
		out += string(v)
	}
	return
}

// AutomatonState represents a state of the automaton composed of indices and the values corresponding to the indices.
// Note that as the data are for internal use, the object is opaque.
//
// For example, instead of representing the state as:
// state = [3,3,3,3,3,3,3,3,3,3,3,2,1,2,3,3,3,3,3,3]
//
// The following is in fact used:
// indices = [11,12,13]
// values = [2,1,2]
type AutomatonState struct {
	indices []int
	values  []int
}

// getHash generates a hash value for a state
func (state AutomatonState) getHash() int {
	var total = 17
	for _, value := range state.indices {
		total = total*37 + value
	}
	for _, value := range state.values {
		total = total*37 + value
	}
	return total
}

// CreateAutomaton creates a new automaton
func CreateAutomaton(searchedTerm string, distanceMax int) *LevenshteinAutomaton {
	return &LevenshteinAutomaton{
		distanceMax:       distanceMax,
		searchedTermRunes: []rune(searchedTerm),
	}
}

// Start gives the initial state allowing to step into the automaton
func (automaton *LevenshteinAutomaton) Start() AutomatonState {
	indices := make([]int, automaton.distanceMax+1)
	for i := 0; i < automaton.distanceMax+1; i++ {
		indices[i] = i
	}
	values := make([]int, automaton.distanceMax+1)
	for i := 0; i < automaton.distanceMax+1; i++ {
		values[i] = i
	}

	return AutomatonState{
		indices: indices,
		values:  values,
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// Step steps through the automaton by generating the next state based on the current one + the given
// char.
func (automaton *LevenshteinAutomaton) Step(state AutomatonState, character rune) AutomatonState {
	var newIndices []int
	var newValues []int

	if (len(state.indices) > 0) && (state.indices[0] == 0) && (state.values[0] < automaton.distanceMax) {
		newIndices = []int{0}
		newValues = []int{state.values[0] + 1}
	} else {
		newIndices = []int{}
		newValues = []int{}
	}

	for counter, value := range state.indices {
		if value == len(automaton.searchedTermRunes) {
			break
		}

		var cost int
		if automaton.searchedTermRunes[value] == character {
			cost = 0
		} else {
			cost = 1
		}

		val := state.values[counter] + cost

		if (len(newIndices) > 0) && (newIndices[len(newIndices)-1] == value) {
			val = min(val, newValues[len(newValues)-1]+1)
		}
		if ((counter + 1) < len(state.indices)) && (state.indices[counter+1] == value+1) {
			val = min(val, state.values[counter+1]+1)
		}
		if val <= automaton.distanceMax {
			newIndices = append(newIndices, value+1)
			newValues = append(newValues, val)
		}
	}

	return AutomatonState{
		indices: newIndices,
		values:  newValues,
	}
}

// IsMatch returns true if the given states is matching
func (automaton *LevenshteinAutomaton) IsMatch(state AutomatonState) bool {
	return (len(state.indices) > 0) && (state.indices[len(state.indices)-1] == len(automaton.searchedTermRunes))
}

// CanMatch returns true if the given states can match
func (automaton *LevenshteinAutomaton) CanMatch(state AutomatonState) bool {
	return len(state.indices) > 0
}

// digraphInformation is a structure filled during the recursive walk of the generated digraph. It holds
// together the information of the digraph
type digraphInformation struct {
	counter     int
	states      map[int]*int
	transitions []transitionDetail
	matching    []int
}

// transitionDetail represents the detail of a single transition
type transitionDetail struct {
	from      int
	to        int
	character rune
}

// CreateDigraph returns the textual representation of an automaton to be rendered with graphviz
func CreateDigraph(searchedTerm string, distanceMax int) []string {
	automaton := CreateAutomaton(searchedTerm, distanceMax)

	digraphInfo := &digraphInformation{
		counter:     0,
		states:      make(map[int]*int),
		transitions: make([]transitionDetail, 0),
		matching:    make([]int, 0),
	}

	automaton.explore(digraphInfo, automaton.Start())

	// Sort the transition for a "more" humanly readable output
	sort.Slice(digraphInfo.transitions, func(i, j int) bool {
		if digraphInfo.transitions[i].from < digraphInfo.transitions[j].from {
			return true
		} else {
			return digraphInfo.transitions[i].to < digraphInfo.transitions[j].to
		}
	})

	result := make([]string, 0, 100)

	result = append(result, "digraph G {\n")
	for _, transition := range digraphInfo.transitions {
		result = append(result, fmt.Sprintf("%v -> %v [label=\" %c \"]\n", transition.from, transition.to, transition.character))
	}

	for _, matching := range digraphInfo.matching {
		result = append(result, fmt.Sprintf("%v [style=filled]\n", matching))
	}
	result = append(result, "}\n")

	return result
}

// explore is an internal function that recursively generates the full automaton. This function receive a
// state to explore and then returns the id of the now explored state
func (automaton *LevenshteinAutomaton) explore(digraphInfo *digraphInformation, state AutomatonState) int {

	// Search if the same state was previously seen
	stateHash := state.getHash()
	previous := digraphInfo.states[stateHash]
	// If yes, just use the previous one
	if previous != nil {
		return *previous
	}

	// Otherwise references the new state
	stateId := digraphInfo.counter
	digraphInfo.counter++
	digraphInfo.states[stateHash] = &stateId

	// If the state is matching, references it in the matching list
	if automaton.IsMatch(state) {
		digraphInfo.matching = append(digraphInfo.matching, stateId)
	}

	// The recursively explore the new found state
	for _, c := range automaton.getAllTransitions(state) {
		newState := automaton.Step(state, c)
		nextStateId := automaton.explore(digraphInfo, newState)
		digraphInfo.transitions = append(digraphInfo.transitions, transitionDetail{
			from:      stateId,
			to:        nextStateId,
			character: c,
		})
	}

	return stateId
}

// getAllTransitions generates all transitions from the given state. These are the transitions
// for the runes that are part of the original word plus the generic '*' transition
func (automaton *LevenshteinAutomaton) getAllTransitions(state AutomatonState) []rune {

	temp := make(map[rune]int)

	// Put each rune pointed by the indices in a map
	// The map is just here to remove duplicate
	for _, value := range state.indices {
		if value < len(automaton.searchedTermRunes) {
			temp[automaton.searchedTermRunes[value]] = 0
		}
	}
	// Convert the map back to a list
	transitions := make([]rune, 0, len(temp)+1)
	for key := range temp {
		transitions = append(transitions, key)
	}

	// Append the generic '*' transition
	transitions = append(automaton.getAllTransitions(state), '*')

	// Luxury, sort the transitions to ensure stable results. As generated digraph
	// are not really useful for "production" this sort should not have any real life
	// impact
	sort.Slice(transitions, func(i, j int) bool {
		return transitions[i] < transitions[j]
	})

	return transitions
}
