package levenshteinsearch

func (dictionary *Dictionary) SearchAll(searchedTerm string, distanceMax int) map[string]*WordInformation {
	// Create the Automaton
	automaton := CreateAutomaton(searchedTerm, distanceMax)

	// Start the search
	state := automaton.Start()

	results := map[string]*WordInformation{}

	dictionary.Root.searchAll(automaton, "", nil, state, &results)

	return results
}

func (trie *RuneTrie) searchAll(automaton *LevenshteinAutomaton, prefix string, nodeCharacter *rune, automatonState AutomatonState, results *map[string]*WordInformation) {

	var newState AutomatonState
	currentWord := ""

	// The first character will be null for the root
	if nodeCharacter != nil {
		// Add the given char to the state
		newState = automaton.Step(automatonState, *nodeCharacter)
		// If the state can't match, stop here
		if !automaton.CanMatch(newState) {
			return
		}

		// Compute the current word
		currentWord = prefix + string(*nodeCharacter)

		// If the node is a word and if the state is a match, add it to the result
		if (trie.information != nil) && automaton.IsMatch(newState) {
			(*results)[currentWord] = trie.information
		}
	} else {
		newState = automaton.Start()
	}

	// Do the children
	for character, child := range trie.children {
		child.searchAll(automaton, currentWord, &character, newState, results)
	}
}
