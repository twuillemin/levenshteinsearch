package levenshteinsearch

import "testing"

func TestSearch(t *testing.T) {

	dict := CreateDictionary()

	dict.Put("banana")
	dict.Put("orange")
	dict.Put("monkey")

	result := dict.SearchAll("banana", 1)
	for word := range result {
		if word != "banana" {
			t.Error("Expected to find 'banana' with a distance of 1")
		}
	}

	result = dict.SearchAll("banan", 1)
	for word := range result {
		if word != "banana" {
			t.Error("Expected to find 'banan' with a distance of 1")
		}
	}

	result = dict.SearchAll("a", 5)
	if len(result) != 2 {
		t.Error("Expected to find 'banana' and 'orange' with a distance of 5")
	}

	result = dict.SearchAll("a", 6)
	if len(result) != 3 {
		t.Error("Expected to find 'banana', 'orange' and 'monkey' with a distance of 6")
	}
}
