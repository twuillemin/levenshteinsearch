package levenshteinsearch

// Dictionary holds the root node of the trie and some other useful information
type Dictionary struct {
	Root            RuneTrie
	WordCount       int
	UniqueWordCount int
}

// WordInformation holds the various information relative to a single word. As of now
// The only information is the number of time the word was added
type WordInformation struct {
	Count int
}

// RuneTrie is a trie of runes with string keys and WordInformation values.
type RuneTrie struct {
	information *WordInformation
	children    map[rune]*RuneTrie
}

// NewRuneTrie allocates and returns a new *RuneTrie.
func CreateDictionary() *Dictionary {
	return &Dictionary{
		Root: RuneTrie{
			information: nil,
			children:    make(map[rune]*RuneTrie),
		},
		WordCount:       0,
		UniqueWordCount: 0,
	}
}

// NewRuneTrie allocates and returns a new *RuneTrie.
func NewRuneTrie() *RuneTrie {
	return &RuneTrie{
		information: nil,
		children:    make(map[rune]*RuneTrie),
	}
}

// Get returns the value stored at the given key. Returns nil if the key is not found.
func (dictionary *Dictionary) Get(key string) *WordInformation {
	node := &dictionary.Root
	for _, r := range key {
		node = node.children[r]
		if node == nil {
			return nil
		}
	}
	return node.information
}

// Put inserts the value into the trie at the given key, updating any
// existing information. It returns true if the put adds a new value, false
// if it replaces an existing value.
func (dictionary *Dictionary) Put(key string) bool {
	node := &dictionary.Root

	// Rune by rune up to the node
	for _, r := range key {
		child, _ := node.children[r]
		if child == nil {
			child = NewRuneTrie()
			node.children[r] = child
		}
		node = child
	}

	// Does node have an existing value?
	var isNewVal bool
	if node.information == nil {
		isNewVal = true
		node.information = &WordInformation{
			Count: 1,
		}
		dictionary.UniqueWordCount++
	} else {
		isNewVal = false
		node.information.Count++
	}

	dictionary.WordCount++

	return isNewVal
}
