# Fast Levenshtein search

This project offers a fast and efficient and efficient fuzzy search on large dictionary. To (try to) offer a good level 
of performances, the two following approaches are used:

* The dictionary is stored as a Trie. So, if the search does not match, it is easy to cut directly a full branch.
* The matching is done by using a Levenshtein automaton which allow to easily propagate the state of the automaton as 
the Trie is walked

The automaton is fully described by Jules Jacobs in this [article](https://julesjacobs.github.io/2015/06/17/disqus-levenshtein-simple-and-fast.html). So please be sure to give him a thumb.

# Usage
## Initialization
### Creation of the dictionary
The first step of the search is the creation of a `Dictionary` structure that is holding all the words. The structure is 
created simply by calling the function `levenshteinsearch.CreateDictionary()` without any specific parameter.

Example
```go
dict := levenshteinsearch.CreateDictionary()
```

### Adding words to the dictionary
Words are added one by one to the dictionary, using its member function Put().

Example
```go
// Add an array of word to the dictionary
for _, word := range allWords {
    dict.Put(word)
}
```

## Requests
### Retrieving the dictionary information
Once initialized, the dictionary has two properties `WordCount` and `UniqueWordCount`, giving information about its content

Example
```go
// Get information about the dictionary
log.Printf("Number of words in the dictionary: %v", dict.WordCount)
log.Printf("Number of unique words in the dictionary: %v", dict.UniqueWordCount)
```

### Retrieving information about a single word
Information about a single word can be retrieved using the function `Get` of the dictionary. The function will return a
pointer to a `WordInformation` structure. Currently this structure only have a single information, named `Count` which 
represents the number of time the word was put in the dictionary. Please, note that `Get` will return _nil_ if the 
requested word is not present in the dictionary.


Example
```go
wordInformation := dict.Get("rabbit")
if wordInformation!=nil {
    log.Printf("Number of times the word 'rabbit' is present: %v", wordInformation.Count)
}	else {
    log.Printf("The word 'rabbit' is not part of the dictionary")
}
```

### Retrieving similar words
The dictionary also allow to query for similar words. The similarity is given by the Levenshtein distance [Wikipedia](https://en.wikipedia.org/wiki/Levenshtein_distance).

For searching the similar words, the dictionary has a function named `SearchAll()`, that takes in parameters the searched 
word and the maximum distance. The function return a `map[string]*WordInformation`. The returned map has as key the found 
similar word and as value the `WordInformation` structure of this word.

```go
// Search all word having maximum Levenshtein distance of 3
wordInformationByWord := dict.SearchAll("rabbit", 3)

// Display the number of similar words found
log.Printf("Number of words close to \"rabbit\" with a distance of %v: %v", distance, len(wordInformationByWord))

// Print all the similar words
for key,value := range wordInformationByWord {
    log.Printf("\tWord: '%v' count: %v", key,value.Count)
}
```

# Example
A full working example is given in the folder `/example/alice/alice.go`.


# Performances
The result was benched against:

 * a very naive: Dictionary is stored as a simple list of strings. For each query, all string of the list is tested
 * a simple map: Dictionary is stored as a map having as key the word and as value the `*WordInformation` structure. Due 
 to the map redundant word are not tested multiple times

 The search is done with text of *Alice's Adventures In Wonderland*. For the tests the sets of words <_rabbit_> and 
 <_rabbit_, _eart_, _the_> are used. Results as follow:

```
go version go1.11 windows/amd64
```
```
goos: windows
goarch: amd64
pkg: bitbucket.org/twuillemin/levenshteinsearch/pkg/levenshteinsearch
BenchmarkNaive1Word-8       	      30	  50888216 ns/op	18399690 B/op	  278693 allocs/op
BenchmarkNaive3Word-8       	      10	 133466320 ns/op	43199491 B/op	  842243 allocs/op
BenchmarkMap1Word-8         	     100	  11002825 ns/op	 3600740 B/op	   36130 allocs/op
BenchmarkMap3Word-8         	      50	  30208276 ns/op	 9774411 B/op	  108587 allocs/op
BenchmarkOptimized1Word-8   	  500000	      2757 ns/op	    2368 B/op	      50 allocs/op
BenchmarkOptimized3Word-8   	  200000	      8162 ns/op	    7104 B/op	     150 allocs/op
PASS
```

# License

Copyright 2018 Thomas Wuillemin  <thomas.wuillemin@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this project or its content except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
