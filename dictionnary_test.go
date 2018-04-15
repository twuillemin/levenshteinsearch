package levenshteinsearch

import "testing"

func TestGetPut(t *testing.T) {

	dict := CreateDictionary()

	dict.Put("banana")
	wordInfo := dict.Get("banana")

	if wordInfo == nil {
		t.Error("Expected to retrieve the word info for 'banana'")
	}
	if wordInfo.Count != 1 {
		t.Error("Expected the word info for 'banana' to have a count of 1")
	}
	if dict.WordCount != 1 {
		t.Error("Expected the dictionnary to have 1 word")
	}
	if dict.UniqueWordCount != 1 {
		t.Error("Expected the dictionnary to have 1 unique word")
	}


	dict.Put("orange")
	wordInfo = dict.Get("orange")

	if wordInfo == nil {
		t.Error("Expected to retrieve the word info for 'orange'")
	}
	if wordInfo.Count != 1 {
		t.Error("Expected the word info for 'orange' to have a count of 1")
	}
	if dict.WordCount != 2 {
		t.Error("Expected the dictionnary to have 2 word")
	}
	if dict.UniqueWordCount != 2 {
		t.Error("Expected the dictionnary to have 2 unique word")
	}

	dict.Put("banana")
	wordInfo = dict.Get("banana")

	if wordInfo == nil {
		t.Error("Expected to retrieve the word info for 'banana'")
	}
	if wordInfo.Count != 2 {
		t.Error("Expected the word info for 'banana' to have a count of 2")
	}
	if dict.WordCount != 3 {
		t.Error("Expected the dictionnary to have 3 word")
	}
	if dict.UniqueWordCount != 2 {
		t.Error("Expected the dictionnary to have 2 unique word")
	}

	wordInfo = dict.Get("monkey")
	if wordInfo != nil {
		t.Error("Expected to not retrieve word info for 'monkey'")
	}
}
